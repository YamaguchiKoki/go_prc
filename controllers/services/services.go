package services

import "github.com/YamaguchiKoki/go_prc/models"

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}