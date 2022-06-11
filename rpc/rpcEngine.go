package rpc

import (
	"SearchEngineV2/rpc/com/Hint"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"os"
)

type RpcEngine struct {
	Engine    *Hint.GetHintServiceClient
	Transport *thrift.TSocket
}

func NewRpcEngine() *RpcEngine {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	//可以选择任意一种通信协议，但是必须确保服务端和客户端的通信协议一致
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{}) //布尔参数strictRead, strictWrite，读和写时是否加入版本校验。
	transport := thrift.NewTSocketConf("127.0.0.1:8090", &thrift.TConfiguration{})
	useTransport, _ := transportFactory.GetTransport(transport)
	client := Hint.NewGetHintServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:8090", " ", err)
		os.Exit(1)
	}
	fmt.Println("RPC connection complete")

	return &RpcEngine{
		Engine:    client,
		Transport: transport,
	}
}
func (c *RpcEngine) TestConnect() {

}

func (c *RpcEngine) FindEnd(word string) []string {
	res, err := c.Engine.GetString(context.Background(), word)
	if err != nil {
		log.Println("Echo failed:", err)
		return nil
	}
	return res.GetHints()
}

func (c *RpcEngine) close() {
	c.Transport.Close()
}
