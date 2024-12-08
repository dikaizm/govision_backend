package controllers

import (
	"net/http"

	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/internal/dto/response"
	controller_intf "github.com/dikaizm/govision_backend/internal/http/controllers/interfaces"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type ArticleController struct {
	articleService service_intf.ArticleService
}

func NewArticleController(articleService service_intf.ArticleService) controller_intf.ArticleController {
	return &ArticleController{articleService: articleService}
}

func (c *ArticleController) Create(w http.ResponseWriter, r *http.Request) {
	req := request.CreateArticle{}

	// Decode the request body
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Validate the struct
	err := validate.Struct(req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Get current user
	author, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	req.AuthorID = author.ID

	// Create the article
	if err := c.articleService.Create(&req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Create article success",
	}, http.StatusCreated)

}

func (c *ArticleController) ViewAll(w http.ResponseWriter, r *http.Request) {
	var articleResponse []*response.GetArticle
	var err error

	// Query params
	var filter request.FilterGetArticle
	filter.SizeParam = r.URL.Query().Get("size")
	if filter.SizeParam != "" {
		// Must be a number
		if err := validate.Var(filter.SizeParam, "numeric"); err != nil {
			helpers.SendResponse(w, response.Response{
				Status: "error",
				Error:  err.Error(),
			}, http.StatusBadRequest)
			return
		}

		filter.Size, err = helpers.StringToInt(filter.SizeParam)
		if err != nil {
			helpers.SendResponse(w, response.Response{
				Status: "error",
				Error:  err.Error(),
			}, http.StatusBadRequest)
			return
		}
	}

	articles, err := c.articleService.FindAll(&filter)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	articleResponse = []*response.GetArticle{}

	for _, article := range articles {
		articleResponse = append(articleResponse, &response.GetArticle{
			ID:        article.ID,
			Title:     article.Title,
			Image:     article.Image,
			ReadCount: article.ReadCount,
			CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	helpers.SendResponse(w, response.Response{
		Status: "success",
		Data:   articleResponse,
	}, http.StatusOK)
}

func (c *ArticleController) View(w http.ResponseWriter, r *http.Request) {
	var articleResponse response.GetArticle

	id := helpers.UrlVars(r, "id")

	article, err := c.articleService.FindByID(id)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	articleResponse = response.GetArticle{
		ID:        article.ID,
		Title:     article.Title,
		Body:      article.Body,
		Image:     article.Image,
		ReadCount: article.ReadCount,
		Author:    article.Author.Name,
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Get article success",
		Data:    articleResponse,
	}, http.StatusOK)
}

func (c *ArticleController) CreateBulk(w http.ResponseWriter, r *http.Request) {
	req := []*request.CreateArticle{}

	// Decode the request body
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Validate the struct
	for _, r := range req {
		if err := validate.Struct(r); err != nil {
			helpers.SendResponse(w, response.Response{
				Status: "error",
				Error:  err.Error(),
			}, http.StatusBadRequest)
			return
		}
	}

	// Get current user
	author, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	for i := range req {
		req[i].AuthorID = author.ID
	}

	// Create the article
	if err := c.articleService.CreateBulk(req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Create article success",
	}, http.StatusCreated)
}
