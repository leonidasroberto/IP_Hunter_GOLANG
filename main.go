package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

////Configurando template
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./Templates/*.html"))
}

////Configurando template fim
func request(w http.ResponseWriter, r *http.Request) {

	///Redirecionar caso o metodo não seja POST
	if r.Method != "POST" || r.FormValue("ip") == "" || strings.Contains(r.FormValue("ip"), ".") == false {
		http.Redirect(w, r, "/erro", http.StatusSeeOther)
		return
	}

	///Objeto que recebe o json
	type jason struct {
		Country_name string
		Region_name  string
		Country_code string
		Ip           string
		City         string
		Zip          string
		Latitude     float64
		Longitude    float64
		Country_flag string
	}

	///Pegando ip do formulario
	fip := r.FormValue("ip")

	///Tratamento Request
	req, _ := http.NewRequest("GET", "http://api.ipapi.com/"+fip+"?access_key=cd975b4561325b163e847342ecad31f8", nil)
	var client http.Client
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	///Passando objeto struct para variavel e convertendo Json pro Objeto
	var t jason
	json.Unmarshal([]byte(body), &t)

	//Escrevendo resposta
	tpl.ExecuteTemplate(w, "resul.html", t)

}

func consulta(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "form.html", nil)
}

func erro(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "erro.html", nil)
}

///Configuração endereços e função, e porta
func main() {
	http.HandleFunc("/request", request)
	http.HandleFunc("/", consulta)
	http.HandleFunc("/erro", erro)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
