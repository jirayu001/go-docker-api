package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/api-server/model"
	"github.com/gin-gonic/gin"
)

// ShowResult godoc
// @Summary Show an result
// @Description show by json result
// @Tags result
// @Accept  json
// @Produce  json
// @Param account body model.DataUsers true "Show result"
// @Success 200 {object} model.Result
// @Failure 500 {object} model.Result
// @Router /auth/loginWithCoulacare [post]
func (c *Controller) ShowResult(ctx *gin.Context) {
	var u model.DataUsers
	var r model.Result

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err := json.Unmarshal(jsonData, &u)
		if err != nil {
			fmt.Println(err.Error())
			r.Status = http.StatusInternalServerError
			r.Data.LiffURL = "null"
			r.Message = "กรุณากรอกข้อมูลให้ครบ"
			ctx.JSON(http.StatusInternalServerError, r)
		} else {
			if u.Token == "" {
				r.Status = http.StatusInternalServerError
				r.Data.LiffURL = "null"
				r.Message = "กรุณากรอกข้อมูลให้ครบ"
				ctx.JSON(http.StatusInternalServerError, r)
			} else {
				r.Status = http.StatusOK
				r.Data.LiffURL = "https://f0ac1522c75e.ngrok.io"
				r.Message = ""
				ctx.JSON(http.StatusOK, r)
			}
		}
	}
}
