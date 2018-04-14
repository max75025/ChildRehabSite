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


type NewsDataModel struct{
	//LatestNews []Article
	Meta BasicPageData
	AllNews []News
}


type News struct{
	Title 		string
	Body 		string
	Img 		string
	Date 		string
	NewsLink 	string
}

type PhotoAlbum struct{
	Img         []string
	Title 		string
	Description string
}

type PhotosDataModel struct{
	Meta 	BasicPageData
	Albums 	[]PhotoAlbum
}

type TeamDataModel struct{
	Meta 		BasicPageData
	OurTeam 	[]Article
}


type ContactsDataModel struct{
	Meta 			BasicPageData
	PhoneNumbers 	[]string
	Emails 			[]string
	Address 		[]string
	Social			SocialLinks

}

type SocialLinks struct{
	Facebook string
	Twitter  string
	Telegram string
	Vk       string
	Ok 		 string
	YouTube  string
}

type DocumentsDataModel struct {
	Meta BasicPageData
	Documents Document
}

type Document struct{
	Name string
	Description string
	Link string
}

type OurHistoryDataModel struct {
	Meta BasicPageData
	History []Article
}
