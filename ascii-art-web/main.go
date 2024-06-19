package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	// a handler function for handling requests
	// first param is a 'pattern' , second a handler
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {
// 	// shadow = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
// 	// standard = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
// 	// thinkertoy = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
// 	const tpl =`
// 	<html>

// 	<body>
// 	{{range .Items}}<div>{{.}}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
// 	</body>
// 	</html>
// 	`
// 	check := func (err error)  {
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	t, err := template.New("webpage").Parse(tpl)
// 	check(err)

// 	data1 := struct{
// 		Title string
// 		Items []string
// 	}{
// 		Title: "My page",
// 		Items: []string{
// 			"My photos",
// 			"My blog",
// 		},
// 	}

// 	noItems := struct{
// 		Title string
// 		Items []string
// 	}{
// 		Title: "My other page",
// 		Items: []string{},
// 	}

// 	err = t.Execute(os.Stdout, data1)
// 	check(err)
// 	err = t.Execute(os.Stdout, noItems)
// 	check(err)
// 	tmpl, err := template.New("example").Parse("<h1>Hello,{{.Name}}!</h>")
// 	check(err)
// 	data := struct {
// 		Name string
// 	}{
// 		"Fred",
// 	}
// 	err = tmpl.Execute(os.Stdout, data)
// 	check(err)

// }
