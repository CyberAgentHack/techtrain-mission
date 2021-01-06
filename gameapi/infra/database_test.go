package infra_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/task4233/techtrain-mission/gameapi/infra"
	"github.com/task4233/techtrain-mission/gameapi/log"
)

var (
	logger = log.MyLogger
)

func ExampleNewDB() {
	db, err := infra.NewDB()
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

func TestNewDB(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		wantDB *sqlx.DB
	}{
		{
			name:   "MySQLとの接続に失敗する",
			wantDB: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := infra.NewDB(); got != tc.wantDB {
				t.Errorf("infra.NewDB() = %v, want = %v", got, tc.wantDB)
			}
		})
	}
}
