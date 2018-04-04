package main

type BasicPageData struct {
	Description string
	Keywords    string
	Title       string
}

type IndexDataModel struct {
	Meta       BasicPageData
	OurMission Article
	OurTeam    []Article
	OurHistory Article
}

type Article struct {
	Title       string
	Body        string
	Img         string
	ArticleLink string
}

/*type PreviewArticle Article{

}*/
