package contollers

import (
	"Golang6/src/Todo_App/app/models"
	"Golang6/src/Todo_App/config"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("./app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error){
	cookie, err := r.Cookie("_cookie")
	if err == nil{
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid Session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9])+$")

func parseURL(fn func(w http.ResponseWriter, r *http.Request, i int)) (http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request){
		// todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil{
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil{
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)

	}
}

func StartMainServer() error{
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("todos/save", todoSave)
	http.HandleFunc("todos/edit/", parseURL(todoEdit))
	http.HandleFunc("todos/update", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete", parseURL(todoDelete))
	return http.ListenAndServe(":" + config.Config.Port, nil)

}
