package sdk

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)
type ClientInfo struct {
	conn *grpc.ClientConn
	Client ultipa.UltipaRpcsClient
	Host string
}

func (t *ClientInfo) init(conn *grpc.ClientConn, host string) {
	client := ultipa.NewUltipaRpcsClient(conn)
	t.Client = client
	t.conn = conn
	t.Host = host
}
func (t *ClientInfo) Close(){
	if t.conn != nil {
		t.conn.Close()
	}
}

type ClientType int32
var (
	ClientType_Default 	ClientType = 0
	ClientType_Algo 	ClientType = 1
	ClientType_Update 	ClientType = 2
	ClientType_Leader 	ClientType = 3
)
var (
	UQL_Command_Global = []string{
		"listUser", "getUser", "createUser", "updateUser",
		"deleteUser", "grant", "revoke", "listPolicy", "getPolicy",
		"createPolicy", "updatePolicy", "deletePolicy", "listPrivilege",
		"stat", "listGraph", "getGraph", "createGraph", "dropGraph", "top","kill",
	}
	UQL_Command_Write = [] string {
		"insert","update","delete","drop","create","LTE","UFE",
		"clearTask","createUser","updateUser","deleteUser","grant",
		"revoke","createPolicy","updatePolicy","deletePolicy","createIndex",
		"dropIndex","createGraph","dropGraph","stopTask",
	}
	UQL_Command_Extra = [] string {
		"top","kill","showTask","stopTask","clearTask","show","stat","listGraph","listAlgo","getGraph",
		"createPolicy","deletePolicy","listPolicy","getPolicy","grant","revoke","listPrivilege","getUser",
		"getSelfInfo","createUser","updateUser","deleteUser",
	}
)

func UqlIsGlobal(uqlStr string)  bool {
	uql := utils.UQL{}
	uql.Parse(uqlStr)
	_, f := Find(UQL_Command_Global, uql.Command)
	return f
}
func UqlIsExtra(uqlStr string)  bool{
	uql := utils.UQL{}
	uql.Parse(uqlStr)
	_, f := Find(UQL_Command_Extra, uql.Command)
	return f
}
func UqlIsWrite(uqlStr string) bool {
	uql := utils.UQL{}
	uql.Parse(uqlStr)
	_, f := Find(UQL_Command_Write, uql.Command)
	return f
}
func UqlIsAlgo(uqlStr string) bool {
	uql := utils.UQL{}
	uql.Parse(uqlStr)
	return uql.Command == "algo"
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

type HostManager struct {
	GraphSetName string
	LeaderHost string
	FollowersHost []string
	RaftReady bool

	username string
	password string
	crtFile string

	leaderClientInfo *ClientInfo
	algoClientInfo *ClientInfo
	defaultClientInfo *ClientInfo
	otherFollowerClientInfos []*ClientInfo
}
func (t *HostManager) Init(graphSetName string, host string, username string, password string, crtFile string)  {
	t.GraphSetName = graphSetName
	t.username = username
	t.password = password
	t.crtFile = crtFile
	t.LeaderHost = host
	clientInfo, _ := t.createClientInfo(host)
	t.defaultClientInfo = clientInfo
}
func (t *HostManager) createClientInfo(host string) (*ClientInfo, error) {
	var opts []grpc.DialOption
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(-1)))
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(-1)))
	if len(t.crtFile) == 0 {
		// 兼容2.0
		opts = append(opts, grpc.WithInsecure())
		conn, _ := grpc.Dial(host, opts...)
		clientInfo := ClientInfo{}
		clientInfo.init(conn, host)
		return &clientInfo, nil
	} else {
		creds, err := credentials.NewClientTLSFromFile(t.crtFile, "ultipa")
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
		conn, _ := grpc.Dial(host, opts...)
		clientInfo := ClientInfo{}
		clientInfo.init(conn, host)
		return &clientInfo, nil
	}
}

