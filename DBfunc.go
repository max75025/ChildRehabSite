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

func deleteNewsFromDB(db *sql.DB, link string) error  {
	stmt, err := db.Prepare("DELETE FROM news WHERE News_link='?'")
	checkErr(err)
	_, err = stmt.Exec(link)
	return err
}

func addAlbumToDB(db *sql.DB, title string, description string) error{
	stmt, err := db.Prepare("INSERT INTO Albums(Title, Description) VALUES (?,?)")
	checkErr(err)
	_, err = stmt.Exec(title, description)
	return err
}

func deleteAlbumFromDB(db *sql.DB,idAlbum int) error{
	stmt, err := db.Prepare("DELETE FROM Albums WHERE ID=?")
	checkErr(err)
	_, err = stmt.Exec(idAlbum)
	return err
}

func getAlbumsFromDB(db *sql.DB)([]PhotoAlbum){
		var albums []PhotoAlbum
		var album PhotoAlbum
		rows, err := db.Query("SELECT ID, Title,Description FROM Albums  " )
		checkErr(err)
		for rows.Next() {
		err = rows.Scan(&album.ID,&album.Title,&album.Description )
		albums = append(albums, album)
		}
		checkErr(err)
		rows.Close()
		return albums
}

func addPhotoFromDB(db *sql.DB, title string, path string, albumName string) error{
	stmt, err := db.Prepare("INSERT INTO photo(Title, path, id_album) VALUES (?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(title, path, albumName)
	return err
}

func deletePhotoFromDB(db *sql.DB,id int) error{
	stmt, err := db.Prepare("DELETE FROM photo WHERE ID=?")
	checkErr(err)
	_, err = stmt.Exec(id)
	return err
}


func getPhotosFromDB(db *sql.DB)([]Photo){
	var photos []Photo
	var photo Photo
	rows, err := db.Query("SELECT ID, title, path, id_album FROM photo  " )
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&photo.ID,&photo.Title,&photo.Img, &photo.IDAlbum )
		photos = append(photos, photo)
	}
	checkErr(err)
	rows.Close()
	return photos
}

