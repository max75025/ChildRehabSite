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
func getNewsFromDB(quantity int, db *sql.DB) ([]News,error){
	var returnNews []News
	var rows *sql.Rows
	var err error
	if quantity>0 {
		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news LIMIT " + string(quantity) + "ORDER BY Date DESC")
	}else{
		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news ORDER BY Date DESC")
	}

	if err!= nil{
		return returnNews,err
	}
	var news News
	for rows.Next() {
		err = rows.Scan(&news.Title, &news.Body, &news.Img, &news.Date, &news.NewsLink)
		if err!= nil{
			return returnNews,err
		}
		returnNews = append(returnNews, news)
	}

	rows.Close()
	return returnNews, nil
}




func addNewsToDB(db *sql.DB, imgPath string, title string, body string ) error {
	stmt, err := db.Prepare("INSERT INTO news(Title, Body, Img, Date, News_link) VALUES (?,?,?,?,?)")
	t:= time.Now().Unix()
	str:= translit.EncodeToICAO(title)
	url := strings.Replace(str," ","-", -1)
	_, err = stmt.Exec(title, body, imgPath, int(t), url)
	return err
}

func deleteNewsFromDB(db *sql.DB, link string) error  {
	stmt, err := db.Prepare("DELETE FROM news WHERE News_link='?'")
	if err!= nil{
		return err
	}
	_, err = stmt.Exec(link)
	return err
}

func addAlbumToDB(db *sql.DB, title string, description string) error{
	stmt, err := db.Prepare("INSERT INTO Albums(Title, Description) VALUES (?,?)")
	if err!= nil{
		return err
	}
	_, err = stmt.Exec(title, description)
	return err
}

func deleteAlbumFromDB(db *sql.DB,idAlbum int) error{
	stmt, err := db.Prepare("DELETE FROM Albums WHERE ID=?")
	if err!= nil{
		return err
	}
	_, err = stmt.Exec(idAlbum)
	return err
}

func getAlbumsFromDB(db *sql.DB)([]PhotoAlbum, error){
		var albums []PhotoAlbum
		var album PhotoAlbum
		rows, err := db.Query("SELECT ID, Title,Description FROM Albums  " )
	if err!= nil{
		return albums,err
	}
		for rows.Next() {
		err = rows.Scan(&album.ID,&album.Title,&album.Description )
		albums = append(albums, album)
		}
	if err!= nil{
		return albums,err
	}
		rows.Close()
		return albums, nil
}

func addPhotoFromDB(db *sql.DB, title string, path string, albumName string) error{
	stmt, err := db.Prepare("INSERT INTO photo(Title, path, id_album) VALUES (?,?,?)")
	if err!= nil{
		return err
	}
	_, err = stmt.Exec(title, path, albumName)
	return err
}

func deletePhotoFromDB(db *sql.DB,id int) error{
	stmt, err := db.Prepare("DELETE FROM photo WHERE ID=?")
	if err!= nil{
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func getPhotoFromDBByID(db *sql.DB, id int )(Photo){
	var photo Photo
	 db.QueryRow("SELECT ID, Title, path, id_album FROM photo WHERE id = ?", id).Scan(photo.ID,photo.Title, photo.Img, photo.IDAlbum )
	 return photo
}


func getPhotosFromDB(db *sql.DB)([]Photo, error){
	var photos []Photo
	var photo Photo
	rows, err := db.Query("SELECT ID, title, path, id_album FROM photo  " )
	if err!= nil{
		return photos,err
	}
	for rows.Next() {
		err = rows.Scan(&photo.ID,&photo.Title,&photo.Img, &photo.IDAlbum )
		photos = append(photos, photo)
	}
	if err!= nil{
		return photos,err
	}
	rows.Close()
	return photos, nil
}

