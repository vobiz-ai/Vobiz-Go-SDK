package trunks

import (
	"fmt"
	"net/http"
	"strings"
)

func setModernPath(req *http.Request, authID, suffix string) {
	s := strings.TrimPrefix(suffix, "/")
	req.URL.Path = fmt.Sprintf("/api/v1/account/%s/%s", authID, s)
}
