package infra

import (
	"log"
	"os"
)

func ExampleNewDB() {
	db, err := NewDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// 何があっても強い気持ちでCloseする
	// クローズに失敗した場合はそれもしっかりロギングする
	defer func() {
		cerr := db.Close()
		if err != nil {
			log.Println(cerr)
		}
	}()
}
