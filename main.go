package main

import (
	"html/template"
	"log"
	"net/http"

	c "random/internal/connection"
	v "random/internal/verification"
)

var tpl *template.Template
var Getphone string

type PageData struct {
	Title         string
	FirstName     string
	Phone         string
	IsAuth        bool
	Verification  bool
	International string
	Country       string
	WrongNumber   bool
	WrongMessage  string
}

// type SearchModel interface {
// 	searching
// }

// type Search struct{}

// type brazil struct{}

// func (s *Search) searching(name string) interface{} {
// 	return true
// }

// func SearchFromDB(s SearchModel, name string) bool {
// 	return true
// }

func init() {

	pd := PageData{}

	pd.Title = "Index Page"
	pd.IsAuth = false
	pd.FirstName = "Путник"

	c.Connection() // database

	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/status", status)
	http.Handle("/templates/", http.FileServer(http.Dir("./")))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// / позже вынести на index package
func index(w http.ResponseWriter, req *http.Request) {

	pd := PageData{
		Title: "Index Page",
	}
	pd.FirstName = "Путник"

	var (
		name  string
		phone string
	)
	/// Если метод выполнен (форма заполнена), то начинается проверка
	if req.Method == http.MethodPost {
		name = req.FormValue("fname")
		phone = req.FormValue("phone")

		Getphone = phone
		var (
			international string
			verified      bool
			country       string
			auth          bool
		)

		//// поправить nesting надо бы
		auth = v.MustCompileNumber(Getphone)
		if !auth {
			pd.IsAuth = false
			pd.WrongNumber = true
			pd.WrongMessage = "wrong number. please, try again."
		} else {
			v.IsNumberCorrect(Getphone)
			international, verified = v.Verification(Getphone)
			if !verified {
				pd.IsAuth = false
				pd.WrongNumber = true
				pd.WrongMessage = "the verification is not passed."
			} else {
				country = v.GetCountryName(Getphone)
				pd.IsAuth = true
				pd.FirstName = name
				pd.Phone = phone
				pd.International = international
				pd.Verification = verified
				pd.Country = country

				c.InsertOne(pd) // в базу из connection
			}
		}
	}

	err := tpl.ExecuteTemplate(w, "index.html", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	//// запрос и вывод на эту страницу с формы на index
}
