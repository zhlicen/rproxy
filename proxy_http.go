package rproxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/http2"
)

var webRequestTotal *prometheus.CounterVec
var webRequestDuration *prometheus.HistogramVec
var ph http.Handler

func init() {
	webRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "web_request_total",
			Help: "Number of requests in total",
		},
		[]string{"method", "endpoint"},
	)

	webRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "web_request_duration_seconds",
			Help:    "web request duration distribution",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1},
		},
		[]string{"method", "endpoint"},
	)
	prometheus.MustRegister(webRequestTotal)
	prometheus.MustRegister(webRequestDuration)
	ph = promhttp.Handler()

}

type proxyHandler struct {
}

func matchPrefix(prefix string, uri string) (relativePath string, match bool) {
	if prefix == "/" {
		relativePath = uri
		return relativePath, true
	}
	if strings.HasPrefix(uri, prefix) {
		relativePath = strings.TrimPrefix(uri, prefix)
		if strings.HasPrefix(relativePath, "/") || relativePath == "" {
			return relativePath, true
		}
	}
	return "", false

}

func (p *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// grpc-web uses http 1.1
	if serveGRPCWeb(w, r) {
		return
	}
	start := time.Now()
	log.Printf("request received %s, headers: [%s] from [%s]", r.RequestURI, r.Header, r.RemoteAddr)
	uri := r.URL.Path

	if strings.HasPrefix(uri, "/metrics") {
		ph.ServeHTTP(w, r)
		return
	}

	matchedPrefix := ""
	HTTP.rangeRule(func(prefix, dst string) bool {
		if _, ok := matchPrefix(prefix, uri); ok {
			if len(matchedPrefix) < len(prefix) {
				matchedPrefix = prefix
			}
		}
		return true
	})
	if matchedPrefix != "" {
		dst, ok := HTTP.Load(matchedPrefix)
		if !ok {
			log.Printf("error loading prefix: " + matchedPrefix)
			return
		}
		rPath, ok := matchPrefix(matchedPrefix, uri)
		if !ok {
			log.Printf("error matching prefix: " + matchedPrefix)
			return
		}
		u := fmt.Sprintf("%s%s", dst, rPath)
		if r.URL.RawQuery != "" {
			u += fmt.Sprintf("?%s", r.URL.RawQuery)
		}
		// rewrite request
		r.RequestURI = ""
		fullURL, err := url.Parse(u)
		if err != nil {
			log.Printf("error parsing url: " + dst.(string))
			return
		}
		r.URL = fullURL
		// forward
		remote, err := url.Parse(fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host))
		if err != nil {
			log.Printf("error parsing host: " + r.URL.Host)
			return
		}
		r.Header.Set("X-Real-IP", r.RemoteAddr)
		proxy := httputil.NewSingleHostReverseProxy(remote)
		log.Printf("matched route %s, forward to %s", matchedPrefix, dst.(string))
		proxy.ServeHTTP(w, r)
		duration := time.Since(start)
		endpoint := matchedPrefix
		if route := w.Header().Get("__route"); route != "" {
			endpoint = route
		}
		webRequestTotal.With(prometheus.Labels{"method": r.Method, "endpoint": endpoint}).Inc()
		webRequestDuration.With(prometheus.Labels{"method": r.Method, "endpoint": endpoint}).Observe(duration.Seconds())
	} else {
		log.Printf("no router is matched")
		http.NotFound(w, r)
	}
}

func registerHTTPServer(listener cmux.CMux) {
	httpListener := listener.Match(cmux.HTTP1Fast())
	h := &proxyHandler{}
	server := &http.Server{Handler: h}
	http2.ConfigureServer(server, &http2.Server{})
	go func() {
		err := server.Serve(httpListener)
		if err != nil {
			log.Fatalln("server.Serve: ", err)
		}
	}()
}
