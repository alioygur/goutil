package goutil

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// DecodeReq decodes request's body to given interface
func DecodeReq(r *http.Request, to interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(to); err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}

// QueryValue returns url's query value, if it not exists then return ""
func QueryValue(k string, r *http.Request) string {
	values := r.URL.Query()[k]

	if len(values) != 0 {
		return values[0]
	}

	return ""
}

// QueryValueInt gets url query value and tries convert it to int.
func QueryValueInt(k string, r *http.Request) (int, error) {
	qv := QueryValue(k, r)
	if qv == "" {
		return 0, nil
	}
	return strconv.Atoi(qv)
}
