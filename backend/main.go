package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yukyoooo/go_next_ddd/application"
	"github.com/yukyoooo/go_next_ddd/domain/model"
	"github.com/yukyoooo/go_next_ddd/domain/model/employee"
	"github.com/yukyoooo/go_next_ddd/domain/model/milestone"
	"github.com/yukyoooo/go_next_ddd/domain/model/project"
	projectassignment "github.com/yukyoooo/go_next_ddd/domain/model/projectAssignment"
	"golang.org/x/net/websocket"
)

var command = flag.String("usecase", "", "usercase of application")

func main() {
	fmt.Println(model.Db)
	// e := echo.New()            // インスタンスを作成
	// e.Use(middleware.Logger()) // ミドルウェアを設定

	// e.GET("/", func(c echo.Context) error { // ルートを設定
	// 	return c.String(http.StatusOK, "Hello, World!!") // 出力
	// })
	// e.GET("/socket", handleWebSocket)

	// e.Logger.Fatal(e.Start(config.Config.Port)) // サーバーをポート番号で起動

	employeeRepository, err := employee.NewEmployeeRepository(model.Db)
	if err != nil {
		log.Fatal(err)
	}
	employeeService, err := employee.NewEmployeeService(employeeRepository)
	if err != nil {
		log.Fatal(err)
	}
	employeeApplicationService := application.NewEmployeeApplicationService(employeeRepository, *employeeService)

	flag.Parse()
	log.Println(*command) //go run main.go -usecase=register
	switch *command {
	case "register":
		if err := employeeApplicationService.Register("yukyooowoaaa", "yukyoooo", "test32@examplo.com", "password", 1); err != nil {
			log.Println(err)
		}
	default:
		log.Printf("%s is not command. choose in ('register', 'get', 'update', 'delete')", *command)
	}

	projectRepository, err := project.NewProjectRepository(model.Db)
	if err != nil {
		log.Fatal(err)
	}
	projectAssignmentRepository, err := projectassignment.NewProjectAssignmentRepository(model.Db)
	if err != nil {
		log.Fatal(err)
	}
	projectApplicationService := application.NewProjectApplicationService(projectRepository, projectAssignmentRepository)
	if err := projectApplicationService.Create(1, "project", time.Now(), time.Now()); err != nil {
		log.Println(err)
	}

	milestoneRepository, err := milestone.NewMilestoneRepository(model.Db)
	if err != nil {
		log.Fatal(err)
	}
	milestoneApplicationService := application.NewMilestoneApplicationService(milestoneRepository)
	if err := milestoneApplicationService.Create(1, "milestone", time.Now(), time.Now()); err != nil {
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
