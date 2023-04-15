package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// /hello のハンドラ
func HelloHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, world")
}

// /article のハンドラ
func PostArticleHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Article...\n")
}

// /article/list のハンドラ
func ArticleListHandler(c echo.Context) error{

	var page int

	pageParam := c.QueryParam("page")
	if pageParam == ""{
		page = 1
	}else{
		var err error
		page, err = strconv.Atoi(c.QueryParam("page"))

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid query parameter")
		}
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)

	return c.String(http.StatusOK,resString)
}

// /article/1 のハンドラ
func ArticleDetailHandler(c echo.Context) error {
	articleID, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid query parameter")
	}
	resString := fmt.Sprintf("Article No. %d \n", articleID)
	return c.String(http.StatusOK, resString)
}

// /article/nice のハンドラ
func PostNiceHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Nice...\n")
}

// /comment のハンドラ
func PostCommentHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Comment...\n")
}