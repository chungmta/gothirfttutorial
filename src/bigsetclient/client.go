package bigsetclient

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thirfttutorial/openstars/core/bigset/generic"
)

var defaultCtx = context.Background()

func handleClient(client *generic.TStringBigSetKVServiceClient) (err error) {

	var bsName generic.TStringKey = "HEY"
	firstBS, err := client.CreateStringBigSet(defaultCtx, bsName)

	fmt.Println("!!", err)

	fmt.Println("----", firstBS)
	//client.AssignBigSetName()

	//put
	rs, _ := client.BsPutItem(defaultCtx, bsName, &generic.TItem{
		Key:   []byte("key1"),
		Value: []byte("lakjfd lajdfl akdf"),
	})
	fmt.Println("putItem", rs)

	//get item
	item, _ := client.BsGetItem(defaultCtx, bsName, []byte("key1"))
	//if err {
	//	fmt.Println("ERR", err)
	//}
	fmt.Println("item", item)

	//check exits
	exits, _ := client.BsExisted(defaultCtx, bsName, []byte("key1"))
	fmt.Println("--------exits", exits)

	return err
}

func RunClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	return handleClient(generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
