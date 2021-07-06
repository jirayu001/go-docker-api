package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/api-server/model"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
}

func (h *CustomerHandler) Post(c *gin.Context) {
	var u model.DataUsers
	var r model.Result

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		fmt.Println(err.Error())
	} else {
		err := json.Unmarshal(jsonData, &u)
		if err != nil {
			fmt.Println(err.Error())
			r.Status = http.StatusInternalServerError
			r.Data.LiffURL = "null"
			r.Message = "กรุณากรอกข้อมูลให้ครบ"
			c.JSON(http.StatusInternalServerError, r)
		} else {
			//c.BindJSON(&u)
			if u.DisplayImage == "" || u.DisplayName == "" || u.Email == "" || u.LineToken == "" {
				r.Status = http.StatusInternalServerError
				r.Data.LiffURL = "null"
				r.Message = "กรุณากรอกข้อมูลให้ครบ"
				c.JSON(http.StatusInternalServerError, r)
			} else {
				r.Status = http.StatusOK
				r.Data.LiffURL = ""
				r.Message = ""
				c.JSON(http.StatusOK, r)
			}
		}
	}
}