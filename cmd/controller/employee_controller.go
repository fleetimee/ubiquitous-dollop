package controller

import (
	"net/http"
	"service-fleetime/cmd/helpers"
	"service-fleetime/cmd/models"
	"service-fleetime/cmd/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func FetchAllEmployee(ctx *gin.Context) {

	employee, err := repository.FetchAllEmployee()

	if employee == nil {
		// Create empty employee slice
		employee = []models.Employee{}

		helpers.APIResponse(ctx, "Data pegawai tidak ditemukan", http.StatusNotFound, employee)

		return

	}

	if err != nil {
		helpers.ErrorResponse(ctx, err.Error())

		return
	}

	helpers.APIResponse(ctx, "Success", http.StatusOK, employee)
}

func FetchAllEmployeeAndSendToPostgres(ctx *gin.Context) {
	startTime := time.Now()

	insertedRows, deletedRows, err := repository.FetchAllEmployeeAndSendToPostgres()

	elapsedTime := time.Since(startTime)

	if err != nil {
		helpers.ErrorResponse(ctx, err.Error())
		return
	}

	response := map[string]interface{}{
		"message":       ":)",
		"elapsed_time":  elapsedTime.String(),
		"rows_inserted": insertedRows,
		"rows_deleted":  deletedRows,
	}

	helpers.APIResponse(ctx, "Success", http.StatusOK, response)
}
