<h1>Blog List</h1>
{{range .}}
    <div><a href="\page\{{.GUID}}">{{.Title}}</a></div>
    <div>{{.TruncatedContent}}</div>
    <div>{{.Date}}</div>
    <br>
{{end}}
