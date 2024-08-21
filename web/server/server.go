package server

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

var (
	//go:embed tmpl
	templateFS embed.FS

	//go:embed static
	content embed.FS
)

type FSError struct {
	location string
	err      error
}

func (e FSError) Error() string {
	return "error creating file system rooted at '" + e.location + "': " + e.err.Error()
}

type WebServer struct {
	html *template.Template
	mux  http.ServeMux
}

func NewWebServer() (*WebServer, error) {
	staticFs, err := fs.Sub(content, "static")
	if err != nil {
		return nil, FSError{"static", err}
	}

	tmpl, err := parseTemplates(templateFS, ".html", true, nil)
	if err != nil {
		return nil, err
	}

	s := &WebServer{
		html: tmpl,
	}

	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFs))))
	s.mux.HandleFunc("/", httpHandler(s.handleIndex))

	return s, nil
}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	return s.html.ExecuteTemplate(w, "index.html", nil)
}

// parseTemplates recursively parses all templates in the FS with the given extension.
// File paths are used as template names to support duplicate file names.
// Use nonRootTemplateNames to exclude root directory from template names
// (e.g. index.html instead of templates/index.html)
func parseTemplates(
	templates fs.FS,
	ext string,
	nonRootTemplateNames bool,
	funcMap template.FuncMap) (*template.Template, error) {

	root := template.New("")
	err := fs.WalkDir(templates, "tmpl", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ext) {
			if err != nil {
				return err
			}
			b, err := fs.ReadFile(templates, path)
			if err != nil {
				return err
			}
			name := ""
			if nonRootTemplateNames {
				//name the template based on the file path (excluding the root)
				parts := strings.Split(path, string(os.PathSeparator))
				name = strings.Join(parts[1:], string(os.PathSeparator))
			}
			t := root.New(name).Funcs(funcMap)
			_, err = t.Parse(string(b))
			if err != nil {
				return err
			}
		}
		return nil
	})

	return root, err
}
