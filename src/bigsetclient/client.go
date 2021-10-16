package bigsetclient

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"thirfttutorial/openstars/core/bigset/generic"
	"thirfttutorial/src/modal"
	store2 "thirfttutorial/src/store"
)

var ctx = context.Background()

func handleClient(client *generic.TStringBigSetKVServiceClient) (err error) {
	//log.Println(": handleClient")
	//var bsName generic.TStringKey = "HEY"
	////firstBS, err := client.CreateStringBigSet(ctx, bsName)
	////if err != nil {
	////	fmt.Println("!!", err)
	////}
	////fmt.Println("----", firstBS)
	////client.AssignBigSetName()
	//
	////put
	//rs, _ := client.BsPutItem(ctx, bsName, &generic.TItem{
	//	Key:   []byte("key1"),
	//	Value: []byte("lakjfd lajdfl akdf"),
	//})
	//fmt.Println("putItem", rs)
	//
	////get item
	//item, _ := client.BsGetItem(ctx, bsName, []byte("key1"))
	////if err {
	////	fmt.Println("ERR", err)
	////}
	//fmt.Println("item", item)
	//
	////check exits
	//exits, _ := client.BsExisted(ctx, bsName, []byte("key1"))
	//fmt.Println("exits--", exits.Existed, exits.Error)
	//
	////total count
	//total, _ := client.GetTotalCount(ctx, bsName)
	//fmt.Println("total:", total)
	//
	////time.Sleep(time.Second)
	//
	////put many item
	//i2 := generic.TItem{Key: []byte("key2"), Value: []byte("value2")}
	//i3 := generic.TItem{Key: []byte("key3"), Value: []byte("value3")}
	//
	//var items []*generic.TItem
	//items = append(items, &i2)
	//items = append(items, &i3)
	//
	//itemSet := generic.TItemSet{
	//	//Items: []*generic.TItem{
	//	//	&i2, &i3,
	//	//},
	//	Items: items,
	//}
	//
	//res, _ := client.BsMultiPut(ctx, bsName, &itemSet, false, false)
	//fmt.Println("put many item-", res)
	//
	//total, _ = client.GetTotalCount(ctx, bsName)
	//fmt.Println("total:", total)
	//
	////get bigset info by name
	//info, _ := client.GetBigSetInfoByName(ctx, bsName)
	//fmt.Println("get bigset info by name", info.GetInfo(), *info.GetInfo().Count)
	//
	////get slice
	//s, _ := client.BsGetSlice(ctx, bsName, 0, int32(total))
	//fmt.Println("get slice:", s.Items)
	//
	//fmt.Println("==================")
	////add 1000 item
	//var items2 []*generic.TItem
	//for i := 0; i < 1000; i++ {
	//	v := generic.TItem{Key: []byte(strconv.FormatInt(time.Now().Unix(), 10) ), Value: []byte("value2")}
	//	items2 = append(items2, &v)
	//}
	//
	//resMul, _ := client.BsMultiPut(ctx,
	//	bsName,
	//	&generic.TItemSet{
	//		Items: items2,
	//	},
	//	false,
	//	false)
	//fmt.Println("Add 1000 item", resMul)
	//
	////get slice BsGetSliceR
	//sr, _ := client.BsGetSliceR(ctx, bsName, 0, 1)
	//fmt.Println("BsGetSliceR", sr.Items, sr.Items.Items[0])
	//fmt.Println("BsGetSliceR", sr.Items.Items[0])
	//fmt.Println("BsGetSliceR", string(sr.Items.Items[0].Key))

	store := store2.GetStoreFactory(store2.BSStoreType, client)

	s1 := modal.Student{
		Id:      []byte("1"),
		Name:    "Student name 1",
		Classes: [][]byte{[]byte("1")},
	}
	s2 := modal.Student{
		Id:      []byte("2"),
		Name:    "Student name 2",
		Classes: [][]byte{[]byte("1")},
	}
	s3 := modal.Student{
		Id:      []byte("3"),
		Name:    "Student name 3",
		Classes: [][]byte{[]byte("1"), []byte("2")},
	}
	s4 := modal.Student{
		Id:      []byte("3"),
		Name:    "Student name 3",
		Classes: [][]byte{[]byte("2")},
	}

	c1 := modal.Class{
		Id:       []byte("1"),
		Name:     "Class name 1",
		Students: [][]byte{[]byte("1"), []byte("2"), []byte("3")},
	}
	c2 := modal.Class{
		Id:       []byte("2"),
		Name:     "Class name 2",
		Students: [][]byte{[]byte("3"), []byte("4")},
	}

	_ = store.CreateStudent(ctx, s1)
	_ = store.CreateStudent(ctx, s2)
	_ = store.CreateStudent(ctx, s3)
	_ = store.CreateStudent(ctx, s4)

	_ = store.CreateClass(ctx, c1)
	_ = store.CreateClass(ctx, c2)

	lStudent, _ := store.GetStudentsByClassID(ctx, c1)

	fmt.Println(lStudent)

	return nil
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
	fmt.Println("isOpenSocket", transport.IsOpen())
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
