package controller

import (
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/articles", GetAllArticles)
/* 	e.POST("/articles", CreateArticle)
	e.GET("/articles/:id", ShowArticle)
	e.PUT("/articles/:id", UpdateArticle)
	e.DELETE("/articles/:id", DeleteArticle) */
	return e
}
