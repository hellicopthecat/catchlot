package commons

import "log"

func BadSQLFile(err error) {
	log.Fatalf("❌ SQL 디렉토리 및 파일 읽기 실패 %s :: ", err)
}
