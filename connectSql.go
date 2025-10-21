package main

import (
	"database/sql"
	"fmt"
	"log"

	// นำเข้าไดรเวอร์สำหรับ SQL Server (ต้องติดตั้ง `go get github.com/go-sql-driver/mysql`)
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// สร้าง Data Source Name (DSN) สำหรับเชื่อมต่อฐานข้อมูล
	// รูปแบบ: user:password@tcp(host:port)/dbname
	// สำหรับ SQL Server: user:password@tcp(host:port)/dbname
	// (ตัวอย่างนี้ใช้ MySQL เนื่องจากเป็นตัวอย่างที่นิยม แต่หลักการเดียวกัน)
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"

	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: %v", err)
	}
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อว่าใช้งานได้หรือไม่
	err = db.Ping()
	if err != nil {
		log.Fatalf("ไม่สามารถเข้าถึงฐานข้อมูลได้: %v", err)
	}

	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จ!")
}
