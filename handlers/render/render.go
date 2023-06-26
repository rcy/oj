package render

import (
	"bytes"
	"html/template"
	"net/http"
)

func Execute(w http.ResponseWriter, t *template.Template, data any) {
	bytes, err := ExecuteToBytes(t, data)
	if err != nil {
		Error(w, err.Error(), 500)
	} else {
		w.Write(bytes)
	}
}

func ExecuteNamed(w http.ResponseWriter, t *template.Template, name string, data any) {
	bytes, err := ExecuteNamedToBytes(t, name, data)
	if err != nil {
		Error(w, err.Error(), 500)
	} else {
		w.Write(bytes)
	}
}

func ExecuteToBytes(t *template.Template, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := t.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func ExecuteNamedToBytes(t *template.Template, name string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := t.ExecuteTemplate(buf, name, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
