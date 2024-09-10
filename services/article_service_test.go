package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/YamaguchiKoki/go_prc/services"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbHost := "db"

	// DB接続情報のフォーマット
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)

	m.Run()
}

//ベンチマークテスト
func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}