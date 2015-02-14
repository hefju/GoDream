package main
import (
	"fmt"
	"net/http"
	"html/template"
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

func render(w http.ResponseWriter, tmplName string, context map[string]interface{}) {
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, context)
	return
}

func main(){
	http.HandleFunc("/", index)

	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("cao")
}
