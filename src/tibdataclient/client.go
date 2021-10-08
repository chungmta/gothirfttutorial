package tibdataclient

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thirfttutorial/openstars/core/bigset/generic"
)

var defaultCtx = context.Background()

func handleClient(client *generic.TIBSDataServiceClient) (err error) {

	r, _ := client.BulkLoad(defaultCtx, generic.TKey(1), &generic.TItemSet{})

	fmt.Println("init", r)
	//client.PutItem(defaultCtx, bsName, &generic.TItem{
	//	Key:   []byte("key1"),
	//	Value: []byte("lakjfd lajdfl akdf"),
	//})

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

	return handleClient(generic.NewTIBSDataServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
