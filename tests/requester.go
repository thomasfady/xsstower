package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type Requester struct {
	Engine  *gin.Engine
	Cookies map[string]string
}

type JSONRequest struct {
	Method string
	Url    string
	Body   interface{}
}

type JSONResonse struct {
	Code          int
	ResponseBytes []byte
}

func NewRequester(engine *gin.Engine) (requester *Requester) {
	requester = &Requester{
		Engine: engine,
	}
	requester.Cookies = make(map[string]string)
	return
}

func (requester *Requester) Get(url string) JSONResonse {
	return requester.DoRequest("GET", url, nil)
}
func (requester *Requester) Post(url string, body interface{}) JSONResonse {
	return requester.DoRequest("POST", url, body)
}

func (requester *Requester) addCookieHeader(r *http.Request) {
	var cookies []string
	for c, v := range requester.Cookies {
		cookies = append(cookies, c+"="+v)
	}
	r.Header.Add("Cookie", strings.Join(cookies, "; "))
}

func (requester *Requester) DoRequest(method string, url string, body interface{}) JSONResonse {
	w := httptest.NewRecorder()
	data, _ := json.Marshal(body)
	r, _ := http.NewRequest(method, url, bytes.NewReader(data))
	r.Header.Add("Content-Type", "application/json")
	requester.addCookieHeader(r)
	requester.Engine.ServeHTTP(w, r)
	h := w.Header().Get("Set-Cookie")
	if h != "" {
		parts := strings.Split(strings.Split(h, ";")[0], "=")
		requester.Cookies[parts[0]] = parts[1]
	}
	return JSONResonse{
		Code:          w.Code,
		ResponseBytes: w.Body.Bytes(),
	}
}

func (res *JSONResonse) ResponseString() string {
	var error_msg string
	if err := json.Unmarshal(res.ResponseBytes, &error_msg); err != nil {
		return string(res.ResponseBytes)
	}
	return error_msg
}
