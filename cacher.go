package main

import (
	"database/sql"
	"errors"

	"time"
	"strconv"

	"fmt"
	"log"
)


var cachedAlbums 	[]PhotoAlbum
var cachedNews 		[]News

func autoRefreshCache(db *sql.DB) {
	for  range time.Tick(30 *time.Second){
		fmt.Println("autorefresh cache...")
		err := refreshCache(db)
		if err!= nil{
			log.Println(err)

		}
	}
}

func refreshCache(db *sql.DB) error{
	err:= cacheAlbums(db)
	if err!= nil{
		return err
	}
	err = cacheNews(db)
	if err!= nil{
		return err
	}
	return nil
}


func cacheAlbums(db *sql.DB) error {
	 Albums, err := getAlbumsFromDB(db)
	if err!= nil{
		return err
	}
	allPhotos , err:=getPhotosFromDB(db)
	if err!= nil{
		return err
	}

	for i,k:= range Albums {
		var photosForAlbum []Photo
		for _,t := range allPhotos{

			if k.ID == t.IDAlbum {
				photosForAlbum = append(photosForAlbum, t)
			}
		}

		Albums[i].Img = photosForAlbum
		if len(photosForAlbum)!=0{
			Albums[i].Cover = photosForAlbum[0].Img
		}else{
			Albums[i].Cover = "http://placehold.it/400x300"
		}

	}
	cachedAlbums = Albums
	return nil
}

func cacheNews(db *sql.DB) error{
	var allNews []News
	var rows *sql.Rows
	var err error

		rows, err = db.Query("SELECT Title, Body, Img, Date, News_link FROM news ORDER BY Date DESC")

	if err!= nil{
		return err
	}
	var news News
	for rows.Next() {
		err = rows.Scan(&news.Title, &news.Body, &news.Img, &news.Date, &news.NewsLink)
		if err!= nil{
			return err
		}
		i,_:=strconv.ParseInt(news.Date, 10, 64)
		t:=time.Unix(i,0)
		convTime:= t.Format("02.01.2006")
		news.Date = convTime
		allNews = append(allNews, news)
	}
	rows.Close()
	cachedNews = allNews
	return nil
}

func getAlbumById(id int)(PhotoAlbum,error){
	for _,k:= range cachedAlbums{
		if k.ID==id{
			return k,nil
		}
	}
	return PhotoAlbum{},errors.New("not found album in cache")
}


func getNewsByUrl(link string) (News,error){
	for _,k:=range cachedNews{
		if k.NewsLink == link{
			return k, nil
		}
	}

	return News{},errors.New("not found news by id")
}





