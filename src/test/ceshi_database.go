package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "l123456.."
	dbname   = "testdb"
)

func main() {
	db := getDB()
	if db != nil {
		SelectUser(db)
		DeleteUser(db)

	}
}

func getDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("successfull connected!")
	return db
}

type User struct {
	username string
	phone    int
	cryptpwd string
	md5pwd   string
}

//查询
func SelectUser(db *sql.DB) {
	rows, err := db.Query("select * from testusers")
	if err != nil {
		log.Fatal(err)
	}
	user := User{}
	for rows.Next() {
		err := rows.Scan(&user.username, &user.phone, &user.cryptpwd, &user.md5pwd)
		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(user)

}

//添加用户
func InsertUser(db *sql.DB) {
	stmt, err := db.Prepare("insert into testusers(username,phone,cryptpwd,md5pwd) values($1,$2,$3,$4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(6, "Windows", 1, "操作系统")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into user_tbl success")
	}
}

func InsertUser2(db *sql.DB) {
	db.QueryRow("INSERT INTO user_tbl(id,username,sex,info) VALUES ($1,$2,$3,$4)", 7, "Linux", 0, "dd")
}

//删除用户
func DeleteUser(db *sql.DB) {
	stmt, err := db.Prepare("DELETE  FROM testusers WHERE  username=$1")
	fmt.Println(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(6)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("delete form user_tbl success")
	}
}

//更新用户
func UpdateUser(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE  user_tbl  set username=$1 WHERE  id=$2")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("Linux", 6)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("udpate user_tbl success")
	}

}
