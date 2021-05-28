package api

import (
	"net/http"
	"strconv"

	"example.com/user/hello/database"
	"example.com/user/hello/model"
	"github.com/labstack/echo/v4"
)

func GetAllArticles(c echo.Context) error {
	var articles []model.Article
	database.DB.Find(&articles)
	return c.JSON(http.StatusOK, &articles)
}

func CreateArticle(c echo.Context) error {
	article := model.Article{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}
	database.DB.Create(&article)
	return c.JSON(http.StatusOK, &article)
}

func ShowArticle(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.First(&article, id)
	database.DB.Model(&article).Update("read", true)
	return c.JSON(http.StatusOK, &article)
}

func UpdateArticle(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.First(&article, id)
	database.DB.Model(&article).Update("title", "new200")

	return c.JSON(http.StatusOK, &article)
}

func DeleteArticle(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&article, id)
	return c.JSON(http.StatusOK, &article)
}
