package numbers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func setModernPath(req *http.Request, authID, suffix string) {
	s := strings.TrimPrefix(suffix, "/")
	req.URL.Path = fmt.Sprintf("/api/v1/account/%s/%s", authID, s)
}

func escNumber(number string) string { return url.PathEscape(number) }
