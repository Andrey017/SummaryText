package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getHome(ctx *gin.Context) {
	token, _ := ctx.Cookie("token")

	if token != "" {
		articles := loadArticle(token)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Главная страница",
			"logged":  true,
			"payload": articles.Data,
		})
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/u/login")
	}
}

type ArticleAll struct {
	Status int `json:"status"`
	Data   []struct {
		Id         int64  `json:"id"`
		Name       string `json:"name"`
		FileName   string `json:"fileName"`
		Note       string `json:"note"`
		DateCreate string `json:"dateCreate"`
		UserId     int64  `json:"userId"`
	} `json:"data"`
}

func loadArticle(token string) ArticleAll {
	articleList := ArticleAll{}

	var bearer = "Bearer " + token

	req, err := http.NewRequest("GET", "http://192.168.17.247:9000/article/", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &articleList)

	if err != nil {
		log.Fatalln(err)
	}

	return articleList
}

func showCreateArticle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create_article.html", gin.H{
		"title":  "Создать",
		"logged": true,
	})
}

func createArticleInfo(ctx *gin.Context) {
	name := ctx.PostForm("name")
	note := ctx.PostForm("note")
	date := ctx.PostForm("dateCreate")

	dateSplit := strings.Split(date, "-")

	date = dateSplit[2] + "." + dateSplit[1] + "." + dateSplit[0]

	if name != "" {
		token, _ := ctx.Cookie("token")

		status, id := CreateArticleInfoReq(name, note, date, token)

		if status == http.StatusCreated {
			ctx.Redirect(http.StatusMovedPermanently, "/api/create/loadFile/"+strconv.Itoa(id))
		} else {
			showCreateArticle(ctx)
		}
	}
}

type CreateArticleResponse struct {
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}

func CreateArticleInfoReq(name, note, date, token string) (int64, int) {
	message := map[string]interface{}{
		"name":        name,
		"file_name":   "",
		"note":        note,
		"date_create": date,
	}

	bytesReq, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article")

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bytesReq))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	createRes := CreateArticleResponse{}

	err = json.NewDecoder(res.Body).Decode(&createRes)

	if err != nil {
		log.Fatalln(err)
	}

	return createRes.Status, int(createRes.Id)
}

func showLoadFileArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	ctx.HTML(http.StatusOK, "load_file.html", gin.H{
		"title":  "Загрузить",
		"logged": true,
		"Id":     id,
	})
}

type Article struct {
	Status int `json:"status"`
	Data   struct {
		Id         int64  `json:"id"`
		Name       string `json:"name"`
		FileName   string `json:"fileName"`
		Note       string `json:"note"`
		DateCreate string `json:"dateCreate"`
		UserId     int64  `json:"userId"`
	} `json:"data"`
}

func showArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	token, _ := ctx.Cookie("token")

	if token != "" {
		article := loadArticleById(token, id)
		summary := loadSummary(token, strconv.Itoa(int(article.Data.Id)))

		ctx.HTML(http.StatusOK, "article.html", gin.H{
			"title":   "Страница статьи",
			"logged":  true,
			"payload": article.Data,
			"summ":    summary.Data,
		})
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func loadArticleById(token, id string) Article {
	articleList := Article{}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article/%s", id)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &articleList)

	if err != nil {
		log.Fatalln(err)
	}

	return articleList
}

type DeleteArticleResponse struct {
	Status int64 `json:"status"`
}

func deleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	token, _ := ctx.Cookie("token")

	if token != "" {
		result := deleteArticleReq(token, id)

		if result.Status == http.StatusOK {
			ctx.Redirect(http.StatusMovedPermanently, "/")
		}
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func deleteArticleReq(token, id string) DeleteArticleResponse {
	result := DeleteArticleResponse{}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article/%s", id)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func showUpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.HTML(http.StatusOK, "update_article.html", gin.H{
		"title":  "Обновить",
		"logged": true,
		"Id":     id,
	})
}

func updateArticleInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	note := ctx.PostForm("note")

	if name != "" {
		token, _ := ctx.Cookie("token")

		status, _ := UpdateArticleInfoReq(id, name, note, token)

		if status == http.StatusOK {
			ctx.Redirect(http.StatusMovedPermanently, "/")
		} else {
			ctx.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

type UpdateArticleResponse struct {
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}

func UpdateArticleInfoReq(id, name, note, token string) (int64, int) {
	message := map[string]interface{}{
		"name":        name,
		"file_name":   "",
		"note":        note,
		"date_create": "",
	}

	bytesReq, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article/%s", id)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bytesReq))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	updateRes := UpdateArticleResponse{}

	err = json.NewDecoder(res.Body).Decode(&updateRes)

	if err != nil {
		log.Fatalln(err)
	}

	return updateRes.Status, int(updateRes.Id)
}

//----------------------------------------------------------------------------------------------------------------------------------

type SummaryAll struct {
	Status int `json:"status"`
	Data   []struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		NameFile  string `json:"nameFile"`
		Type      string `json:"type"`
		Note      string `json:"note"`
		ArticleId int64  `json:"articleId"`
	} `json:"data"`
}

func loadSummary(token, id string) SummaryAll {
	summaryList := SummaryAll{}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article/summary/%s", id)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &summaryList)

	if err != nil {
		log.Fatalln(err)
	}

	return summaryList
}

func showCreateSummary(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.HTML(http.StatusOK, "create_summary.html", gin.H{
		"title":  "Создать",
		"logged": true,
		"Id":     id,
	})
}

func createSummary(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	typeSumm := ctx.PostForm("type")
	note := ctx.PostForm("note")

	countWord, _ := strconv.ParseInt(ctx.PostForm("wordCount"), 10, 32)
	countSentense, _ := strconv.ParseInt(ctx.PostForm("sentenseCount"), 10, 32)
	minCountSentense, _ := strconv.ParseInt(ctx.PostForm("minsentenseCount"), 10, 32)
	maxCountSentense, _ := strconv.ParseInt(ctx.PostForm("maxsentenseCount"), 10, 32)

	fmt.Println(countWord)

	if name != "" {
		token, _ := ctx.Cookie("token")

		fileName := loadArticleById(token, id)

		articleId, _ := strconv.ParseInt(id, 10, 32)

		id := CreateSummaryReq(name, typeSumm, note, fileName.Data.FileName, articleId, token, countWord, countSentense, minCountSentense, maxCountSentense)

		fmt.Println(id)

		if id != 0 {
			ctx.Redirect(http.StatusMovedPermanently, "/api/article/"+strconv.Itoa(int(articleId)))
		} else {
			showCreateSummary(ctx)
		}
	}
}

type CreateSummaryResponse struct {
	Id int64 `json:"id"`
}

func CreateSummaryReq(name, typeSumm, note, fileName string, articleId int64, token string, countWord, countSentense, minCountSentense, maxCountSentense int64) int {
	message := map[string]interface{}{
		"name":             name,
		"file_name":        fileName,
		"type":             typeSumm,
		"note":             note,
		"article_id":       articleId,
		"countWord":        countWord,
		"countSentense":    countSentense,
		"minCountSentense": minCountSentense,
		"maxCountSentese":  maxCountSentense,
	}

	fmt.Println(message)

	bytesReq, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + token

	var url string

	if typeSumm == "Создать короткую анотацию" {
		url = fmt.Sprintf("http://192.168.17.247:9000/summaryservice/rank")
	} else if typeSumm == "Создать резюме" {
		url = fmt.Sprintf("http://192.168.17.247:9000/summaryservice/lsa")
	} else if typeSumm == "Создать реферат" {
		url = fmt.Sprintf("http://192.168.17.247:9000/summaryservice/t5")
	} else {
		return 0
	}

	fmt.Println(url)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bytesReq))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	createRes := CreateSummaryResponse{}

	err = json.NewDecoder(res.Body).Decode(&createRes)

	if err != nil {
		log.Fatalln(err)
	}

	return int(createRes.Id)
}

type DeleteSummaryResponse struct {
	Status    int64 `json:"status"`
	ArticleId int64 `json:"articleId"`
}

func deleteSummary(ctx *gin.Context) {
	id := ctx.Param("id")
	token, _ := ctx.Cookie("token")

	if token != "" {
		result := deleteSummaryReq(token, id)

		if result.Status == http.StatusOK {
			ctx.Redirect(http.StatusMovedPermanently, "/api/article/"+strconv.Itoa(int(result.ArticleId)))
		}
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func deleteSummaryReq(token, id string) DeleteSummaryResponse {
	result := DeleteSummaryResponse{}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.247:9000/article/summary/%s", id)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatalln(err)
	}

	return result
}
