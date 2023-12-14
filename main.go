package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"web"
)

type student struct {
	Name string
	Age  int8
}

func onlyForV2() web.HandlerFunc {
	return func(c *web.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := web.NewEngine()
	r.Use(web.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	v1 := r.Group("/v1")
	{

		v1.POST("/hello", func(c *web.Context) {
			// Parse query parameter
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.POST("/login", func(c *web.Context) {
			c.JSON(http.StatusOK, web.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

		v2.POST("/loginByBody", func(c *web.Context) {
			c.JSON(http.StatusOK, web.H{
				"username": c.PostBody("username"),
				"password": c.PostBody("password"),
			})
		})

		v2.GET("/hello/:name", func(c *web.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{"filepath": c.Param("filepath")})
	})

	r.GET("/students", func(c *web.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", web.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *web.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", web.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
