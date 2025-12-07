package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ForwardRequest(targetURL string) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, fmt.Errorf("invalid target URL: %w", err)
	}

	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path
			req.URL.RawQuery = target.RawQuery

			req.Host = target.Host

		},
		ModifyResponse: func(resp *http.Response) error {
			return nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, "Upstream service error: "+err.Error(), http.StatusBadGateway)
		},
	}, nil
}
