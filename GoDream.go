package main

import (
	"fmt"
	"godream/model"
	"html/template"
	"net/http"
	//"os"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	//	r.ParseForm()
	//	for k,v:=range r.Form{
	//		fmt.Println(k,":",v)
	//	}

	locals := make(map[string]interface{})
	locals["userName"] = "caocao"
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

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("web")))
	http.Handle("/imgs/", http.FileServer(http.Dir("web")))

	//http.Handle("/", http.StripPrefix("/template/", http.FileServer(http.Dir("/template"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/setting", setting)
	http.HandleFunc("/about", about)
	http.HandleFunc("/updatesetting", updatesetting)
	http.HandleFunc("/addsetting", addsetting)
	http.HandleFunc("/deletesetting", deleteSetting)

	//http.FileServer(http.Dir("web"))

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cao")
}
