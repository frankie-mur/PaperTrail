package article

import (
	"context"

	"github.com/frankie-mur/PaperTrail/db"
)

type ArticleService struct {
	db db.DB
}

// NewArticleService initializes the service with a database implementation.
func NewArticleService(database db.DB) *ArticleService {
	return &ArticleService{db: database}
}

// SaveArticle saves an article through the database interface.
func (s *ArticleService) SaveArticle(ctx context.Context, article db.Article) error {
	return s.db.SaveArticle(ctx, article)
}

// GetUserArticles retrieves all articles for a user.
func (s *ArticleService) GetUserArticles(ctx context.Context, userID string) ([]db.Article, error) {
	return s.db.GetArticlesByUser(ctx, userID)
}
