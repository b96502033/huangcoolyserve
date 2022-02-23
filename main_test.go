package main

import (
	"bytes"
	"encoding/json"
	_ "huangcoolyserver/docs"
	"huangcoolyserver/src/controller"
	"huangcoolyserver/src/dto"
	"huangcoolyserver/src/paramDto"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setEnviroment() {
	os.Setenv("MONGO", "mongodb+srv://Hao:Hao12345678@huangcluster.fmtgp.mongodb.net")
}

func unSetEnviroment() {
	os.Unsetenv("MONGO")
}

func Test_setUserRouter(t *testing.T) {

	setEnviroment()

	defer unSetEnviroment()

	router := gin.Default()
	c := controller.NewController()
	router = setUserRouter(router, c)

	t.Run("getRoleAccessSystem", func(t *testing.T) {
		w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
		req, _ := http.NewRequest("GET", "/huangcoolyserver/user/getRoleAccessSystem/62046c4512cc08ba0fe17ae9", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var userLoginDto dto.UserLoginDto
		json.NewDecoder(w.Body).Decode(&userLoginDto)
		assert.Equal(t, userLoginDto.UserId, "62046c4512cc08ba0fe17ae9")

	})

	t.Run("login", func(t *testing.T) {
		w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
		paramUserLoginDto := &paramDto.ParamUserLoginDto{
			UserWorkId:   "admin",
			UserPassword: "admin",
		}

		js, _ := json.Marshal(paramUserLoginDto)
		req, _ := http.NewRequest("POST", "/huangcoolyserver/user/login", bytes.NewReader(js))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "62046c4512cc08ba0fe17ae9")
	})
}
