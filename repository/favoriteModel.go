package repository

import (
	"TikTokLite/common"
	"errors"

	"github.com/jinzhu/gorm"
)

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func LikeAction(uid, vid int64) error {
	db := common.GetDB()
	favorite := Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Find(&Favorite{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have liked this video")
	}
	err = db.Create(&favorite).Error
	if err != nil {
		return err
	}
	authorid, _ := CacheGetAuthor(vid)
	// go func() {
	// 	CacheChangeUserCount(uid, add, "like")
	// 	CacheChangeUserCount(authorid, add, "liked")
	// }()
	go CacheChangeUserCount(uid, add, "like")
	go CacheChangeUserCount(authorid, add, "liked")
	return nil
}

func UnLikeAction(uid, vid int64) error {
	db := common.GetDB()
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	authorid, _ := CacheGetAuthor(vid)
	// go func() {
	go CacheChangeUserCount(uid, sub, "like")
	go CacheChangeUserCount(authorid, sub, "liked")
	// }()
	return nil
}

func GetFavoriteList(uid int64) ([]Video, error) {
	var videos []Video
	db := common.GetDB()
	err := db.Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return []Video{}, nil
	} else if err != nil {
		return nil, err
	}
	for i, v := range videos {
		author, err := GetUserInfo(v.AuthorId)
		if err != nil {
			return videos, err
		}
		videos[i].Author = author
	}
	return videos, nil
}
