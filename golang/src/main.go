package main

import (
  "fmt"
  "log"
  "os"
	"net/http"
	"time"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gin-gonic/gin"
  // "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primary_key"`
	Task        string         `gorm:"size:255"`
	IsCompleted bool           `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func dns() string {
  mysqlUser := os.Getenv("MYSQL_USER")
  mysqlPassword := os.Getenv("MYSQL_PASSWORD")
  mysqlDatabase := os.Getenv("MYSQL_DATABASE")
  return fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlDatabase)
}

func connectDB(dns string) (db *gorm.DB) {
  db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
  if err != nil {
    log.Fatal("Error opening database:", err)
  }
  return
}

func main() {
  dns := dns()
  db := connectDB(dns)

  // データベースにテーブルを作成
  db.AutoMigrate(&Task{})



  // Ginエンジンのインスタンスを作成
  r := gin.Default()

  // タスクを取得するエンドポイント
	r.GET("/tasks", func(c *gin.Context) {
		var tasks []Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})

	// 新しいタスクを作成するエンドポイント
	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}
		db.Create(&task)
		c.JSON(http.StatusOK, task)
	})

	// タスクを更新するエンドポイント
	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")

		if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
		}

		if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}

		db.Save(&task)
		c.JSON(http.StatusOK, task)
	})

	// タスクを削除するエンドポイント
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")

		if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
		}

		db.Delete(&task)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

  // 8080ポートでサーバーを起動
  r.Run(":8080")
}