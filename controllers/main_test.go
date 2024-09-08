package controllers_test

import (
	"testing"

	"github.com/YamaguchiKoki/go_prc/controllers"
	"github.com/YamaguchiKoki/go_prc/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

/*
Unit Testの流れ
1: 対象リソース作成
2: テスト対象の関数に入れるinputを定義
3: テスト対象の関数を実行してoutputを得る
4: outputが期待通りかチェック
*/

// 1
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}