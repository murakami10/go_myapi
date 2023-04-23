package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/murakami10/myapi/models"
	"net/http"
	"strconv"
)

// /hello のハンドラ
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}

// /article のハンドラ
func PostArticleHandler(c echo.Context) error {

	var article models.Article

	if err := c.Bind(article); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, article)
}

// /article/list のハンドラ
func ArticleListHandler(c echo.Context) error {

	var page int

	pageParam := c.QueryParam("page")
	if pageParam == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(c.QueryParam("page"))

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid query parameter")
		}
	}

	articleList := []models.Article{models.Article1, models.Article2}
	if len(articleList) == 0 {
		errMsg := fmt.Sprintf("fail to fetch articles in page: %d", page)
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, articleList)
}

// /article/1 のハンドラ
func ArticleDetailHandler(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid query parameter")
	}

	article := &models.Article1
	if article == nil {
		errMsg := fmt.Sprintf("fail to fetch article in article ID: %d", articleID)
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, *article)
}

// /article/nice のハンドラ
func PostNiceHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Nice...\n")
}

// /comment のハンドラ
func PostCommentHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Posting Comment...\n")
}
