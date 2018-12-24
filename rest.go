package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// JSON is a map alias, just for convenience
type JSON map[string]interface{}

// RenderJSON sends data as json
func RenderJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	b, err := json.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal data")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(b); err != nil {
		return errors.Wrapf(err, "failed to send response to %s", r.RemoteAddr)
	}
	return nil
}

// RenderJSONFromBytes sends binary data as json
func RenderJSONFromBytes(w http.ResponseWriter, r *http.Request, data []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(data); err != nil {
		return errors.Wrapf(err, "failed to send response to %s", r.RemoteAddr)
	}
	return nil
}

// RenderJSONWithHTML allows html tags and forces charset=utf-8
func RenderJSONWithHTML(w http.ResponseWriter, r *http.Request, v interface{}) error {

	encodeJSONWithHTML := func(v interface{}) ([]byte, error) {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(v); err != nil {
			return nil, errors.Wrap(err, "json encoding failed")
		}
		return buf.Bytes(), nil
	}

	data, err := encodeJSONWithHTML(v)
	if err != nil {
		return errors.Wrap(err, "json encoding failed")
	}
	return RenderJSONFromBytes(w, r, data)
}
