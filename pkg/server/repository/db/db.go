package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// DB 各repositoryで利用するDB接続情報
var DB *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	DB, err = sql.Open(driverName,
		fmt.Sprintf("root:%s@/papero", password)+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
}
