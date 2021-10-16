package store

import (
	"context"
	"encoding/json"
	"thirfttutorial/openstars/core/bigset/generic"
	"thirfttutorial/src/modal"
)

type BSStore struct {
	db *generic.TStringBigSetKVServiceClient
}

func (bss *BSStore) CreateStudent(ctx context.Context, s modal.Student) error {
	sm, _ := json.Marshal(s)
	_, err := bss.db.BsPutItem(ctx, generic.TStringKey(modal.Student{}.TableName()), &generic.TItem{Key: s.Id, Value: sm})

	if err != nil {
		return err
	}
	return nil
}

func (bss *BSStore) CreateClass(ctx context.Context, s modal.Class) error {
	sm, _ := json.Marshal(s)
	_, err := bss.db.BsPutItem(ctx, generic.TStringKey(modal.Class{}.TableName()), &generic.TItem{Key: s.Id, Value: sm})

	if err != nil {
		return err
	}
	return nil
}

func (bss *BSStore) GetStudentsByClassID(ctx context.Context, class modal.Class) ([]modal.Student, error) {
	//get class in bs
	result, _ := bss.db.BsGetItem(ctx, generic.TStringKey(modal.Class{}.TableName()), class.Id)

	var cl modal.Class
	if err := json.Unmarshal(result.Item.Value, &cl); err != nil {
		return nil, err
	}

	lstStudent := make([]modal.Student, len(cl.Students))
	//get student in class
	for i, sid := range cl.Students {
		rs, _ := bss.db.BsGetItem(ctx, generic.TStringKey(modal.Student{}.TableName()), sid)

		if err := json.Unmarshal(rs.Item.Value, &lstStudent[i]); err != nil {
			return nil, err
		}
	}

	return lstStudent, nil
}
