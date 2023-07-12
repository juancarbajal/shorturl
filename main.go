package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juancarbajal/shorturl/models"
	"github.com/juancarbajal/shorturl/pkg"
)

type ShortUrlBind struct {
	ID string `uri:"id" binding"required,uuid"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	route := gin.Default()
	sm := models.ShortenModel{}

	//route url shorten
	route.GET("/:id", func(c *gin.Context) {
		var sub ShortUrlBind
		if err := c.ShouldBindUri(&sub); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		if url := sm.GetUrl(sub.ID); url != "" {
			c.Redirect(http.StatusMovedPermanently, url)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	//route form
	route.LoadHTMLGlob("templates/*")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Shorten URL ",
		})
	})
	//route save form
	route.POST("/", func(c *gin.Context) {
		url := c.PostForm("url")
		shortURL := pkg.ShortenURL(url)
		sm.SaveUrl(url, shortURL)
		protocol := "http://"
		if c.Request.TLS != nil {
			protocol = "https://"
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"shortURL": protocol + c.Request.Host + "/" + shortURL,
			"url":      url,
		})
	})

	//route ping
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "ok"})
	})
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	route.Run(":" + PORT)
}
