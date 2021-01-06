package infra

import (
	"os"
)

func ExampleNewDB() {
	db, err := NewDB()
	if err != nil {
		logger.Warnf("failed NewDB(): %w", err)
		os.Exit(1)
	}
	// 何があっても強い気持ちでCloseする
	// クローズに失敗した場合はそれもしっかりロギングする
	defer func() {
		cerr := db.Close()
		if err != nil {
			logger.Warnf("failed db.Close(): %w", cerr)
		}
	}()
}