func (t *HostManager)GetAllHosts()  *[]string{
	hosts := []string{
		t.LeaderHost,
	}
	if t.FollowersHost != nil && len(t.FollowersHost) > 0 {
		hosts = append(hosts, t.FollowersHost...)
	}
	return &hosts

}
func (t*HostManager) chooseClientInfo(clientType ClientType, uql string, readModeNonConsistency bool, useHost string) *ClientInfo  {
	if useHost != "" {
		for _, clientInfo := range t.getAllClientInfos()  {
			if clientInfo.Host == useHost {
				return clientInfo
			}
		}
	}
	if clientType == ClientType_Default && uql != "" {
		if (UqlIsAlgo(uql)) {
			clientType = ClientType_Algo
		} else if (UqlIsWrite(uql)) {
			clientType = ClientType_Update
		}
	}
	if clientType == ClientType_Algo {
		if t.algoClientInfo != nil {
			return t.algoClientInfo
		}
		return t.defaultClientInfo
	}
	if clientType == ClientType_Update || clientType == ClientType_Leader || readModeNonConsistency == false {
		if t.leaderClientInfo != nil {
			return t.leaderClientInfo
		}
		return t.defaultClientInfo
	}

	// 负载均衡，随机挑一个
	all := t.getAllClientInfos()
	return all[rand.Intn(len(all))]
}
func (t *HostManager) getAllClientInfos() []*ClientInfo  {
	all := []*ClientInfo{
		t.defaultClientInfo,
	}
	if t.algoClientInfo != nil {
		all = append(all, t.algoClientInfo)
	}
	if t.leaderClientInfo != nil && t.leaderClientInfo.Host != t.defaultClientInfo.Host{
		all = append(all, t.leaderClientInfo)
	}
	if t.otherFollowerClientInfos != nil {
		all = append(all, t.otherFollowerClientInfos...)
	}
	return all
}
func (t *HostManager) SetClients(leaderHost string, followersHost [] string)  {
	t.LeaderHost = leaderHost
	if t.defaultClientInfo.Host == leaderHost {
		t.leaderClientInfo = t.defaultClientInfo
	} else {
		clientInfo, _ := t.createClientInfo(leaderHost)
		t.leaderClientInfo = clientInfo
	}
	t.FollowersHost = followersHost
	if t.FollowersHost != nil && len(t.FollowersHost) > 0 {
		info, _ := t.createClientInfo(followersHost[rand.Intn(len(followersHost))])
		t.algoClientInfo = info
		for _, host := range followersHost {
			if host == info.Host {
				continue
			}
			info, _ := t.createClientInfo(host)
			t.otherFollowerClientInfos = append(t.otherFollowerClientInfos, info)
		}
	} else {
		t.algoClientInfo = t.leaderClientInfo
	}

}

type HostManagerControl struct {
	InitHost string
	username string
	password string
	crtFile string
	ReadModeNonConsistency bool
	AllHostManager map[string]*HostManager
}

func (t *HostManagerControl) Init(initHost string, username string, password string, crtFile string, readModeNonConsistency bool)  {
	t.InitHost = initHost
	t.username = username
	t.password = password
	t.ReadModeNonConsistency = readModeNonConsistency
	t.AllHostManager = map[string]*HostManager{}
}

func (t*HostManagerControl) chooseClientInfo(graphSetName string, clientType ClientType, uql string, useHost string) *ClientInfo {
	hostManager := t.getHostManager(graphSetName)
	return hostManager.chooseClientInfo(clientType, uql, t.ReadModeNonConsistency, useHost)
}
func (t*HostManagerControl) getHostManager(graphSetName string) *HostManager {
	hostManager := t.AllHostManager[graphSetName]
	if hostManager == nil {
		hostManager = t.upsetHostManager(graphSetName, t.InitHost)
	}
	return hostManager
}
func (t*HostManagerControl) upsetHostManager(graphSetName string, initHost string) *HostManager{
	hostManager := HostManager{}
	hostManager.Init(graphSetName, initHost, t.username, t.password, t.crtFile)
	t.AllHostManager[graphSetName] = &hostManager
	return &hostManager
}

const RAFT_GLOBAL = "global"

