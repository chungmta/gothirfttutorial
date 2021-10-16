package store

import "thirfttutorial/openstars/core/bigset/generic"

const (
	BSStoreType = iota + 1
)

func GetStoreFactory(store int, db interface{}) Store {
	switch store {
	case BSStoreType:
		return &BSStore{
			db: db.(*generic.TStringBigSetKVServiceClient),
		}
	}
	return nil
}
