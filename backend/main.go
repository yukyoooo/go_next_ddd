package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yukyoooo/go_next_ddd/config"
	"github.com/yukyoooo/go_next_ddd/domain/models"
	employee "github.com/yukyoooo/go_next_ddd/domain/models/employees"
	"github.com/yukyoooo/go_next_ddd/enum"
	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println(models.Db)
	e := echo.New()            // インスタンスを作成
	e.Use(middleware.Logger()) // ミドルウェアを設定

	e.GET("/", func(c echo.Context) error { // ルートを設定
		return c.String(http.StatusOK, "Hello, World!!") // 出力
	})
	e.GET("/socket", handleWebSocket)

	e.Logger.Fatal(e.Start(config.Config.Port)) // サーバーをポート番号で起動


	// err := CreateEmployee(models.Db, "taro", "yamadaa", "testtes2t@test.com", "MyP@ssw0rd", enum.Waiting)
	// if err != nil {
	// 	log.Println(err)
	// }
}

func CreateEmployee(Db *sql.DB, firstName string, lastName string, email string, password string, role enum.Role) (err error) {
	newEmployeeName, err := employee.NewFullName(firstName, lastName)
	if err != nil {
		return err
	}

	newEmail, err := employee.NewEmail(email)
	if err != nil {
		return err
	}

	newPassword, err := employee.NewPassword(password)
	if err != nil {
		return err
	}

	newEmployee, err := employee.NewEmployee(*newEmployeeName, *newEmail, *newPassword, role)
	if err != nil {
		return err
	}
	fmt.Println(newEmployee)

	userService, err := employee.NewEmployeeService(models.Db)
	if err != nil {
		return err
	}
	isExists, err := userService.Exists(newEmployee)
	if err != nil {
		return err
	}

	if isExists {
		log.Println("userservice.Exists() :既に存在する名前、もしくはメールアドレスです")
	} else {
		newEmployee.CreateEmployee()
	}
	return nil
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