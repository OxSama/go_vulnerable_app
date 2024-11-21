package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Test Application</title>
</head>
<body>
    <h1>Security Test Application</h1>
    <p>This is a test application with deliberately missing security headers.</p>
    <p style="color: red;">WARNING: This application is for local testing only!</p>
    
    <form method="GET" action="/search">
        <input type="text" name="q" placeholder="Enter search term">
        <input type="submit" value="Search">
    </form>

    {{if .Query}}
        <div>
            <h2>Search Results</h2>
            <p>You searched for: {{.Query}}</p>
        </div>
    {{end}}
</body>
</html>
`

type PageData struct {
	Query string
}

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	tmpl := template.Must(template.New("index").Parse(htmlTemplate))

	// Main page handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Search handler - deliberately missing security headers
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		data := PageData{
			Query: query,
		}
		tmpl.Execute(w, data)
	})

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting test server on http://localhost%s\n", addr)
	fmt.Println("WARNING: This is a test application with deliberate security issues.")
	fmt.Println("DO NOT deploy this in a production environment!")

	log.Fatal(http.ListenAndServe(addr, nil))
}
