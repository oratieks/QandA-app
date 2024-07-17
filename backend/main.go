package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}

	// データベース接続を初期化
	initDB()

	// Gin ルーターを設定
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// CORSミドルウェアを追加
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Reactアプリのアドレス
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// ルートを設定
	setupRoutes(r)

	// サーバーを起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func setupRoutes(r *gin.Engine) {
	r.GET("/api/subjects", getSubjects)
	r.GET("/api/subjects/:id/questions", getQuestionsBySubject)
	r.POST("/api/questions", createQuestion)
	r.GET("/api/questions/:id", getQuestion)
	r.POST("/api/questions/:id/answers", createAnswer)
	r.DELETE("/api/questions/:id", deleteQuestion)
}
