package db

import (
	"fmt"
	"strconv"
)

func Insert(){
	for i := 10000; i <= 20000; i++{
		insert(i)
	}
}

func insert(index int){
	stm,err1 := GDB.Prepare("INSERT INTO tb_user(uid,nickname,chips,level,avatar,sex) values(?,?,?,?,?,?)")
	if err1 != nil{
		fmt.Println(err1.Error())
	}
	defer stm.Close()

	_,err := stm.Exec(index,"user"+strconv.Itoa(index),10000,1,"avatar",0)
	if err != nil{
		fmt.Println(err.Error())
	}
}

func Query(index string) ([]map[string]string,error) {
	sql := "select * from tb_user where uid="
	sql += index

	var results []map[string]string

	rows,err := GDB.Query(sql)
	defer rows.Close()

	if err != nil{
		return nil,err
	}else{
		cols,err := rows.Columns()
		if err != nil{
			return nil,err
		}
		vals := make([][]byte,len(cols))
		scans := make([]interface{},len(cols))
		for i := range vals{
			scans[i] = &vals[i]
		}


		for rows.Next(){
			err = rows.Scan(scans...)
			if err != nil{
				return nil,err
			}

			row := make(map[string]string)
			for k,v := range vals{
				key := cols[k]
				row[key] = string(v)
			}

			results = append(results,row)
		}

		for k,v := range results{
			fmt.Println(k,v)
			id := v["uid"]
			nn := v["nickname"]
			sex := v["sex"]
			fmt.Println(id,nn,sex)
		}
	}
	return results,nil
}

