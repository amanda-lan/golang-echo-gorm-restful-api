package api

import (
	"net/http"

	"example.com/user/hello/database"
	"example.com/user/hello/model"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello World!")

}

func GetAllArticles(c echo.Context) error {

	articles := []model.Article{}
	results := database.DB.Find(&articles)
	// spew.Dump(json.Marshal(users))
	return c.JSON(http.StatusOK, results.RowsAffected)
	//	return c.String(http.StatusOK, result.RowsAffected)
}

func SaveArticle(c echo.Context) error {
	article := model.Article{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}
	result := database.DB.Select("title", "body").Create(&article)
	return c.String(http.StatusOK, result.Name())
}

func ShowArticle(c echo.Context) error {

	articles := []model.Article{}
	results := database.DB.Find(&articles)
	// spew.Dump(json.Marshal(users))
	return c.JSON(http.StatusOK, results.RowsAffected)
	//	return c.String(http.StatusOK, result.RowsAffected)
}

func UpdateArticle(c echo.Context) error {

	articles := []model.Article{}
	results := database.DB.Find(&articles)
	// spew.Dump(json.Marshal(users))
	return c.JSON(http.StatusOK, results.RowsAffected)
	//	return c.String(http.StatusOK, result.RowsAffected)
}

func DeleteArticle(c echo.Context) error {
	article := model.Article{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}
	result := database.DB.Select("title", "body").Create(&article)
	return c.String(http.StatusOK, result.Name())
}
