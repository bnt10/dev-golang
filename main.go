package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// 환경 변수에서 데이터베이스 연결 정보 가져오기
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 환경 변수 출력 (디버깅 목적)
	fmt.Printf("host=%s user=%s password=%s dbname=%s sslmode=disable\n",
		dbHost, dbUser, dbPassword, dbName)

	// 데이터베이스 연결 문자열 구성
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)

	// 데이터베이스 연결
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 연결 확인
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	// 간단한 쿼리 실행 (users 테이블 조회)
	rows, err := db.Query("SELECT id, username, email, created_at, last_login FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 조회 결과 출력
	for rows.Next() {
		var id int
		var username, email string
		var createdAt, lastLogin sql.NullTime
		err := rows.Scan(&id, &username, &email, &createdAt, &lastLogin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s, CreatedAt: %v, LastLogin: %v\n",
			id, username, email, createdAt, lastLogin)
	}

	// 에러 확인
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
