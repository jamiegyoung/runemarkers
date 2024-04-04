package devserver

import (
	"github.com/gin-gonic/gin"
	"github.com/jamiegyoung/runemarkers-go/builder"
)

func Start() {
	r := gin.Default()

	r.Static("/api", "./public/api")
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")
	r.Static("/thumbnails", "./public/thumbnails")

  r.Use(func(c *gin.Context) {
    // build again before each path, this is why it's a dev server ;)
    builder.Build(true)
    c.Next()
  })

	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/:page", func(c *gin.Context) {
		c.File("./public/" + c.Param("page") + ".html")
	})

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
