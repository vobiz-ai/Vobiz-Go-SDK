package compliance

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func newFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		fi, err := file.Stat()
		if err != nil {
			return nil, err
		}
		file.Close()
		part, err := writer.CreateFormFile(paramName, fi.Name())
		if err != nil {
			return nil, err
		}
		if _, err := part.Write(contents); err != nil {
			return nil, err
		}
	}
	for k, v := range params {
		_ = writer.WriteField(k, v)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}

func structToStringMap(v interface{}, skipField string) map[string]string {
	result := make(map[string]string)
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	for i := 0; i < t.NumField(); i++ {
		tag := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if tag == skipField || tag == "" || tag == "-" {
			continue
		}
		result[tag] = val.Field(i).String()
	}
	return result
}
