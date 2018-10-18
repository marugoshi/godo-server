package httputils

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Api struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (a *Api) Html(code int, name string, path string, data interface{}) error {
	body, err := a.parse(name, path, data)
	if err != nil {
		log.Printf("1: %s", err)
		return err
	}
	return a.show(code, body, ContentTypeTextHtml)
}

func (a *Api) RawText(code int, body string) error {
	return a.show(code, body, ContentTypeTextPlain)
}

func (a *Api) parse(name string, path string, data interface{}) (string, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func (a *Api) show(code int, body string, contentType string) error {
	a.ResponseWriter.Header().Set("Content-Type", contentType)
	a.WriteHeader(code)
	_, err := io.WriteString(a.ResponseWriter, fmt.Sprintf("%s\n", body))
	return err
}

type Func func(ctx context.Context, api Api) error
