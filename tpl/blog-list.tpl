<h1>Blog List</h1>
{{range .}}
    <div><a href="\page\{{.GUID}}">{{.Title}}</a></div>
    <div>{{.Content}}</div>
    <div>{{.Date}}</div>
    <br>
{{end}}
