package repositories

import (
	"github.com/dikaizm/govision_backend/pkg/domain"
	repo_intf "github.com/dikaizm/govision_backend/pkg/repositories/interfaces"
	"gorm.io/gorm"
)

type DbArticleRepository struct {
	DB *gorm.DB
}

func NewDbArticleRepository(db *gorm.DB) repo_intf.ArticleRepository {
	return &DbArticleRepository{DB: db}
}

func (r *DbArticleRepository) Create(article *domain.Article) error {
	if err := r.DB.Create(article).Error; err != nil {
		return err
	}

	return nil
}

func (r *DbArticleRepository) FindAll() ([]*domain.Article, error) {
	var articles []*domain.Article

	if err := r.DB.Order("created_at desc").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *DbArticleRepository) FindByID(id string) (*domain.Article, error) {
	var article domain.Article

	if err := r.DB.Where("id = ?", id).Preload("Author").First(&article).Error; err != nil {
		return nil, err
	}

	// Increment the read count
	article.ReadCount++

	if err := r.Update(&article); err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *DbArticleRepository) CreateBulk(articles []*domain.Article) error {
	if err := r.DB.Create(articles).Error; err != nil {
		return err
	}

	return nil
}

func (r *DbArticleRepository) Update(article *domain.Article) error {
	if err := r.DB.Save(article).Error; err != nil {
		return err
	}

	return nil
}
