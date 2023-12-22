package routes

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"service-fleetime/cmd/controller"
	"service-fleetime/cmd/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

var NewLogger logger.Interface

func ClearDirectory(dir string) error {
	// Open the directory.
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	// Read all file names from the directory.
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Loop over the file names and delete each one.
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}

func FetcherRoutes() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("../static/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.StaticFS("/static", http.Dir("../static"))
	r.GET("/clear-static", func(c *gin.Context) {
		// Open the directory.
		d, err := os.Open("../static")
		if err != nil {
			helpers.ErrorResponse(c, err.Error())
			return
		}
		defer d.Close()

		// Read all file names from the directory.
		names, err := d.Readdirnames(-1)
		if err != nil {
			helpers.ErrorResponse(c, err.Error())
			return
		}

		// Loop over the file names and clear each one.
		for _, name := range names {
			f, err := os.OpenFile(filepath.Join("../static", name), os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				helpers.ErrorResponse(c, err.Error())
				return
			}
			defer f.Close()

			err = f.Truncate(0)
			if err != nil {
				helpers.ErrorResponse(c, err.Error())
				return
			}
		}

		helpers.APIResponse(c, "Logs Cleared", http.StatusOK, nil)
	})

	r.GET("/fetcher", controller.GetData)
	r.GET("/fetcher/all", controller.FetchAllEmployee)
	r.GET("/fetcher/all/send", controller.FetchAllEmployeeAndSendToPostgres)

	r.Run(":8080")
}
