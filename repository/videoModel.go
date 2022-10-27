package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"TikTokLite/util"
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Title         string `gorm:"column:title;"`
	Author        User   `gorm:"foreignkey:AuthorId"`
}

func (Video) TableName() string {
	return "videos"
}

func InsertVideo(authorid int64, playurl, coverurl, title string) error {
	video := Video{
		AuthorId:      authorid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   util.GetCurrentTime(),
		Title:         title,
	}
	db := common.GetDB()
	err := db.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func GetVideoList(AuthorId int64) ([]Video, error) {
	var videos []Video
	author, err := GetUserInfo(AuthorId)
	if err != nil {
		return videos, err
	}
	db := common.GetDB()
	err = db.Where("author_id = ?", AuthorId).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	for i := range videos {
		videos[i].Author = author
	}
	return videos, nil
}

func GetVideoListByFeed(currentTime int64) ([]Video, error) {
	var videos []Video
	db := common.GetDB()
	err := db.Where("publish_time < ?", currentTime).Limit(20).Order("video_id DESC").Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	for i, v := range videos {
		author, err := GetUserInfo(v.AuthorId)
		CacheSetAuthor(v.Id, v.AuthorId)
		if err != nil {
			return videos, err
		}
		videos[i].Author = author
	}
	return videos, nil
}

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := common.CacheHSet("video", key, authorid)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := common.CacheHGet("video", key)
	if err != nil {
		return 0, err
	}
	uid := int64(0)
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
