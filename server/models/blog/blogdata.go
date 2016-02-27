package blog

import (
	"time"
)

type BlogSyncProgress func(message string)

type BlogSync struct {
	BlogSyncId   int `xorm:"autoincr"`
	UserId       int
	AccessToken  string
	GitUrl       string
	SyncType     int
	State        int
	StateMessage string
	CreateTime   time.Time `xorm:"created"`
	ModifyTime   time.Time `xorm:"updated"`
}

type BlogSyncs struct {
	Data  []BlogSync
	Count int
}

type BlogCategory struct {
	Id   int
	Name string
}

type BlogArticle struct {
	Id       int
	Title    string
	Content  string
	Category string
}

type Blog struct {
	Categorys []BlogCategory
	Articles  []BlogArticle
}

type BlogDiff struct {
	AddCategorys []BlogCategory
	DelCategorys []BlogCategory
	AddArticles  []BlogArticle
	ModArticles  []BlogArticle
	DelArticles  []BlogArticle
}