func (t *HostManagerControl) GetAllHosts() *[]string{
	hostManager := t.getHostManager(RAFT_GLOBAL)
	return hostManager.GetAllHosts()
}
func (t *HostManagerControl) CloseAll()  {
	for _, hostManager := range t.AllHostManager{
		if hostManager != nil {
			saveClose(hostManager.defaultClientInfo)
			saveClose(hostManager.leaderClientInfo)
			saveClose(hostManager.algoClientInfo)
		}
	}
}


type DefaultConfig struct {
	GraphSetName string
	TimeoutWithSeconds uint32
	ResponseWithRequestInfo bool
	ReadModeNonConsistency bool
}

type Connection struct {
	HostManagerControl *HostManagerControl
	metadataKV *[]string
	username string
	password string
	crtFile string
	DefaultConfig *DefaultConfig
}

func GetConnection(host string, username string, password string, crtFile string, defaultConfig *DefaultConfig) (*Connection, error){
	connect := Connection{}
	err := connect.init(host, username, password, crtFile, defaultConfig)
	if err != nil {
		return nil, err
	}
	return &connect, nil
}
func (t *Connection) SetDefaultConfig(defaultConfig *DefaultConfig)  {
	if t.DefaultConfig == nil {
		t.DefaultConfig = &DefaultConfig{ "default", 15, false, false}
	}
	if defaultConfig != nil {
		if &defaultConfig.GraphSetName != nil {
			t.DefaultConfig.GraphSetName = defaultConfig.GraphSetName
		}
		if &defaultConfig.TimeoutWithSeconds != nil {
			t.DefaultConfig.TimeoutWithSeconds = defaultConfig.TimeoutWithSeconds
		}
		if &defaultConfig.ResponseWithRequestInfo != nil {
			t.DefaultConfig.ResponseWithRequestInfo = defaultConfig.ResponseWithRequestInfo
		}
		if &defaultConfig.ReadModeNonConsistency != nil {
			t.DefaultConfig.ReadModeNonConsistency = defaultConfig.ReadModeNonConsistency
		}
	}
}
func saveClose(clientInfo *ClientInfo)  {
	if clientInfo != nil {
		clientInfo.Close()
	}
}
func (t *Connection) CloseAll()  {
	t.HostManagerControl.CloseAll()
}

func (t *Connection) init(host string, username string, password string, crt string, config *DefaultConfig)  error {
	t.username = username
	t.password = password
	t.crtFile = crt
	t.SetDefaultConfig(config)
	kv := []string{"user", username, "password", password}
	t.metadataKV = &kv
	hmControl := HostManagerControl{}
	hmControl.Init(host, username, password, crt, t.DefaultConfig.ReadModeNonConsistency)
	t.HostManagerControl = &hmControl
	return nil
}

const (
	TIMEOUT_DEFAUL time.Duration = time.Minute
)

type GetClientInfoResult struct {
	ClientInfo *ClientInfo
	Context context.Context
	CancelFunc context.CancelFunc
	Host string
	GraphSetName string
}
type GetClientInfoParams struct {
	GraphSetName string
	ClientType ClientType
	Uql string
	IsGlobal bool
	IgnoreRaft bool
	UseHost string
}

