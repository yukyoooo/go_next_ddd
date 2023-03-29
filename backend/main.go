package main

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	application "github.com/yukyoooo/go_next_ddd/application"
	"github.com/yukyoooo/go_next_ddd/domain/model"
	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println(model.Db)
	// e := echo.New()            // インスタンスを作成
	// e.Use(middleware.Logger()) // ミドルウェアを設定

	// e.GET("/", func(c echo.Context) error { // ルートを設定
	// 	return c.String(http.StatusOK, "Hello, World!!") // 出力
	// })
	// e.GET("/socket", handleWebSocket)

	// e.Logger.Fatal(e.Start(config.Config.Port)) // サーバーをポート番号で起動

	err := application.RegisterEmployeeService("taro", "yamada", "test@example.com", "MyP@ssw0rd", 1);
	if err != nil {
		log.Println(err)
	}

	err = application.RegisterProjectService(1, "test project", 1, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}
}



func handleWebSocket(c echo.Context) error {
	log.Println("Serving at localhost:8000...")
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// 初回のメッセージを送信
		err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			// Client からのメッセージを読み込む
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				if err.Error() == "EOF" {
					log.Println(fmt.Errorf("read %s", err))
					break
				}
				log.Println(fmt.Errorf("read %s", err))
				c.Logger().Error(err)
			}

			// Client からのメッセージを元に返すメッセージを作成し送信する
			err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}