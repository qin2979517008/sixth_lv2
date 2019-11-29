package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var name string
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/students?charset=utf8")
	if err != nil {
		fmt.Println("打开数据库失败")
	}else{
		fmt.Println("打开数据库！")
	}
	defer db.Close()
	defer fmt.Println("退出")
	for i := 2019211530; i <= 2019211565; i++ {
		a := strconv.Itoa(i)
		website := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + a
		res, err := http.Get(website)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 读取资源数据 body: []byte
		body, err := ioutil.ReadAll(res.Body)
		// 关闭资源流
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", website, err)
			os.Exit(1)
		}
		//写入文件
		ioutil.WriteFile("site.txt", body, 0644)
		file, _ := ioutil.ReadFile("site.txt")
		regall := `>>\d{10}([\s\S]*?)<`
		rega := regexp.MustCompile(regall)
		all := rega.FindAllStringSubmatch(string(file), 1)
		for _, v := range all {
			if v [1] != "  " {
				name = v[1]
				fmt.Println(name)
			}
		}
		insert(db,name,i)
	}
}
func insert(db *sql.DB , name string ,i int )  {
	_,err := db.Exec("INSERT into class(id,name) values (?,?)",i,name)
	if (err !=nil){
		fmt.Println(err)
	}else{
		fmt.Println("插入成功")
	}
}