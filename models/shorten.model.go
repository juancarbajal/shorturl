package models
import (
	"github.com/juancarbajal/shorturl/configs"
	"github.com/juancarbajal/shorturl/entities"
)

type ShortenModel struct{
	
}

func (shortenModel ShortenModel) SaveUrl(url string, shortenURL string)(bool){
	db := configs.DbConnect()
	u := entities.Shorten{ID: shortenURL, Url: url}
	result := db.Create(&u)
	return result.RowsAffected > 0
}

func (shortenModel ShortenModel) GetUrl(shortenURL string) string{
	db := configs.DbConnect()
	var s entities.Shorten
	_= db.First(&s, "ID=?", shortenURL)
	return s.Url
}
