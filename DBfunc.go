package main

import (
	"fmt"
	"os"
	"io"
	"database/sql"
	"time"
	"strings"
	"github.com/essentialkaos/translit"
	"mime/multipart"
)

func upload(file multipart.File, filename string,  path string ) string {

	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("."+path+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()
	io.Copy(f, file)
	return path+filename
}

/* set quantity 0 to get all news*/
func getNewsFromDB(quantity int, db *sql.DB) []News{
	var returnNews []News
	var rows *sql.Rows
	var err error
	if quantity>0 {
		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news LIMIT " + string(quantity) + "ORDER BY Date DESC")
	}else{
		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news ORDER BY Date DESC")
	}

	checkErr(err)
	var news News
	for rows.Next() {
		err = rows.Scan(&news.Title, &news.Body, &news.Img, &news.Date, &news.NewsLink)
		checkErr(err)
		returnNews = append(returnNews, news)
	}

	rows.Close()
	return returnNews
}


func getNewsByUrlFromDB(link string, db *sql.DB) News{
	var news News
	rows, err := db.Query("SELECT Title, Body, Img, Date, News_link FROM news WHERE News_link = '" + link +"'" )
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&news.Title, &news.Body, &news.Img, &news.Date, &news.NewsLink)
	}
	checkErr(err)
	rows.Close()
	return news
}

func addNewsToDB(db *sql.DB, imgPath string, title string, body string ) error {
	stmt, err := db.Prepare("INSERT INTO news(Title, Body, Img, Date, News_link) VALUES (?,?,?,?,?)")
	checkErr(err)
	t:= time.Now().Unix()
	str:= translit.EncodeToICAO(title)
	url := strings.Replace(str," ","-", -1)
	_, err = stmt.Exec(title, body, imgPath, int(t), url)
	return err
}

