package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/revett/projects/internal/uci-engine-wrapper/handlers"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/calculate", handlers.Calculate)

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler is required by Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
			log.Println(err)
	}
	fmt.Println(path)

	var files []string

	root := "/var/task/handler"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
	})
	if err != nil {
			log.Println(err)
	}
	for _, file := range files {
			fmt.Println(file)
	}
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
