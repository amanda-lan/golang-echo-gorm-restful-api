package service

import (
	"net/http"
/* 	"strconv"
	"encoding/json"
	"github.com/labstack/echo/v4" */
	"example.com/hello/dao"
	"gorm.io/datatypes"
)

func GetAllArticlesService() error {
	return dao.CreateArticleDao()
}

/* func CreateArticleService(c echo.Context) error {
	return dao.CreateArticleDao(c echo.Context)
}

func ShowArticleService(c echo.Context) error {

	return dao.ShowArticleDao(c echo.Context)

func UpdateArticleService(c echo.Context) error {
	return dao.UpdateArticleDao(c echo.Context)
}

func DeleteArticleService(c echo.Context) error {
	return dao.DeleteArticleDao(c echo.Context)
} */
