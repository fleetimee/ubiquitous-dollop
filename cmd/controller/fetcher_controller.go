package controller

import (
	"service-fleetime/cmd/helpers"
	"service-fleetime/cmd/repository"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {

	email := ctx.DefaultQuery("email", "")

	if email == "" {
		helpers.ErrorResponse(ctx, "Email is required")
		return
	}

	if !helpers.IsBpddiyEmail(email) {
		helpers.ErrorResponse(ctx, "Email tidak valid")
		return
	}

	fetcher, err := repository.FetchByEmail(email)

	if fetcher.Nrp == 0  {
		helpers.ErrorResponse(ctx, "Data pegawai tidak ditemukan")

		return
	}

	

	if err != nil {
		helpers.ErrorResponse(ctx, err.Error())

		return
	}

	helpers.APIResponse(ctx, "Success", 200, fetcher)

}
