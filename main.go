// ChildrensPortal project main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"


	"strconv"
	"log"
)


/*func checkErr(err error,) {
	if err != nil {
		panic(err)

	}
}*/

func checkSiteErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// upload one img and return path or ""


func main() {
	db, err:= sql.Open("sqlite3", "./database/db.db")
	if err!= nil{
		log.Println(err)
		return
	}
	router := httprouter.New()

	//Serve static
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.ServeFiles("/newsImgs/*filepath", http.Dir("newsImgs"))
	router.ServeFiles("/photo/*filepath", http.Dir("photo"))


	//Views
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := new(IndexDataModel)
		newArticle := Article{Title: "Тест", Img: "/static/none.png", Body: `Психолог`}
		data.OurTeam = append(data.OurTeam, newArticle, newArticle, newArticle, newArticle)
		newArticle = Article{Title: "Наша миссия", Img: "/static/bg.png", Body: `Есть много вариантов Lorem Ipsum, но большинство из них имеет не всегда приемлемые модификации, например, юмористические вставки или слова, которые даже отдалённо который они просто повторяют, пока не достигнут нужный объём. Это делает предлагаемый здесь генератор единственным настоящим Lorem Ipsum генератором. Он использует словарь из более чем 200 латинских слов, а также набор моделей предложений. В результате сгенерированный Lorem Ipsum выглядит правдоподобно, не имеет повторяющихся абзацей или "невозможных" слов.`}
		data.OurMission = newArticle
		newArticle = Article{Title: "Наша история", Img: "/static/foot.jpg", Body: `Многие думают, что Lorem Ipsum - но это не совсем так. Его корни уходят в один фрагмент классической латыни 45 года н.э., то есть более двух тысячелетий назад. Ричард МакКлинток, и занялся его поисками в классической латинской литературе. В результате он нашёл неоспоримый первоисточник Lorem Ipsum в разделах 1.10.32 и 1.10.33 книги "de Finibus Bonorum et Malorum" ("О пределах добра и зла"), написанной Цицероном в 45 году н.э. Этот трактат по теории этики был очень популярен в эпоху Возрождения. Первая строка Lorem Ipsum, "Lorem ipsum dolor sit amet..", происходит от одной из строк в разделе 1.10.32`}
		data.OurHistory = newArticle
		tmpl := template.Must(template.ParseFiles("tmpls/index.html"))
		tmpl.Execute(w, data)
	})

	router.GET("/news", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		data:= new(NewsDataModel)

		// var manyNews []News
		// manyNews = append(manyNews, newPost, newPost)

		data.AllNews = cachedNews
		tmpl:= template.Must(template.ParseFiles("tmpls/news.html"))
		tmpl.Execute(w,data)
	})
	router.GET("/news/:link", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		data:= new(NewsDataModel)
		link :=  ps.ByName("link")
		news,err:= getNewsByUrl(link)
		if err!= nil{
			http.Redirect(w,r,"/404", http.StatusSeeOther)
		}
		data.AllNews = append(data.AllNews,news )
		tmpl:= template.Must(template.ParseFiles("tmpls/oneNews.html"))
		tmpl.Execute(w,data)
	})

	router.GET("/albums", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		data:= new(PhotosDataModel)

		data.Albums = cachedAlbums
		tmpl:= template.Must(template.ParseFiles("tmpls/Albums.html"))
		tmpl.Execute(w,data)
	})
	router.GET("/albums/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		id,err:=strconv.Atoi(ps.ByName("id"))
		if err!= nil {
			log.Println(err)
			http.Error(w, "id not int", 500)
			return
		}
		album,err:= getAlbumById(id)
		if err!= nil {
			log.Println(err)
			http.Error(w, "not found album by id", 500)
			return
		}

		tmpl:= template.Must(template.ParseFiles("tmpls/photos.html"))
		tmpl.Execute(w,album)
	})

	router.GET("/history", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := new(OurHistoryDataModel)
		article1:=Article{
			Title: "Подзаголовок",
			Img:   "/static/350x200.png",
			Body:  `Lorem ipsum dolor sit amet, consectetur adipisicing elit.  Accusamus architecto asperiores at eaque eligendi hic maxime mollitia nam, neque nisi optio placeat praesentium provident quaerat quam quod reiciendis sit tenetur veritatis vitae. Aspernatur assumenda corporis culpa cumque dignissimos dolores eaque et explicabo facilis, fugiat in ipsa iure laboriosam molestiae neque non numquam officiis omnis optio possimus praesentium provident quas qui quidem ratione, recusandae rerum sunt temporibus ullam voluptatibus. `,
		}
		article2:=Article{
			Title: "Подзаголовок",
			Img:   "",
			Body:  `Lorem ipsum dolor sit amet, consectetur adipisicing elit.  Accusamus architecto asperiores at eaque eligendi hic maxime mollitia nam, neque nisi optio placeat praesentium provident quaerat quam quod reiciendis sit tenetur veritatis vitae. Aspernatur assumenda corporis culpa cumque dignissimos dolores eaque et explicabo facilis, fugiat in ipsa iure laboriosam molestiae neque non numquam officiis omnis optio possimus praesentium provident quas qui quidem ratione, recusandae rerum sunt temporibus ullam voluptatibus.            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab assumenda cupiditate distinctio doloremque eligendi ex expedita, explicabo facilis harum in inventore modi nam nemo nesciunt non nulla numquam officiis perferendis quaerat quasi quia sit suscipit ullam. Amet autem consequuntur debitis ea hic inventore magnam mollitia odio officiis, pariatur placeat porro quidem quisquam reprehenderit unde ut voluptate! Aliquid amet deleniti doloremque enim? Architecto aspernatur atque cum debitis eos error eum explicabo id illum ipsam ipsum iure iusto laborum maxime natus necessitatibus odio possimus, rem repellat, sit soluta sunt unde vero! Animi cupiditate dolore doloremque et explicabo incidunt neque, perspiciatis provident quas! `,
		}
		data.History = append(data.History, article1, article2, article1)
		tmpl := template.Must(template.ParseFiles("tmpls/history.html"))
		tmpl.Execute(w, data)
	})

	router.GET("/team", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//fmt.Fprintf(w," our team")
		data:= new(TeamDataModel)
		specialist:= Specialist{
			Position:    "психолог",
			Img:         "http://placehold.it/200x200",
			Name:        "Иванов Иван Иваныч",
			Discription: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eos molestias quas animi architecto rem consequatur quia nesciunt, culpa? Recusandae error molestias impedit dolorem. Nisi dicta reprehenderit, ut minus cupiditate architecto excepturi iure esse corrupti hic iste, unde ad incidunt fuga ex facere minima maxime rerum sint consequatur fugit natus debitis!Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eos molestias quas animi architecto rem consequatur quia nesciunt, culpa? Recusandae error molestias impedit dolorem. Nisi dicta reprehenderit, ut minus cupiditate architecto excepturi iure esse corrupti hic iste, unde ad incidunt fuga ex facere minima maxime rerum sint consequatur fugit natus debitis!",
			Link:        "/link",
		}
		data.OurTeam= append(data.OurTeam, specialist, specialist, specialist, specialist)
		tmpl:=template.Must(template.ParseFiles("tmpls/team.html"))
		tmpl.Execute(w,data)
	})

	router.GET("/doc", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data:= new (DocumentsDataModel)
		doc:= Document{
			Name:        "Document",
			Description: "",
			Link:        "document.txt",
		}
		data.Documents= append(data.Documents,doc, doc,doc,doc,doc )
		tmpl:=template.Must(template.ParseFiles("tmpls/documents.html"))
		tmpl.Execute(w,data)
		//fmt.Fprintf(w," our documents")
	})
	router.GET("/contacts", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tmpl:=template.Must(template.ParseFiles("tmpls/contacts.html"))
		tmpl.Execute(w,"")
		//fmt.Fprintf(w," contacts")
	})


	/************************************admin panel and auth*******************************************/


	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		/*name := "admin"
		password :="admin"*/
			t, _ := template.ParseFiles("tmpls/login.html")
			t.Execute(w, nil)
	})
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//name := "admin"
		//password :="admin"
		/*if r.Form["username"][0] == name && r.Form["password"][0] == password {
			fmt.Fprintf(w, "success")
			//
		}*/
		name := r.FormValue("username")
		pass := r.FormValue("password")
		redirectTarget := "/404"
		if name != "" && pass != "" && checkLoginData(name, pass) {
			// .. check credentials ..
			setSession(name, w)
			redirectTarget = "/AdminPanel"
		}
		http.Redirect(w, r, redirectTarget, 302)
	} )

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logoutHandler(w,r)
	})

	router.GET("/AdminPanel" ,func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		checkLogin(w, r)
		//fmt.Fprintf(w,"hello %s", getUserName(r))
		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/AdminPanel.html"))
		tmpl.Execute(w,nil)
	})

	router.GET("/AdminPanel/news", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		news:= new(NewsDataModel)
		news.AllNews = cachedNews
		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/news.html"))
		tmpl.Execute(w,news)
	})
	router.GET("/AdminPanel/addNews", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w,r)
		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/addNews.html"))
		tmpl.Execute(w,nil)
		//fmt.Println(time.Now().Unix())
	})

	router.POST("/AdminPanel/addNews", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w,r)

		r.ParseForm()
		imgPath:= ""
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("img")
		if err != nil {
			log.Println(err)
			return
		}
		imgPath = upload(file, handler.Filename, "/newsImgs/")

		err = addNewsToDB(db, imgPath, r.Form["title"][0],r.Form["body"][0] )
		if err!=nil {
			fmt.Fprintf(w, "введены некоректные или ранее используемые данные")
			return
		}
		fmt.Println("news add")
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/news", http.StatusSeeOther)
		//fmt.Println(time.Unix(t,0).Format("02.01.2006"))

	})

	router.GET("/AdminPanel/deleteNews/:link", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w,r)
		_, err:= db.Exec("Delete FROM news WHERE News_link = '" + ps.ByName("link") + "'")
		if err!= nil {
			log.Println(err)
			http.Error(w, "error delete news", 500)
			return
		}
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/news", http.StatusSeeOther)


	})


	router.GET("/AdminPanel/Albums", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w,r)
		albums := cachedAlbums
		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/PhotoAlbums.html"))
		tmpl.Execute(w,albums)
	})

	router.GET("/AdminPanel/addAlbums", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//checkLogin(w,r)

		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/addAlbums.html"))
		tmpl.Execute(w,nil)
	})
	router.POST("/AdminPanel/Albums/add", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w,r)
		r.ParseForm()
		//err = addAlbumToDB(db, "test","test" )
		err = addAlbumToDB(db, r.Form["title"][0],r.Form["description"][0] )
		if err!=nil {
			fmt.Fprintf(w, "введены некоректные или ранее используемые данные")
			return
		}
		fmt.Println("album add")
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/Albums", http.StatusSeeOther)
	})
	router.GET("/AdminPanel/Album/:id/delete",func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//checkLogin(w,r)
		i,_:=strconv.Atoi(ps.ByName("id"))
		err = deleteAlbumFromDB(db, i)
		if err!= nil {
			log.Println(err)
			http.Error(w, "error delete album", 500)
			return
		}
		fmt.Println("album delete")
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/Albums", http.StatusSeeOther)
	})

	router.GET("/AdminPanel/Albums/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		id,err:=strconv.Atoi(ps.ByName("id"))
		if err!= nil {
			log.Println(err)
			http.Error(w, "id not int", 500)
			return
		}
        album,err:= getAlbumById(id)
		if err!= nil {
			log.Println(err)
			http.Error(w, "not found album by id", 500)
			return
		}

		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/photos.html"))
		tmpl.Execute(w,album)
	})
	router.GET("/AdminPanel/Albums/:id/addPhoto", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {


		tmpl:= template.Must(template.ParseFiles("tmpls/AdminPanel/addPhoto.html"))
		tmpl.Execute(w,ps.ByName("id"))
	})

	router.POST("/AdminPanel/Albums/addPhoto/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		imgPath:= ""
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("img")
		if err != nil {
			log.Println(err)
			return
		}
		imgPath = upload(file, handler.Filename, "/photo/")
		err = addPhotoFromDB(db,r.Form["title"][0],imgPath, ps.ByName("id"))
		if err!=nil {
			fmt.Fprintf(w, "введены некоректные или ранее используемые данные")
			return
		}
		fmt.Println("add photo")
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/Albums/"+ps.ByName("id"), http.StatusSeeOther)
	})

	router.GET("/AdminPanel/deletePhoto/:idImg/:idAlbum", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idImg,_:= strconv.Atoi(ps.ByName("idImg"))
		deletePhotoFromDB(db, idImg)
		refreshCache(db)
		http.Redirect(w, r, "/AdminPanel/Albums/"+ps.ByName("idAlbum"), http.StatusSeeOther)
	})







	/*some test */
	/*router.GET("/db", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rows, err:= db.Query("SELECT * FROM news")
		checkErr(err)
		var id int
		var title string
		var body string
		var img string
		var link string
		var date string

		for rows.Next(){
			err = rows.Scan(&id, &title, &body, &img ,& link, &date)
			checkErr(err)
			fmt.Fprintf(w,"%d\n",id)
			fmt.Fprintf(w,"%s\n",title)
			fmt.Fprintf(w,"%s\n",body)
			fmt.Fprintf(w,"%s\n",img)
			fmt.Fprintf(w,"%s\n",link)
			fmt.Fprintf(w,"%s\n",date)
		}
	})*/


	/*router.GET("/socialJson", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		testLink:= SocialLinks{
			Facebook: "fb.com",
			Twitter:  "twitter.com",
			Telegram: "telegram.com",
			Vk:       "vk.com",
			Ok:       "ok.com",
			YouTube:  "YouTube.com",
		}
		testJson,_:= json.Marshal(testLink)
		testDecodeJson := SocialLinks{}
		json.Unmarshal(testJson, &testDecodeJson)
		fmt.Fprintf(w,"%s\n",testLink)
		fmt.Fprintf(w,"%s\n",testJson)
		fmt.Fprintf(w,"%s\n",testDecodeJson)
	})*/






	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
		fmt.Fprintf(w,"answer: 42")

	})



	//Iit Server
	server := http.Server{
		Addr:         ":8000",
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		Handler:      router,
	}
	refreshCache(db)
	go autoRefreshCache(db)
	fmt.Println("server listen and serve...")
	fmt.Println(server.ListenAndServe())

	db.Close()

}






