package sqls

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const createSQL = "./sqls/schemas/create/"
const insertSQL = "./sqls/ddls/insert/"

func InitDB() {
	// DB
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("❌ Database is Not Opened :: %d", err)
	}
	defer db.Close()

	dirs, err := os.ReadDir(createSQL)

	if err != nil {
		badSQLDir(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		path := createSQL + dir.Name()
		log.Println(path)
	}

	log.Println("✅ SQL 디렉토리 읽기 성공")

	initializedGakSoos(db)
}

func initializedGakSoos(db *sql.DB) {
	dirs, err := os.ReadDir(insertSQL)
	if err != nil {
		badSQLDir(err)
	}
	for _, dir := range dirs {
		log.Printf("dirdir :: %s", dir)
		if dir.Name() == "" {
		}

	}
	// for i := 1; i <= 45; i++ {
	// 	db.Exec()

	// }
}

func badSQLDir(err error) {
	log.Fatalf("❌ SQL 디렉토리 읽기 실패 %s :: ", err)
}
