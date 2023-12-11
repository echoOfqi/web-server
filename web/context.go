package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//origin objects
	Writer  http.ResponseWriter
	Request *http.Request

	//request info
	Path   string
	Method string
	Body   map[string]string
	Params map[string]string

	//response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	bodyByte, _ := io.ReadAll(r.Body)
	body := make(map[string]string)
	if len(bodyByte) > 0 {
		if err := json.Unmarshal(bodyByte, &body); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		Body:    body,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) PostBody(key string) string {
	return c.Body[key]
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
