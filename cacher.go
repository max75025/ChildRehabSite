package main

import (
	"database/sql"
	"errors"

	"time"
	"strconv"

)


var cachedAlbums 	[]PhotoAlbum
var cachedNews 		[]News


func refreshCache(db *sql.DB){
	cacheAlbums(db)
	cacheNews(db)
}


func cacheAlbums(db *sql.DB)  {
	 Albums := getAlbumsFromDB(db)
	allPhotos:=getPhotosFromDB(db)


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
		i,_:=strconv.ParseInt(news.Date, 10, 64)
		t:=time.Unix(i,0)
		convTime:= t.Format("02.01.2006")
		news.Date = convTime
		allNews = append(allNews, news)
	}
	rows.Close()
	cachedNews = allNews
}

func getAlbumById(id int)(PhotoAlbum,error){
	for _,k:= range cachedAlbums{
		if k.ID==id{
			return k,nil
		}
	}
	return PhotoAlbum{},errors.New("not found")
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

