package server

import (
	"context"
	"embed"
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

type WebServer struct {
	mux  http.ServeMux
	html *template.Template
}

// NewWebServer creates a new WebServer
func NewWebServer(ctx context.Context) (*WebServer, error) {
	staticFS, err := fs.Sub(contentFS, "static")
	if err != nil {
		return nil, &FSError{"static", err}
	}

	s := &WebServer{}

	if err := s.parseTemplates(templateFS, TemplateExt, nil); err != nil {
		return nil, err
	}

	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
	s.mux.HandleFunc("/", httpHandler(s.handleIndex))

	return s, nil
}

// ServeHTTP calls the underlying mux.ServeHTTP method
func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	return s.html.ExecuteTemplate(w, IndexTmpl, nil)
}

func (s *WebServer) parseTemplates(templates fs.FS, ext string, funcMap template.FuncMap) error {
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
