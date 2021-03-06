package main

import (
	"fmt"
	"github.com/hefju/GoDream/model"
	"html/template"
	"net/http"
	"strconv"
    "encoding/json"
    "os"
//    "io/ioutil"
    "log"
)


func main() {
    http.Handle("/js/", http.FileServer(http.Dir("web")))
	http.Handle("/css/", http.FileServer(http.Dir("web")))
	http.Handle("/imgs/", http.FileServer(http.Dir("web")))

	//http.Handle("/", http.StripPrefix("/template/", http.FileServer(http.Dir("/template"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/setting", setting)
	http.HandleFunc("/about", about)
	http.HandleFunc("/updatesetting", updatesetting)
	http.HandleFunc("/addsetting", addsetting)
	http.HandleFunc("/deletesetting", deleteSetting)
	http.HandleFunc("/addUpdateLog", addUpdateLog)
	http.HandleFunc("/task",task)
    http.HandleFunc("/addtask",addtask)
    http.HandleFunc("/gettask",gettask)
    http.HandleFunc("/updateTask",updateTask)

    http.HandleFunc("/msg",message)
    http.HandleFunc("/test",test)

    addr:=GetDefaultListenInfo()
	err := http.ListenAndServe(addr, nil)//":9000"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cao")
}

func test(w http.ResponseWriter, r *http.Request) {
    render(w, "web/test.html", nil)
}

func task(w http.ResponseWriter, r *http.Request) {
    fmt.Println("func task: ")
    date:=""
    lst := model.GetTask(date)
    locals := make(map[string]interface{})
    locals["mytasks"] = lst
	render(w, "web/task.html", locals)
}
func updateTask(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("func updateTask: ")
    id,err:=strconv.ParseInt(r.FormValue("txtID"),10,64)
    if err!=nil{
        fmt.Println(err)
    }

    t:=new(model.MyTask)
    t.Id=id
    t.Title=r.FormValue("txtTitle")
    t.Content=r.FormValue("txtContent")
    fmt.Println(t,id)
    var result string
    _,err=model.UpdateTask(t)
    if err!=nil{
        fmt.Println(err)
        result="执行失败."

    }else{
        result="执行成功."
    }
fmt.Println(result)
    js, err := json.Marshal(result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
func gettask(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    strID:= r.FormValue("name1")
    id,err:=strconv.ParseInt(strID,10,64)
    if err!=nil{
        fmt.Println(err)
        return
    }
    //fmt.Println("func gettask: ",strID)
    t:=model.GetTaskByID(id)
    js, err := json.Marshal(t)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println(t)
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
//添加任务
func addtask(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    id,_:=strconv.ParseInt("0",10,64)
    newtask:=r.FormValue("txtInputTask")
    fmt.Println("func addtask: ",newtask)
    t:=&model.MyTask{Id:id,Title:newtask}
   affect,_:= model.Insert(t)
    fmt.Println(affect)
    task(w,r)
}


func index(w http.ResponseWriter, r *http.Request) {
	lst := model.GetAllUpdateLog()
	locals := make(map[string]interface{})
	locals["UpdateLog"] = lst
	locals["count"]=0

	render(w, "web/index.html", locals)
}

func about(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("func about")
	//render(w, "web/about.html", nil)
	t, _ := template.ParseFiles("web/about.html", "web/tmpl/navbar.tmpl")
	//t.Execute(os.Stdout, nil)
	t.Execute(w, nil)
	//t.ExecuteTemplate(w,nil)
}
func setting(w http.ResponseWriter, r *http.Request) {
	lst := model.GetAllSetting()
	locals := make(map[string]interface{})
	locals["settings"] = lst
	locals["count"]=0
	render(w, "web/setting.html", locals)
}
func updatesetting(w http.ResponseWriter,r *http.Request){
	fmt.Println("func updatesetting")
		r.ParseForm()
		id,err:=strconv.ParseInt(r.FormValue("Id"),10,64)
	if err!=nil{
		fmt.Println("id not int")
		return
	}
		keyname:=r.FormValue("key")
		keyvalue:=r.FormValue("value")
	set:=&model.Mysetting{Id:id,KeyName:keyname,KeyValue:keyvalue}
	count:=set.Save()
	fmt.Println(count,keyname,keyvalue)
	setting(w,r)
}
func addsetting(w http.ResponseWriter,r *http.Request) {
	//fmt.Println("func addsetting ")
	r.ParseForm()
	id,err:=strconv.ParseInt("0",10,64)
	if err!=nil{
		fmt.Println("id not int")
		return
	}
	keyname:=r.FormValue("txtName")
	keyvalue:=r.FormValue("txtValue")
	fmt.Println("func addsetting ",keyname,keyvalue)
	set:=&model.Mysetting{Id:id,KeyName:keyname,KeyValue:keyvalue}
	count:=set.Save()
	fmt.Println(count)
	setting(w,r)
}
func deleteSetting(w http.ResponseWriter,r *http.Request) {
	fmt.Println("func deleteSetting ")
	r.ParseForm()
	id,err:=strconv.ParseInt(r.FormValue("Id"),10,64)
	if err!=nil{
		fmt.Println("id not int")
		return
	}

	set:=&model.Mysetting{Id:id}
	count:=set.Delete()
	fmt.Println(count)
	setting(w,r)
}
func addUpdateLog(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	info:=new(model.UpdateLog)
	info.Title=r.FormValue("txtTitle")
	info.Content=r.FormValue("txtContent")

	count:=info.Save()
	fmt.Println(count)
	index(w,r)
}

//接收来自客户的信息
func message(w http.ResponseWriter, r *http.Request) {
    if r.Method=="POST"{
        decoder := json.NewDecoder(r.Body)
        var msg  model.ReportMsg
        err := decoder.Decode(&msg)
        if err != nil {
            log.Println(err)  // panic()
        }
        model.Insert(msg)
        fmt.Println(msg)
        //        body, _ := ioutil.ReadAll(r.Body)
        //        fmt.Println(string(body))
        fmt.Fprintln(w,"200")

    }else{
        fmt.Fprintln(w,"server receive a message.")
    }
}
func render(w http.ResponseWriter, tmplName string, context map[string]interface{}) {
	//tmpl, err := template.ParseFiles(tmplName)
	tmpl, err := template.ParseFiles(tmplName, "web/tmpl/navbar.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//tmpl.Execute(os.Stdout, nil)
	tmpl.Execute(w, context)
	//tmpl.ExecuteTemplate(w, context)
	return
}

func GetDefaultListenInfo() string {
    host := os.Getenv("HOST")
    if len(host) == 0 {
        host = "0.0.0.0"
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "9000"
    }
    return host + ":" + port
}

