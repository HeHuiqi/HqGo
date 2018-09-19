package hqdatabase
/*
安装mysql驱动
go get -u github.com/go-sql-driver/mysql
*/
import (
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func HqMysqlUse() {

	fmt.Println("-----------------------MySQLUse")
	//打开数据库 dataSourceName格式 user:password@/databaseName?charset=utf8
	db, err := sql.Open("mysql", "root:h12345678@/hhq")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}


	/*
	//插入数据

	inserSql := "insert userinfo set  username=?,departname=?,created=?"
	// insert into userInfo values()要占位所有字段值，然后传入所以字段值
	// inserSql := "insert into userInfo values(?,?,?,?)"
	  stmt, err := db.Prepare(inserSql)
	  checkErr(err)

	  res, err := stmt.Exec("孟子", "研发部门","1656-09-15")
	  checkErr(err)

	  id, err := res.LastInsertId()
	  checkErr(err)

	  fmt.Println("插入Id",id)
	  */



	//更新数据
	updateSql := "update userinfo set username = ? where uid =?"
	smt,err := db.Prepare(updateSql)
	checkErr(err)
	affect,err := smt.Exec("书法",1)
	checkErr(err)
	affectRows,err := affect.RowsAffected()
	fmt.Println("更新的行数",affectRows)
	/*
		//删除数据
		deleteSql := "delete from userinfo where uid = ? "
		smt, err = db.Prepare(deleteSql)
		checkErr(err)
		affect, err = smt.Exec(8)
		checkErr(err)
		affectRows, err = affect.RowsAffected()
		fmt.Println("删除影响的行数",affectRows)
		*/


	//查询数据
	rows, err := db.Query("select * from userInfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid,username,department,created)
	}

}
