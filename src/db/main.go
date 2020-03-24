package main

import (
	"database/sql"
	"db/database_source"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func ConnectMysql(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed,detail error:%s", err.Error())
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed,detail error:%s", err.Error())
		return nil
	}
	return db
}

func ExecuRecord(Db *sql.DB, ExecuSql string) error {
	result, err := Db.Exec(ExecuSql)
	if err != nil {
		fmt.Printf("%s occured err:%s", ExecuSql, err.Error())
		return err
	}
	num, _ := result.RowsAffected()
	fmt.Printf("insert %d rows success", num)
	return nil
}

func QuerySingleRecord(Db *sql.DB, QuerySql string) error {
	var id, name, age, sex string
	row := Db.QueryRow(QuerySql)
	err := row.Scan(&id, &name, &age, &sex)
	if err != nil {
		fmt.Printf("Scan error:%s", err.Error())
		return err
	}
	fmt.Printf("id:%s,name:%s,age:%s,sex:%s", id, name, age, sex)
	return nil
}

func QueryMultiRecord(Db *sql.DB, QuerySql string) error {
	rows, err := Db.Query(QuerySql)
	if err != nil {
		fmt.Printf("query rows error:%s", err.Error())
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name, age, sex string
		err = rows.Scan(&id, &name, &age, &sex)
		if err != nil {
			fmt.Printf("Scan error:%s", err.Error())
			return err
		}
		fmt.Printf("id:%s,name:%s,age:%s,sex:%s\n", id, name, age, sex)
	}
	return nil
}

func main() {
	source := database_source.DataSourceName{
		Username: "root",
		Password: "a8016577A",
		Ip:       "127.0.0.1",
		Port:     3306,
		Database: "gjchen",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		source.Username,
		source.Password,
		source.Ip,
		source.Port,
		source.Database)
	db := ConnectMysql(dsn)
	if db != nil {
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(16)
		db.SetConnMaxLifetime(100 * time.Second)
		QueryMultiRecord(db, "select * from student;")
		ExecuRecord(db, "insert into student (id,name,age,sex) values (7,'tom',44,1);")
		QueryMultiRecord(db, "select * from student;")
	}
}
