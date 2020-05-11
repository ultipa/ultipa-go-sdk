package sdk

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)
type ClientInfo struct {
	conn *grpc.ClientConn
	Client ultipa.UltipaRpcsClient
}

func (t *ClientInfo) init(conn *grpc.ClientConn) {
	client := ultipa.NewUltipaRpcsClient(conn)
	t.Client = client
	t.conn = conn
}
func (t *ClientInfo) Close(){
	if t.conn != nil {
		t.conn.Close()
	}
}

type Connection struct {
	clientInfo *ClientInfo
	metadataKV *[]string
}

func (t *Connection) CloseAll()  {
	if t.clientInfo != nil {
		t.clientInfo.Close()
	}
}
func (t *Connection) Init(host string, username string, password string, crt string)  error {
	kv := []string{"user-metadata", username, "passwd-metadata", password}
	t.metadataKV = &kv
	var opts []grpc.DialOption
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(-1)))
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(-1)))
	if len(crt) == 0 {
	 	opts = append(opts, grpc.WithInsecure())
		conn, _ := grpc.Dial(host, opts...)
		clientInfo := ClientInfo{}
		clientInfo.init(conn)
		t.clientInfo = &clientInfo
		return nil
	}
	creds, err := credentials.NewClientTLSFromFile(crt, "ultipa")
	if err != nil {
		return err
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, _ := grpc.Dial(host, opts...)
	clientInfo := ClientInfo{}
	clientInfo.init(conn)
	t.clientInfo = &clientInfo
	return nil
}

const (
	TIMEOUT_DEFAUL time.Duration = time.Minute
)

func (t *Connection) choiseClient(timeout time.Duration) (_clientInfo *ClientInfo, _context context.Context, _cancelFunc context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx = metadata.AppendToOutgoingContext(ctx, *t.metadataKV...)
	//defer cancel()
	return t.clientInfo, ctx, cancel
}
func (t *Connection) TestConnect()  (bool, error) {
	clientInfo, ctx, cancel := t.choiseClient(time.Second * 3)
	defer cancel()
	name := "MyTest"
	res, err := clientInfo.Client.SayHello(ctx, &ultipa.HelloUltipaRequest{
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
