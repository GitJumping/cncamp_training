package main

import (
	"html/template"
	"log"
	"net/http"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

func main() {

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "我的第一个HTML页面",
		Items: []string{
			"我的相册",
			"我的博客",
		},
	}

	http.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		err = t.Execute(writer, data)
		check(err)
	})
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
