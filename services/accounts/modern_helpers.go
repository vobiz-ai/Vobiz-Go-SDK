package accounts

import (
	"fmt"
	"net/http"
	"strings"
)

func setModernAccountPath(req *http.Request, authID, suffix string) {
	s := strings.TrimPrefix(suffix, "/")
	req.URL.Path = fmt.Sprintf("/api/v1/account/%s/%s", authID, s)
}
