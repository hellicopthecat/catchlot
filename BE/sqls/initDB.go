package sqls

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
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
		badSQLFile(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		path := createSQL + dir.Name()
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("❌ Initialized Create Schemas is Failed")
		}

		var exists int
		tableName := strings.TrimSuffix(strings.TrimPrefix(dir.Name(), "c_"), ".sql")
		db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&exists)
		if exists > 0 {
			log.Printf("⚠️ %s is Already Exists", tableName)
			continue
		}
		db.Exec(string(content))
		log.Println("✅ SQL Create Is Success")
	}

	log.Println("✅ SQL 디렉토리 읽기 성공")

	initializedGakSoos(db)
}

func initializedGakSoos(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM gak_soo").Scan(&count)
	if count >= 45 {
		log.Printf("⚠️ 45 Numbers is Already Exists")
		return
	}

	gakSooSQL, err := os.ReadFile(insertSQL + "i_gak_soo.sql")
	if err != nil {
		badSQLFile(err)
	}

	gakSooStatusSQL, err := os.ReadFile(insertSQL + "i_gak_soo_status.sql")
	if err != nil {
		badSQLFile(err)
	}

	tx, _ := db.Begin()
	defer tx.Rollback()

	for i := 1; i <= 45; i++ {
		gakSooID, _ := uuid.NewV7()
		statusID, _ := uuid.NewV7()
		tx.Exec(string(gakSooSQL), gakSooID.String(), i)
		tx.Exec(string(gakSooStatusSQL), statusID.String(), gakSooID.String())
		log.Println("✅ Gak_Soo Insert Is Complete")
	}
	if err := tx.Commit(); err != nil {
		log.Fatalln("❌ Gak_Soo Insert Is UnComplete")
	}
}

func badSQLFile(err error) {
	log.Fatalf("❌ SQL 디렉토리 및 파일 읽기 실패 %s :: ", err)
}
