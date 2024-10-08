package db

import (
	"apiBook/common/log"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_DB(t *testing.T) {
	SetLocalDB("../../data.db")

	table := "test1"
	key := "test1"

	type A struct {
		A1 string
		A2 int
	}

	var val = &A{
		A1: "aaaaa",
		A2: 111,
	}
	var data *A

	err := DB.Set(table, key, val)
	if err != nil {
		t.Log("Set err = ", err)
		return
	}

	err = DB.Get(table, key, &data)
	if err != nil {
		t.Log("Get err = ", err)
		return
	}

	t.Log("data = ", data)

}

func Test_DBGetAll(t *testing.T) {
	SetLocalDB("../../data.db")

	table := "test2"

	type A struct {
		A1 string
		A2 int
	}

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("test%d", i)
		err := DB.Set(table, key, &A{
			A1: key,
			A2: i,
		})
		if err != nil {
			t.Log("Set err = ", err)
			return
		}
	}

	list := make([]*A, 0)

	err := DB.GetAll(table, func(k, v []byte) {
		data := &A{}
		err := json.Unmarshal(v, &data)
		if err != nil {
			log.Error(err)
		} else {

			DB.GetAllSetCache(table, k, data)

			list = append(list, data)
		}
	})

	if err != nil {
		t.Log("Set err = ", err)
		return
	}

	t.Log("list = ", list)
	for _, v := range list {
		t.Log(v)
	}

}
