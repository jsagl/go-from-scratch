package storage

import (
	"database/sql"
	"github.com/jsagl/go-from-scratch/models"
)

type PostgresArticleStore struct {
	Connection *sql.DB
}

func NewPostgresArticleStore(connection *sql.DB) ArticleStore {
	return &PostgresArticleStore{Connection: connection}
}

func (store *PostgresArticleStore) GetArticle() (*models.Article, error) {
	var (
		id int64
		title string
		content string
	)

	query := `SELECT id, title, content FROM articles WHERE ID = 1`
	if err := store.Connection.QueryRow(query).Scan(&id, &title, &content); err != nil {
		return nil, err
	}

	return &models.Article{ID: id, Title: title, Content: content}, nil
}
