package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	// フィールドはデフォルトでは `ID` という名前のフィールドがプライマリフィールドとして使われます。
	// 参照: https://gorm.io/ja_JP/docs/conventions.html
	ID      int    `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Message{})

	e := echo.New()

	// 別ドメインから Fetch API を用いてリクエストを送信可能にするために必要
	// WEB_ORIGIN に設定したドメインからのリクエストのみ許可する
	// 参考: https://developer.mozilla.org/ja/docs/Web/HTTP/CORS
	if web_origin := os.Getenv("WEB_ORIGIN"); web_origin != "" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{web_origin},
			AllowCredentials: true, // これがないと cookie を送信することができない
		}))
	}
	// ミドルウェアを使う。
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/messages", func(c echo.Context) error {
		// get all messages from db
		var messages []Message
		err := db.Find(&messages).Error
		if err != nil {
			c.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, messages)
	})

	e.POST("/send", func(c echo.Context) error {
		// Form 版: https://echo.labstack.com/docs/binding
		var body struct {
			Content string `json:"content"`
		}
		err := json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid JSON")
		}

		if body.Content == "" {
			return c.String(http.StatusBadRequest, "Empty content")
		}

		err = db.Create(&Message{
			Content: body.Content,
		}).Error
		if err != nil {
			c.Error(err)
			return err
		}
		return c.String(http.StatusAccepted, "")
	})

	if err := e.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
