package main

import (
	"github.com/labstack/echo/v4"
	"github.com/murakami10/myapi/handlers"
	"log"
)

func main(){
	e := echo.New()

	e.GET("/hello", handlers.HelloHandler)
	e.GET("/article/list", handlers.ArticleListHandler)
	e.POST("/article", handlers.PostArticleHandler)
	e.GET("/article/:id", handlers.ArticleDetailHandler)
	e.POST("/article/nice", handlers.PostNiceHandler)
	e.POST("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	e.Logger.Fatal((e.Start(":8080")))
}


