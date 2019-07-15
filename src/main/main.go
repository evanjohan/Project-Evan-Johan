package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"main/configs"
	"main/controllers"
	"main/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Model Table

func main() {
	configs.ConnectingToMySQL()
	router := mux.NewRouter()
	router.HandleFunc("/News", CreateNews).Methods("POST")
	router.HandleFunc("/News", GetNews).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//functionCreateNews
func CreateNews(w http.ResponseWriter, r *http.Request) {
	var newsReq structs.NewsReq
	_ = json.NewDecoder(r.Body).Decode(&newsReq)
	news := structs.News{Author: newsReq.Author, Body: newsReq.Body}
	controllers.SendMessage(news)
	response := structs.Response{Message: "Send News Success: " + news.Author}
	json.NewEncoder(w).Encode(response)
}

//functionGetNews
func GetNews(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["page"]
	var newsGood structs.NewsResponse
	var arr_news []structs.NewsResponse
	var response structs.ResponseNews
	var count int
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	offset, err := strconv.Atoi(keys[0])
	if err != nil {
		panic(err)
	}
	limit := 10
	offset = (offset * 10) - 10
	db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	controllers.FailOnError(err, "Failed to connect DB mysql")
	rows, err := db.Raw("Select id,author,body from news order by id DESC limit ?, ?", offset, limit).Rows()
	if err != nil {
		log.Print(err)
	}
	rowsAll, err := db.Raw("Select count(*) from news").Rows()
	if err != nil {
		log.Print(err)
	}
	for rowsAll.Next() {
		if err := rowsAll.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	for rows.Next() {
		if err := rows.Scan(&newsGood.Id, &newsGood.Author, &newsGood.Body); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_news = append(arr_news, newsGood)
		}
	}
	countPage := count / limit
	if count%10 != 0 {
		countPage += 1
	}
	mess := "Success Get Data News"
	if len(arr_news) == 0 {
		count = 0
		countPage = 0
		mess = "Data news not found"
	}

	response.Status = 1
	response.Message = mess
	response.TotalData = count
	response.TotalPage = countPage
	response.Data = arr_news

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("Url Param 'key' is: " + string(keys[0]))
}
