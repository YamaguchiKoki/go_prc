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
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)

	//go文は戻り値のない関数にしか使えないため、戻り値のない無名関数でラップし、即時実行
	//メインゴルーチンの変数を直接参照しないように引数を利用
	//チャネルを通じて別のゴルーチンの変数に書き込み
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList *[]models.Comment
		err error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
		
	}(commentChan, s.db, articleID)

	//複数後ルーチンからの結果の受け取りイディオム
	for i:= 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
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