package sdk

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
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
	t.conn.Close()
}

type Connection struct {
	ClientInfo *ClientInfo
	MetadataKV *[]string
}

func (t *Connection) Init(host string, username string, password string, crt string)  error {
	kv := []string{"user-metadata", username, "passwd-metadata", password}
	t.MetadataKV = &kv
	var opts []grpc.DialOption
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(-1)))
	//opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(-1)))
	if len(crt) == 0 {
	 	opts = append(opts, grpc.WithInsecure())
		conn, _ := grpc.Dial(host, opts...)
		clientInfo := ClientInfo{}
		clientInfo.init(conn)
		t.ClientInfo = &clientInfo
		return nil
	}
	creds, err := credentials.NewClientTLSFromFile(crt, "ultipa")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, _ := grpc.Dial(host, opts...)
	clientInfo := ClientInfo{}
	clientInfo.init(conn)
	t.ClientInfo = &clientInfo
	return nil
}

const (
	TIMEOUT_DEFAUL time.Duration = time.Minute
)

func (t *Connection) choiseClient(timeout time.Duration) (_clientInfo *ClientInfo, _context context.Context, _cancelFunc context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx = metadata.AppendToOutgoingContext(ctx, *t.MetadataKV...)
	//defer cancel()
	return t.ClientInfo, ctx, cancel
}
func (t *Connection) TestConnect()  (bool, error) {
	clientInfo, ctx, cancel := t.choiseClient(time.Second * 3)
	defer cancel()
	name := "MyTest"
	res, err := clientInfo.Client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: name,
	})
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(res)
	if res.Message != name + " Welcome To Ultipa!"{
		return false, err
	}
	return true, nil
}
func (t *Connection) TestUQLConnect()  {

}

//constructor(host, username, password, crt) {
//if (!crt) {
//this.client = new ultipa_grpc_pb_1.UltipaRpcsClient(host, grpc_1.default.credentials.createInsecure(), {
//"grpc.max_send_message_length": -1,
//"grpc.max_receive_message_length": -1,
//});
//return;
//}
//var ssl_creds = grpc_1.default.credentials.createSsl(crt);
//this.metadata = new grpc_1.default.Metadata();
//this.metadata.add("user-metadata", username);
//this.metadata.add("passwd-metadata", password);
//this.client = new ultipa_grpc_pb_1.UltipaRpcsClient(host, ssl_creds, {
//"grpc.ssl_target_name_override": "ultipa",
//"grpc.default_authority": "ultipa",
//"grpc.max_send_message_length": -1,
//"grpc.max_receive_message_length": -1,
//});
//}

// func Connect(username, password, host string) (_client Client, _conn *ClientConn, _err error) {
// 	// ultipa.SayHello()
// 	var opts []grpc.DialOption
// 	opts = append(opts, grpc.WithInsecure())
// 	conn, err := grpc.Dial(host, opts...)

// 	if err != nil {
// 		log.Printf("fail to dial: %v", err)
// 		return nil, nil, err
// 	}

// 	client := ultipa.NewUltipaRpcsClient(conn)

// 	return client, conn, nil

// }

// // TestConnect return whether connection is alive
// func TestConnect(client Client) (bool, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

// 	defer cancel()

// 	_, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
// 		Name: "test",
// 	})

// 	if err != nil {
// 		if strings.Contains(err.Error(), "code = Unavailable") {
// 			return false, errors.New("Unavailable")
// 		}
// 		return false, err
// 	}
// 	return true, nil
// }

//class _Connection {
//client: UltipaRpcsClient;
//metadata: grpc.Metadata;
//constructor(host, username, password, crt) {
//if (!crt) {
//this.client = new UltipaRpcsClient(
//host,
//grpc.credentials.createInsecure(),
//{
//"grpc.max_send_message_length": -1,
//"grpc.max_receive_message_length": -1,
//}
//);
//return;
//}
//
//var ssl_creds = grpc.credentials.createSsl(crt);
//
//this.metadata = new grpc.Metadata();
//this.metadata.add("user-metadata", username);
//this.metadata.add("passwd-metadata", password);
//
//this.client = new UltipaRpcsClient(host, ssl_creds, {
//"grpc.ssl_target_name_override": "ultipa",
//"grpc.default_authority": "ultipa",
//"grpc.max_send_message_length": -1,
//"grpc.max_receive_message_length": -1,
//});
//}
///**
// * test connection is available
// */
//async test() {
//let request = new HelloUltipaRequest();
//request.setName("ultipa");
//
//return await new Promise<boolean>((resolve, reject) => {
//this.client.sayHello(request, this.metadata, (err, res) => {
//if (err) {
//reject(err);
//return;
//}
//let msg = res.getMessage();
//if (msg != "ultipa Welcome To Ultipa!") {
//resolve(false);
//}
//resolve(true);
//});
//});
//}
//}

// Connect a ultipa db host by hostname or ip
//func Connect(username, password, host string) (_client ultipa.UltipaRpcsClient, _conn *ClientConn, _err error) {
//	// ultipa.SayHello()
//	var opts []grpc.DialOption
//	opts = append(opts, grpc.WithInsecure())
//	conn, err := grpc.Dial(host, opts...)
//
//	if err != nil {
//		log.Printf("fail to dial: %v", err)
//		return nil, nil, err
//	}
//
//	client := ultipa.NewUltipaRpcsClient(conn)
//
//	return client, conn, nil
//
//}
//
//// TestConnect return whether connection is alive
//func TestConnect(client Client) (bool, error) {
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
//
//	defer cancel()
//
//	_, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
//		Name: "test",
//	})
//
//	if err != nil {
//		if strings.Contains(err.Error(), "code = Unavailable") {
//			return false, errors.New("Unavailable")
//		}
//		return false, err
//	}
//	return true, nil
//}
