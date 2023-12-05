package server

import (
	_ "embed"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
)

var (
	//go:embed index.html
	indexhtml string
	indexTmpl = template.Must(template.New("index").Parse(indexhtml))
)

type Introduce struct {
	Name string
	Age  string
}

type Server struct {
	mux        *http.ServeMux
	httpserver *http.Server
	template   *template.Template
}

func New(addr string) *Server {
	mux := http.NewServeMux()
	s := &Server{
		mux: mux,
		httpserver: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		template: indexTmpl,
	}
	s.init()
	return s
}

func (s *Server) ListenAndServe() error {
	return s.httpserver.ListenAndServe()
}

func (s *Server) init() {
	s.mux.HandleFunc("/", s.handleIndex)
	s.mux.HandleFunc("/self", s.handleSelf)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	data := Introduce{
		Name: "Yamato",
		Age:  "21",
	}

	if err := s.template.Execute(w, data); err != nil {
		slog.Error("error", "err", err)
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}
}

func (s *Server) handleSelf(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if name == "" || age == "" {
		slog.Error("error", "err", errors.New("name or age is empty"))
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status), status)
		return
	}

	data := Introduce{
		Name: name,
		Age:  age,
	}

	if err := s.template.Execute(w, data); err != nil {
		slog.Error("error", "err", err)
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}
}
