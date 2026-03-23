package media

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { return quoteEscaper.Replace(s) }

// Upload uploads one or more media files.
func (s *Service) Upload(params UploadParams) (*UploadResponseBody, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, f := range params.Files {
		file, err := os.Open(f.FilePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		filename := filepath.Base(f.FilePath)
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes("file"), escapeQuotes(filename)))
		h.Set("Content-Type", f.ContentType)
		part, err := writer.CreatePart(h)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, file); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	uploadURL := fmt.Sprintf("%s/api/v1/Account/%s/Media", s.baseURL, s.authID)
	req, err := http.NewRequest("POST", uploadURL, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(s.authID, s.authToken)

	resp := &UploadResponseBody{}
	return resp, s.r.ExecuteRequest(req, resp)
}
