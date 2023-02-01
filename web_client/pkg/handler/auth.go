package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title":  "Авторизация",
		"logged": false,
	})
}

type LoginResponse struct {
	Status int64  `json:"status"`
	Token  string `json:"token"`
	Name   string `json:"name"`
}

func LoginRequest(email, password string) string {
	message := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	bytesReq, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.Post("http://192.168.17.247:9000/auth/login", "application/json", bytes.NewBuffer(bytesReq))

	if err != nil {
		log.Fatalln(err)
	}

	loginRes := LoginResponse{}

	err = json.NewDecoder(res.Body).Decode(&loginRes)

	if err != nil {
		log.Fatalln(err)
	}

	return loginRes.Token
}

func performLogin(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	var sameSiteString string

	if email != "" && password != "" {
		token := LoginRequest(email, password)

		if token != "" {
			ctx.SetCookie("token", token, 3600, "", sameSiteString, false, true)

			ctx.Redirect(http.StatusMovedPermanently, "/")
		} else {
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"ErrorTitle":   "Ошибка авторизации",
				"ErrorMessage": "Неверный логин и пароль",
			})
		}
	} else {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"ErrorTitle":   "Ошибка авторизации",
			"ErrorMessage": "Неверный логин и пароль",
		})
	}
}

func showRegPage(c *gin.Context) {
	c.HTML(http.StatusOK, "reg.html", gin.H{
		"title": "Регистрация",
	})
}

type RegResponse struct {
	Status int64 `json:"status"`
}

func RegRequest(name, email, password string) int64 {
	message := map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": password,
	}

	bytesReq, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.Post("http://192.168.17.247:9000/auth/register", "application/json", bytes.NewBuffer(bytesReq))

	if err != nil {
		log.Fatalln(err)
	}

	regRes := RegResponse{}

	err = json.NewDecoder(res.Body).Decode(&regRes)

	if err != nil {
		log.Fatalln(err)
	}

	return regRes.Status
}

func performReg(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if name != "" && email != "" && password != "" {
		statusReg := RegRequest(name, email, password)

		if statusReg == http.StatusCreated {
			showLoginPage(ctx)

			ctx.Redirect(http.StatusMovedPermanently, "/u/login")
		} else {
			ctx.HTML(http.StatusOK, "reg.html", gin.H{
				"ErrorTitle":   "Ошибка регистрации",
				"ErrorMessage": "Пользователь с таким email уже зарегистрирован",
			})
		}
	} else {
		ctx.HTML(http.StatusOK, "reg.html", gin.H{
			"ErrorTitle":   "Ошибка регистрации",
			"ErrorMessage": "Заполните все поля",
		})
	}
}
