package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

// Result ...
type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

type ajaxController struct {
}

func (p *ajaxController) LoginAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := r.ParseForm()
	if err != nil {
		outPutJSON(w, 0, "参数错误", nil)
		return
	}

	adminName := r.FormValue("admin_name")
	adminPassword := r.FormValue("admin_password")

	if adminName == "" || adminPassword == "" {
		outPutJSON(w, 0, "参数错误", nil)
		return
	}

	db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "123456", "webdemo")
	if err := db.Connect(); err != nil {
		log.Println(err)
		outPutJSON(w, 0, "连接数据库失败", nil)
		return
	}
	defer db.Close()

	rows, res, err := db.Query("select * from webdemo_admin where admin_name = '%s'", adminName)
	if err != nil {
		log.Println(err)
		outPutJSON(w, 0, "查询数据库失败", nil)
		return
	}

	index := res.Map("admin_password")
	adminPasswordDB := rows[0].Str(index)

	if adminPasswordDB != adminPassword {
		outPutJSON(w, 0, "密码输入错误", nil)
		return
	}

	// 存入cookie,使用cookie存储
	cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
	http.SetCookie(w, &cookie)

	outPutJSON(w, 1, "操作成功", nil)
	return
}

func outPutJSON(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}
