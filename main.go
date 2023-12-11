package main

import (
	"net/http"
	"web"
)

func main() {
	r := web.NewEngine()
	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.POST("/hello", func(c *web.Context) {
		// Parse query parameter
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.POST("/loginByBody", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{
			"username": c.PostBody("username"),
			"password": c.PostBody("password"),
		})
	})

	r.GET("/hello/:name", func(c *web.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
