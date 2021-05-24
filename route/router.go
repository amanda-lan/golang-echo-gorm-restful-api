package route

import (
	"example.com/user/hello/api"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/articles", api.GetAllArticles)
	e.POST("/articles", api.SaveArticle)
	e.GET("/articles/:id", api.ShowArticle)
	e.PUT("/articles/:id", api.UpdateArticle)
	e.DELETE("/articles/:id", api.DeleteArticle)
	return e
}
