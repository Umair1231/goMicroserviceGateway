package proxy

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

var serviceMap = map[string]string{
	"/users": "http://localhost:8000",
}

func GetTargetURL(apiPath string, query string) (string, error) {
	apiPath = strings.TrimPrefix(apiPath, "/api")

	for prefix, baseURL := range serviceMap {
		if strings.HasPrefix(apiPath, prefix) {
			base, err := url.Parse(baseURL)
			if err != nil {
				return "", fmt.Errorf("invalid base URL for service %s", prefix)
			}

			safePath := path.Clean(apiPath)
			base.Path = path.Join(base.Path, safePath)
			if query != "" {
				base.RawQuery = query
			}
			return base.String(), nil
		}
	}

	return "", fmt.Errorf("no matching service for path")
}
