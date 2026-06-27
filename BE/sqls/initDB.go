package sqls

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	// DB
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("❌ Database is Not Opened :: %d", err)
	}
	defer db.Close()

	const createSQL = "./sqls/schemas/create/"

	dirs, err := os.ReadDir(createSQL)

	if err != nil {
		log.Fatalf("❌ SQL 디렉토리 읽기 실패 %s :: ", err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		path := createSQL + dir.Name()
		log.Println(path)
	}

	log.Fatalf("✅ SQL 디렉토리 읽기 성공")
}