func (t *Connection) getClientInfo(params *GetClientInfoParams) *GetClientInfoResult {
	goGraphSetName := t.getGraphSetName(params.GraphSetName, params.Uql, params.IsGlobal)
	timeout := TIMEOUT_DEFAUL
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	kv := []string{"graph_name", goGraphSetName}
	kv = append(kv, *t.metadataKV...)
	ctx = metadata.AppendToOutgoingContext(ctx, kv...)
	//ctx = metadata.AppendToOutgoingContext(ctx, *t.metadataKV...)
	//defer cancel()

	if params.IgnoreRaft == false && t.HostManagerControl.getHostManager(goGraphSetName).RaftReady == false {
		t.RefreshRaftLeader(t.HostManagerControl.InitHost, &SdkRequest_Common{
			GraphSetName: goGraphSetName,
		})
		t.HostManagerControl.getHostManager(goGraphSetName).RaftReady = true
	}

	clientInfo := t.HostManagerControl.chooseClientInfo(goGraphSetName, params.ClientType, params.Uql, params.UseHost)
	return &GetClientInfoResult{
		ClientInfo: clientInfo,
		Context: ctx,
		CancelFunc: cancel,
		Host: clientInfo.Host,
		GraphSetName: goGraphSetName,
	}
}
func (t *Connection) getGraphSetName(currentGraphName string,uql string, isGlobal bool) string{
	if isGlobal || (uql != "" && UqlIsGlobal(uql)) {
		return RAFT_GLOBAL
	}
	if currentGraphName != "" {
		return currentGraphName
	}
	return t.DefaultConfig.GraphSetName
}
func (t *Connection) TestConnect(commonReq *SdkRequest_Common)  (bool, error) {
	if commonReq == nil {
		commonReq = &SdkRequest_Common{}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		IsGlobal: true,
		UseHost: commonReq.UseHost,
	})
	defer clientInfo.CancelFunc()
	name := "MyTest"
	res, err := clientInfo.ClientInfo.Client.SayHello(clientInfo.Context, &ultipa.HelloUltipaRequest{
		Name: name,
	})
	if err != nil {
		return false, err
	}
	if res.Message != name + " Welcome To Ultipa!"{
		return false, err
	}
	return true, nil
}

type RaftLeaderResSimple struct {
	Code ultipa.ErrorCode
	Message string
	LeaderHost string
	FollowersHost []string
}
func (t *Connection) autoGetRaftLeader(host string, commonReq *SdkRequest_Common, retry int) (*RaftLeaderResSimple,error){
	conn, err := GetConnection(host, t.username, t.password, t.crtFile, t.DefaultConfig)
	// 用一次就关掉
	defer conn.CloseAll()
	if err != nil {
		return nil, err
	}
	res := conn.GetLeaderReuqest(commonReq)
	errorCode := res.Status.Code
	switch errorCode {
	case types.ErrorCode_SUCCESS:
		followers := res.Status.ClusterInfo.RaftPeers
		return &RaftLeaderResSimple{
			Code: errorCode,
			LeaderHost: host,
			FollowersHost: utils.Remove(followers, host),
		}, nil
	case types.ErrorCode_NOT_RAFT_MODE:
		return &RaftLeaderResSimple{
			Code:          types.ErrorCode_SUCCESS,
			LeaderHost:    host,
			FollowersHost: nil,
		}, nil
	case types.ErrorCode_RAFT_REDIRECT:
		if retry > 1 {
			return &RaftLeaderResSimple{
				Code: types.ErrorCode_UNKNOW,
				Message: "raft redirect too many times",
			}, nil
		}
		return t.autoGetRaftLeader(res.Status.ClusterInfo.Redirect, commonReq, retry+1)
	}
	return &RaftLeaderResSimple{
		Code: errorCode,
		Message: res.Status.Message,
	}, nil
}

type Retry struct {
	Current uint32
	Max uint32
}
type (
	SdkRequest_Common struct {
		GraphSetName string
		TimeoutSeconds time.Duration
		Retry *Retry
		UseHost string
	}
)

func (t *Connection)  RefreshRaftLeader(redirectHost string, commonReq *SdkRequest_Common) error{
	if commonReq == nil {
		commonReq = &SdkRequest_Common{}
	}
	var hosts []string
	if redirectHost != "" {
		hosts = append(hosts, redirectHost)
	} else {
		hosts = append(hosts, *t.HostManagerControl.GetAllHosts()...)
	}
	goGraphName := t.getGraphSetName(commonReq.GraphSetName, "", false)
	for _, host := range hosts{
		res, err := t.autoGetRaftLeader(host, commonReq, 0)
		if err != nil {
			return err
		}
		if res.Code == types.ErrorCode_SUCCESS {
			leaderHost := res.LeaderHost
			followersHost := res.FollowersHost
			hostManager := t.HostManagerControl.upsetHostManager(goGraphName, leaderHost)
			hostManager.SetClients(leaderHost, followersHost)
			return nil
		}
		return fmt.Errorf("%v - %v", res.Code.String(), res.Message)
	}
	return fmt.Errorf("Unknow error! ")

}
