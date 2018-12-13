package db

import (
	"fmt"
	"strconv"
	"time"
)

func insert(index int){

	sql := "insert into t_user(user_id,acct_name)VALUES"
	onedata := "(" + strconv.Itoa(index)  + ", '20180103002930')"

	_,err := GDB.Exec(sql + onedata)
	if err != nil{
		panic(err.Error())
	}
}

func insert1(index int){
	stm,err1 := GDB.Prepare("INSERT INTO t_user(user_id,acct_name) values(?,?)")
	if err1 != nil{
		fmt.Println(err1.Error())
	}
	defer stm.Close()

	_,err := stm.Exec(index,"user"+strconv.Itoa(index))
	if err != nil{
		fmt.Println(err.Error())
	}
}
func Create_db_data(){
	for i := 1; i <= 10000; i++{
		go insert2(i)
		time.Sleep(time.Millisecond * 10)
	}
}

func insert2(index int){
	tx,_ := GDB.Begin()

	defer tx.Commit()

	_,err := tx.Exec("INSERT INTO t_user(user_id,acct_name) values(?,?)",index,"user"+strconv.Itoa(index))
	if err != nil{
		fmt.Println(err.Error())
	}
}

