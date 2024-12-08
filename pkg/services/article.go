package services

import (
	"github.com/dikaizm/govision_backend/internal/dto/request"
	"github.com/dikaizm/govision_backend/pkg/domain"
	"github.com/dikaizm/govision_backend/pkg/helpers"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	service_intf "github.com/dikaizm/govision_backend/pkg/services/interfaces"
)

type ArticleService struct {
	articleRepo repo_intf.ArticleRepository
}

func NewArticleService(articleRepo repo_intf.ArticleRepository) service_intf.ArticleService {
	return &ArticleService{articleRepo: articleRepo}
}

func (s *ArticleService) Create(p *request.CreateArticle) error {
	article := &domain.Article{
		ID:        helpers.GenerateUserID(),
		Title:     p.Title,
		Body:      p.Body,
		Image:     p.Image,
		ReadCount: 0,
		AuthorID:  p.AuthorID,
	}

	if err := s.articleRepo.Create(article); err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) FindAll(filter *request.FilterGetArticle) ([]*domain.Article, error) {
	articles, err := s.articleRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *ArticleService) FindByID(id string) (*domain.Article, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) CreateBulk(p []*request.CreateArticle) error {
	articles := make([]*domain.Article, 0, len(p))

	for _, article := range p {
		articles = append(articles, &domain.Article{
			ID:        helpers.GenerateUserID(),
			Title:     article.Title,
			Body:      article.Body,
			Image:     article.Image,
			ReadCount: 0,
			AuthorID:  article.AuthorID,
		})
	}

	if err := s.articleRepo.CreateBulk(articles); err != nil {
		return err
	}

	return nil
}
