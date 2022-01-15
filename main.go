package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("test.leveldb", nil)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	if err = db.Put([]byte("key"), []byte("value"), nil); err != nil {
		fmt.Println(err)
		return
	}

	var value []byte
	if value, err = db.Get([]byte("key"), nil); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(value))

	if err = db.Delete([]byte("key"), nil); err != nil {
		fmt.Println(err)
		return
	}
}