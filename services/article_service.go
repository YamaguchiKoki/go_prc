package services

import (
	"database/sql"
	"errors"

	"github.com/YamaguchiKoki/go_prc/apperrors"
	"github.com/YamaguchiKoki/go_prc/models"
	"github.com/YamaguchiKoki/go_prc/repositories"
)

//記事投稿
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

// 指定ページ分だけ記事取得
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err = apperrors.GetDataFailed.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

//指定IDの記事一覧を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

//いいね＋１
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist tarfet article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return models.Article{
		ID: article.ID,
		Title: article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
		NiceNum: article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}