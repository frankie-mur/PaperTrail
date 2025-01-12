package db

import (
	"context"
	"time"
)

// Article represents a saved article.
type Article struct {
	ID        string            `json:"id"`
	UserID    string            `json:"user_id"`
	URL       string            `json:"url"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Metadata  map[string]string `json:"metadata"`
	Tags      []string          `json:"tags"`
	CreatedAt time.Time         `json:"created_at"`
}

// DB defines the interface for database operations.
type DB interface {
	SaveArticle(ctx context.Context, article Article) error
	GetArticlesByUser(ctx context.Context, userID string) ([]Article, error)
	GetArticle(ctx context.Context, articleID string) (*Article, error)
	UpdateArticle(ctx context.Context, articleID string, updates map[string]interface{}) error
	DeleteArticle(ctx context.Context, articleID string) error
}
