package web

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

const (
	TemplateExt  string = ".html"
	TemplatesDir string = "tmpl"

	RootTmpl  string = ""
	IndexTmpl string = "index.html"

	NonRootTmplNames bool = true
)

var (
	//go:embed tmpl
	templateFS embed.FS

	//go:embed static
	contentFS embed.FS
)

type FSError struct {
	dir string
	err error
}

func (e *FSError) Error() string {
	return "error creating file system rooted at " + e.dir + ": " + e.err.Error()
}

type ServerHandler func(http.ResponseWriter, *http.Request) error

func httpHandler(fn ServerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type Server struct {
	mux  http.ServeMux
	html *template.Template
}

// NewServer creates a new WebServer
func NewServer(ctx context.Context) (*Server, error) {
	staticFS, err := fs.Sub(contentFS, "static")
	if err != nil {
		return nil, &FSError{"static", err}
	}

	s := &Server{}

	if err := s.parseTemplates(templateFS, TemplateExt, nil); err != nil {
		return nil, err
	}

	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	s.mux.HandleFunc("/", httpHandler(s.handleIndex))
	s.mux.HandleFunc("/dashboard", httpHandler(s.handleDashboard))
	s.mux.HandleFunc("/tables", httpHandler(s.handleTables))
	s.mux.HandleFunc("/billing", httpHandler(s.handleBilling))
	s.mux.HandleFunc("/profile", httpHandler(s.handleProfile))
	s.mux.HandleFunc("/login", httpHandler(s.handleLogin))
	s.mux.HandleFunc("/registration", httpHandler(s.handleRegistration))

	return s, nil
}

// ServeHTTP calls the underlying mux.ServeHTTP method
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/index recieved request\n")

	return s.html.ExecuteTemplate(w, IndexTmpl, nil)
}

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/dashboard recieved request\n")

	return s.html.ExecuteTemplate(w, "dashboard", nil)
}

func (s *Server) handleTables(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/tables recieved request\n")

	return s.html.ExecuteTemplate(w, "tables", nil)
}

func (s *Server) handleBilling(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/billing recieved request\n")

	return s.html.ExecuteTemplate(w, "billing", nil)
}

func (s *Server) handleProfile(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/profile recieved request\n")

	return s.html.ExecuteTemplate(w, "profile", nil)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/login recieved request\n")

	return s.html.ExecuteTemplate(w, "login", nil)
}

func (s *Server) handleRegistration(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(os.Stdout, "/registration recieved request\n")

	return s.html.ExecuteTemplate(w, "registration", nil)
}

func (s *Server) parseTemplates(templates fs.FS, ext string, funcMap template.FuncMap) error {
	root := template.New(RootTmpl)
	walkDir := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(path, ext) {
			return nil
		}

		b, err := fs.ReadFile(templates, path)
		if err != nil {
			return err
		}

		var name string
		name = path

		if NonRootTmplNames {
			parts := strings.Split(path, string(os.PathSeparator))
			name = strings.Join(parts[1:], string(os.PathSeparator))
		}

		t := root.New(name).Funcs(funcMap)

		_, err = t.Parse(string(b))
		if err != nil {
			return err
		}

		return nil
	}

	if err := fs.WalkDir(templates, TemplatesDir, walkDir); err != nil {
		return err
	}

	s.html = root

	return nil
}
