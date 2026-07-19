package sqls

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// DB
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("❌ Database is Not Opened :: %d", err)
	}

	dirs, err := os.ReadDir(constants.CreateSQL)

	if err != nil {
		commons.BadSQLFile(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		path := constants.CreateSQL + dir.Name()
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

	return db
}

func initializedGakSoos(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM gak_soo").Scan(&count)
	if count >= 45 {
		log.Printf("⚠️ 45 Numbers is Already Exists")
		return
	}

	gakSooSQL, err := os.ReadFile(constants.InsertSQL + "i_gak_soo.sql")
	if err != nil {
		commons.BadSQLFile(err)
	}

	gakSooStatusSQL, err := os.ReadFile(constants.InsertSQL + "i_gak_soo_status.sql")
	if err != nil {
		commons.BadSQLFile(err)
	}

	tx, _ := db.Begin()
	defer tx.Rollback()

	for i := 1; i <= 45; i++ {
		gakSooID, _ := uuid.NewV7()
		statusID, _ := uuid.NewV7()
		tx.Exec(string(gakSooSQL), gakSooID.String(), i)
		tx.Exec(string(gakSooStatusSQL), statusID.String(), strconv.Itoa(i))
		log.Println("✅ Gak_Soo Insert Is Complete")
	}
	if err := tx.Commit(); err != nil {
		log.Fatalln("❌ Gak_Soo Insert Is UnComplete")
	}
}
