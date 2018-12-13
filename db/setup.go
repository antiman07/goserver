package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var GDB *sql.DB

func init(){
	db,err := sql.Open("mysql","root:root@/game")
	if err != nil{
		log.Fatalf("Open database error:%s\n",err)
	}

	GDB = db
	GDB.SetMaxOpenConns(2000)
	//gdb.SetMaxIdleConns(200)
}

func Release(){
	GDB.Close()
}

