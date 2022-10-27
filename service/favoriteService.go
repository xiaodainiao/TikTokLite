package service

import (
	"TikTokLite/log"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
)

func FavoriteAction(uid, vid int64, action int8) error {
	if action == 1 {
		log.Infof("like action uid:%v,vid:%v", uid, vid)
		err := repository.LikeAction(uid, vid)
		if err != nil {
			return err
		}
	} else {
		log.Infof("unlike action uid:%v,vid:%v", uid, vid)
		err := repository.UnLikeAction(uid, vid)
		if err != nil {
			return err
		}
	}
	return nil
}

func FavoriteList(tokenUid, uid int64) (*message.DouyinFavoriteListResponse, error) {
	favList, err := repository.GetFavoriteList(uid)
	if err != nil {
		return nil, err
	}
	// log.Infof("user:%v, followList:%+v", uid, favList)

	favListResponse := message.DouyinFavoriteListResponse{
		VideoList: VideoList(favList, tokenUid),
	}

	return &favListResponse, nil
}

func tokenFavList(tokenUserId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	list, err := repository.GetFavoriteList(tokenUserId)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		m[v.Id] = struct{}{}
	}
	return m, nil
}
