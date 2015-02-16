package main
import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

func index(w http.ResponseWriter,r *http.Request){
//	r.ParseForm()
//	for k,v:=range r.Form{
//		fmt.Println(k,":",v)
//	}

	locals := make(map[string]interface{})
	locals["userName"]="caocao"
	render(w, "web/index.html", locals)
}

func about(w http.ResponseWriter,r *http.Request){
	//fmt.Println("func about")
	//render(w, "web/about.html", nil)
	t, _ := template.ParseFiles("web/about.html", "web/block.tmpl")
	//t.Execute(os.Stdout, nil)
	t.Execute(w, nil)
  //t.ExecuteTemplate(w,nil)
}

func render(w http.ResponseWriter, tmplName string, context map[string]interface{}) {
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//tmpl.Execute(os.Stdout, nil)
	tmpl.Execute(w, context)
	//tmpl.ExecuteTemplate(w, context)
	return
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("cao")
}
