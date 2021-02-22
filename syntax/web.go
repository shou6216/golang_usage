package syntax

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page3 struct {
	Title string
	Body  []byte
}

func (p *Page3) save() error {
	filename := "text/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page3, error) {
	filename := "./text/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page3{Title: title, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("view/edit.html", "view/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page3) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page3{Title: title}

	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page3{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func Web() {
	// ルーティングの設定
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// 第2引数はhandler。nilはdefault handler
	// default handlerは、/で遷移するとnotfoundエラーを返す
	// ListenAndServeはエラーメッセージを返す
	log.Fatal(http.ListenAndServe(":8080", nil))
}
