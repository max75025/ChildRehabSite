// ChildrensPortal project main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	//Serve static
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	//Views
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := new(IndexDataModel)
		newArticle := Article{Title: "Тест", Img: "/static/none.png", Body: `Психолог`}
		data.OurTeam = append(data.OurTeam, newArticle, newArticle, newArticle, newArticle)
		newArticle = Article{Title: "Наша миссия", Img: "/static/hands.png", Body: `Есть много вариантов Lorem Ipsum, но большинство из них имеет не всегда приемлемые модификации, например, юмористические вставки или слова, которые даже отдалённо который они просто повторяют, пока не достигнут нужный объём. Это делает предлагаемый здесь генератор единственным настоящим Lorem Ipsum генератором. Он использует словарь из более чем 200 латинских слов, а также набор моделей предложений. В результате сгенерированный Lorem Ipsum выглядит правдоподобно, не имеет повторяющихся абзацей или "невозможных" слов.`}
		data.OurMission = newArticle
		newArticle = Article{Title: "Наша история", Img: "/static/foot.jpg", Body: `Многие думают, что Lorem Ipsum - но это не совсем так. Его корни уходят в один фрагмент классической латыни 45 года н.э., то есть более двух тысячелетий назад. Ричард МакКлинток, и занялся его поисками в классической латинской литературе. В результате он нашёл неоспоримый первоисточник Lorem Ipsum в разделах 1.10.32 и 1.10.33 книги "de Finibus Bonorum et Malorum" ("О пределах добра и зла"), написанной Цицероном в 45 году н.э. Этот трактат по теории этики был очень популярен в эпоху Возрождения. Первая строка Lorem Ipsum, "Lorem ipsum dolor sit amet..", происходит от одной из строк в разделе 1.10.32`}
		data.OurHistory = newArticle
		tmpl := template.Must(template.ParseFiles("tmpls/index.html"))
		tmpl.Execute(w, data)
	})
	//Iit Server
	server := http.Server{
		Addr:         ":8888",
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		Handler:      router,
	}
	fmt.Println(server.ListenAndServe())
}


