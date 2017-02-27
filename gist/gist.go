package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Server struct {
	logger *log.Logger
}

type Results []Result

type Result struct {
	FullName, URL, Description string
	Star                       int
}

func search(query, sort, order string) (Results, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/search/repositories", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)
	q.Set("sort", sort)
	q.Set("order", order)
	req.URL.RawQuery = q.Encode()

	var results Results
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		TotalCount int `json:"total_count"`
		Items      []struct {
			FullName   string `json:"full_name"`
			HTMLURL    string `json:"html_url"`
			Descripton string `json:"description"`
			Star       int    `json:"stargazers_count"`
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for _, res := range data.Items {
		results = append(results, Result{
			FullName:    res.FullName,
			URL:         res.HTMLURL,
			Description: res.Descripton,
			Star:        res.Star,
		})
	}
	return results, nil
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func (s *Server) ramenHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "q is empty", http.StatusBadRequest)
		return
	}
	results, err := search(query, "", "")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := resultsTemplate.Execute(w, struct {
		Results Results
	}{
		results,
	}); err != nil {
		log.Print(err)
		return
	}
}

func New() *Server {
	return &Server{
		logger: log.New(os.Stdout, "", 0),
	}
}

func main() {
	s := New()
	http.HandleFunc("/hi", s.handler)
	http.HandleFunc("/ramen", s.ramenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
<ol>
	{{range .Results}}
	<li>{{.Description}} - <a href="{{.URL}}">{{.URL}}</a></li>
	{{end}}
</ol>
</body>
</html>
`))
