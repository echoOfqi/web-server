package main

import (
	"net/http"
	"web"
)

func main() {
	r := web.NewEngine()
	r.GET("/index", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *web.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gee</h1>")
		})

		v1.POST("/hello", func(c *web.Context) {
			// Parse query parameter
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
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

	r.Run(":9999")
}
