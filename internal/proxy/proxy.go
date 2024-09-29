package proxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"

	"github.com/golanguzb70/stream-cdn/config"
)

type Proxy struct {
	Config config.Config
}

func NewProxy() Proxy {
	return Proxy{
		Config: config.New(),
	}
}

func (p *Proxy) ReverseProxy(target string) http.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		fmt.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Modify the director to properly forward the request URL
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = url.Path + req.URL.Path
		req.Host = url.Host
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)

		rec := httptest.NewRecorder()
		proxy.ServeHTTP(rec, r)

		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
	}
}

func (p *Proxy) Start() error {
	mux := http.NewServeMux()

	// Wrap the reverse proxy handler with the CORS middleware
	mux.HandleFunc("/", http.HandlerFunc(p.ReverseProxy(p.Config.OriginServerURL)))

	// Start the server
	return http.ListenAndServe(fmt.Sprintf(":%d", p.Config.Port), p.Middleware(mux))
}

func (p *Proxy) HasAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (p *Proxy) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (OPTIONS method)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range p.Config.AllowedOrigins {
			if origin == allowedOrigin || allowedOrigin == "*" {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
