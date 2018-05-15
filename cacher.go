package main

import (
	"database/sql"

	"encoding/json"
	"fmt"
)



var cachedAlbums 	[]PhotoAlbum
var cachedNews 		[]News


func cacheAlbums(db *sql.DB)  {
	var allAlbum []PhotoAlbum
	var album PhotoAlbum

	var id int
	var path string
	var title string
	var albumStr string

	currentAlbum:= ""
	rows, err := db.Query("SELECT * FROM photo GROUP BY album" )
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&id, &path, &title, &albumStr)
		photo := Photo{
			ID:    id,
			Img:   path,
			Title: title,
		}
		if(currentAlbum!=albumStr){

		}

		allAlbum = append(allAlbum, album)
	}
	checkErr(err)
	rows.Close()
	fmt.Println("albums cashed")
}

func cacheNews(db *sql.DB){
	var allNews []News
	var rows *sql.Rows
	var err error

		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news ORDER BY Date DESC")

	checkErr(err)
	var news News
	for rows.Next() {
		err = rows.Scan(&news.Title, &news.Body, &news.Img, &news.Date, &news.NewsLink)
		checkErr(err)
		allNews = append(allNews, news)
	}
	rows.Close()
	cachedNews = allNews
}

/*func testInputPhotoToDB(db *sql.DB){
	stmt, err := db.Prepare("INSERT INTO photo(path,title) VALUES (?,?)")
	checkErr(err)

	ex, err := stmt.Exec("test","test")
	checkErr(err)
	fmt.Println("add test photo")
	fmt.Println(ex)
}*/

/*func testInputJsonToDB(db *sql.DB){
	stmt, err := db.Prepare("INSERT INTO json_object(json, type) VALUES (?,?)")
	checkErr(err)

	album:= PhotoAlbum{
		Img:         []string{"test","test", "test"},
		Title:       "test",
		Description: "test",
	}
	jsonStr,err:=json.Marshal(album)
	checkErr(err)
	_, err = stmt.Exec(jsonStr,"album")
	checkErr(err)
	fmt.Println("add test json album")
}*/

