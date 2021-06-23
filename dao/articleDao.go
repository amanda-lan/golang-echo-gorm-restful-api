package dao

import (
/* 	"net/http"
	"strconv"
	"encoding/json"
	"github.com/labstack/echo/v4" */
	"example.com/hello/database"
	"example.com/hello/model"
)

func (articles *[]model.Article) GetAllArticlesDao() error {
	database.DB.Find(&articles)
	return &articles
}

/* func CreateArticleDao(c echo.Context) error {
    body, _ := json.Marshal(c.FormValue("body"))
	article := model.Article{
		Title: c.FormValue("title"),
		Body:  body,
	}
	database.DB.Create(&article)
	return c.JSON(http.StatusOK, &article)
}

func ShowArticleDao(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.First(&article, id)
	database.DB.Model(&article).Update("read", true)
	return c.JSON(http.StatusOK, &article)
}

func UpdateArticleDao(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.First(&article, id)
	database.DB.Model(&article).Update("title", "new200")

	return c.JSON(http.StatusOK, &article)
}

func DeleteArticleDao(c echo.Context) error {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&article, id)
	return c.JSON(http.StatusOK, &article)
} */
