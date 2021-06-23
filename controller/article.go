package controller

import (
	"net/http"
	"strconv"
    "encoding/json"
	"example.com/hello/service"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

func GetAllArticles(c echo.Context) error {
	service.GetAllArticlesService()
	return return c.JSON(http.StatusOK, &article)
}

/* func CreateArticle(c echo.Context) error {
	return service.CreateArticleService(c echo.Context)
}

func ShowArticleService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return service.ShowArticleService(id)

func UpdateArticleService(c echo.Context) error {
	return service.UpdateArticleService(c echo.Context)
}

func DeleteArticleService(c echo.Context) error {
	return service.DeleteArticleService(c echo.Context)
} */
