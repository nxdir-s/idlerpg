package web

import (
	"context"
	"embed"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	mux    http.ServeMux
	html   *template.Template
	logger *slog.Logger
}

// NewServer creates a new WebServer
func NewServer(ctx context.Context, logger *slog.Logger) (*Server, error) {
	staticFS, err := fs.Sub(contentFS, "static")
	if err != nil {
		return nil, &FSError{"static", err}
	}

	s := &Server{
		logger: logger,
	}

	if err := s.parseTemplates(templateFS, TemplateExt, nil); err != nil {
		return nil, err
	}

	s.mux.Handle("/static/", otelhttp.NewHandler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))), "web.static"))

	s.mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleIndex)), "web.index"))
	s.mux.Handle("/dashboard", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleDashboard)), "web.dashboard"))
	s.mux.Handle("/tables", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleTables)), "web.tables"))
	s.mux.Handle("/billing", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleBilling)), "web.billing"))
	s.mux.Handle("/profile", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleProfile)), "web.profile"))
	s.mux.Handle("/login", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleLogin)), "web.login"))
	s.mux.Handle("/registration", otelhttp.NewHandler(http.HandlerFunc(httpHandler(s.handleRegistration)), "web.registration"))

	return s, nil
}

// ServeHTTP calls the underlying mux.ServeHTTP method
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/index recieved request")

	return s.html.ExecuteTemplate(w, IndexTmpl, nil)
}

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/dashboard recieved request")

	return s.html.ExecuteTemplate(w, "dashboard", nil)
}

func (s *Server) handleTables(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/tables recieved request")

	return s.html.ExecuteTemplate(w, "tables", nil)
}

func (s *Server) handleBilling(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/billing recieved request")

	return s.html.ExecuteTemplate(w, "billing", nil)
}

func (s *Server) handleProfile(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/profile recieved request")

	return s.html.ExecuteTemplate(w, "profile", nil)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/login recieved request")

	return s.html.ExecuteTemplate(w, "login", nil)
}

func (s *Server) handleRegistration(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	s.logger.Info("/registration recieved request")

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
