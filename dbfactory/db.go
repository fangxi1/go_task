package dbfactory

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connection_db() *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("连接数据库失败=", err)
		return nil
	}

	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库不可用=", err)
		return nil
	}
	fmt.Println("数据库连接成功")
	return db

}

func DbTest1(db *sql.DB) {
	// 新增
	_, err := db.Exec("insert into students(name,age,grade) values (?,?,?)", "张三", 16, "三年级")
	if err != nil {
		fmt.Println("插入数据异常=", err)

	}
	// 新增
	_, err = db.Exec("insert into students(name,age,grade) values (?,?,?)", "李四", 17, "五年级")
	if err != nil {
		fmt.Println("插入数据异常=", err)

	}
	_, err = db.Exec("update students set age = ? where name=?", 14, "张三")
	if err != nil {
		fmt.Println("修改数据异常=", err)

	}

	rows, err := db.Query("select name,age,grade from students where age >?", "18")
	if err != nil {
		fmt.Println("查询数据异常=", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var age int
		var grade string
		err := rows.Scan(&name, &age, &grade)
		if err != nil {
			fmt.Println("读取数据异常=", err)
			continue
		}
		fmt.Printf("姓名: %s, 年龄: %d, 年级: %s\n", name, age, grade)
	}

	_, err = db.Exec("delete from students where age<?", 15)
	if err != nil {
		fmt.Println("删除数据异常=", err)
	}
}

func DbTest2(db *sql.DB) {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("开启事务失败=", err)
		return
	}
	_, err = tx.Exec("insert into accounts(name,balance) values (?,?)", "A", 200)
	if err != nil {
		fmt.Println("新增数据失败=", err)
		return
	}
	_, err = tx.Exec("insert into accounts(name,balance) values (?,?)", "B", 200)
	if err != nil {
		fmt.Println("新增数据失败=", err)
		return
	}
	var idA uint32
	var balanceA float64
	var amount float64 = 100
	err = tx.QueryRow("select id,balance from accounts where name=?", "A").Scan(&idA, &balanceA)
	if err != nil {
		fmt.Println("查询数据A失败=", err)
		return
	}
	if balanceA < float64(100) {
		fmt.Println("账户A余额不足=", balanceA)
		tx.Rollback()
	}
	balanceA -= amount
	var idB uint32
	var balanceB float64
	err = tx.QueryRow("select id,balance from accounts where name=?", "B").Scan(&idB, &balanceB)
	if err != nil {
		fmt.Println("查询数据B失败=", err)
		return
	}
	balanceA += amount

	_, err = tx.Exec("update accounts set balance = ? where id = ?", balanceA, idA)
	if err != nil {
		fmt.Println("更新数据A余额失败=", err)
		tx.Rollback()
	}
	_, err = tx.Exec("update accounts set balance = ? where id = ?", balanceB, idB)
	if err != nil {
		fmt.Println("更新数据B余额失败=", err)
		tx.Rollback()
	}
	_, err = tx.Exec("insert into transactions(from_account_id,to_account_id,amount) values (?,?,?)", idA, idB, amount)
	if err != nil {
		fmt.Println("新增交易记录失败=", err)
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交事务失败=", err)
	}
	fmt.Println("转账成功")

}
