package models

import (
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

type MovieInfo struct {
	Id                 int64
	MovieId            int64
	MovieName          string
	MoviePic           string
	MovieDirector      string
	MovieWriter        string
	MovieCountry       string
	MovieLanguage      string
	MovieMainCharacter string
	MovieType          string
	MovieOnTime        string
	MovieSpan          string
	MovieGrade         string
	_CreateTime        string
}

var (
	db orm.Ormer

)


var regMap map[string]*regexp.Regexp= map[string]*regexp.Regexp{
	"name" : regexp.MustCompile(`<span property="v:itemreviewed">(.*?)</span>`),
	"pic" : regexp.MustCompile(`<img src="(.*?)" title="点击看更多海报"`),
	"dct" : regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`),
	"wtr" : regexp.MustCompile(`编剧</span>: <span class='attrs'><a href=".*?">(.*?)</a>`),
	"cnt" : regexp.MustCompile(`制片国家/地区:</span> (.*?)<br/>`),
	"lang" : regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`),
	"mc" : regexp.MustCompile(`主演</span>: <span class='attrs'><a href=".*?" rel="v:starring">(.*?)</a>`),
	"type" : regexp.MustCompile(`<span property="v:genre">(.*?)</span>`),
	"time" : regexp.MustCompile(`<span class="year">\((.*?)\)</span>`),
	"span" : regexp.MustCompile(`<span property="v:runtime" content=".*?">(.*?)</span>`),
	"grd" : regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*?)</strong>`),
}
func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:19950213@tcp(localhost:3306)/test?charset=utf8")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movieInfo *MovieInfo) (int64, error) {
	id, err := db.Insert(movieInfo)
	return id, err
}

//func GetMovieDirector(movieHtml string) string {
//	if movieHtml == "" {
//		return "" + "\n"
//	}
//
//	result := DirectorReg.FindAllStringSubmatch(movieHtml, -1)
//	return result[0][1] + "\n"
//}

func GetMovieDetail(movieHtml, item string) string {
	if movieHtml == "" {
		return "" + "\n"
	}

	result := regMap[item].FindAllStringSubmatch(movieHtml, -1)
	return result[0][1] + "\n"
}

//func GetMovieName(movieHtml string) string {
//	if movieHtml == "" {
//		return "" + "\n"
//	}
//
//	result := NameReg.FindAllStringSubmatch(movieHtml, -1)
//	return result[0][1] + "\n"
//}