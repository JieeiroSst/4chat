package db

import (
	"404Chat/controller"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	NameDB    = "mysql"
	MysqlUser = "root"
	MysqlPass = " "
	MysqlDB   = "4chat"
	MyPort    = "3306"
	MyUrl     = "localhost"
)

var (
	db        *sql.DB
	accocunts []controller.Account
)

func HandleDB() {
	db, err := sql.Open(NameDB, MysqlUser+":"+MysqlPass+"@("+MyUrl+":"+MyPort+")/"+MysqlDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mysql database")
}

func GetList() []controller.Account {
	var account controller.Account
	result, err := db.Query("SELECT* from `Account`")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	accocunts = append(accocunts, account)

	return accocunts
}

func SelectUser(id string) {
	var result, err = db.Query("SELECT UserName,Password from `Account` where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	UserNames := make([]string, 0)
	Passwords := make([]string, 0)
	for result.Next() {
		var UserName string
		var Password string
		err := result.Scan(&UserName, &Password)
		if err != nil {
			log.Fatal(err)
		}
		UserNames = append(UserNames, UserName)
		Passwords = append(Passwords, Password)
	}

	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
}

type project struct {
	UserName string
	Password string
}

var projects []project

func InsertUser(UserName string, Password string) {
	stmt, err := db.Prepare("insert into `Account`(id,Username,Passowrd) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	//uppdate data
	var p project
	p.UserName = UserName
	p.Password = Password

	for id, user := range projects {
		if err, _ := stmt.Exec(id+1, user.UserName, user.Password); err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteUser(id string) {
	result, err := db.Query("delete from `Account` where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
}
