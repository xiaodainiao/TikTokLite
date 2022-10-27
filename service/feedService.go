package service

import (
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"TikTokLite/util"
)

func GetFeedList(currentTime, tokenUserId int64) (*message.DouyinFeedResponse, error) {
	videoList, err := repository.GetVideoListByFeed(currentTime)
	if err != nil {
		return nil, err
	}
	feed := &message.DouyinFeedResponse{
		VideoList: VideoList(videoList, tokenUserId),
	}

	nextTime := util.GetCurrentTime()
	if len(videoList) == 20 {
		nextTime = videoList[len(videoList)-1].PublishTime
	}
	feed.NextTime = nextTime
	return feed, nil
}

func VideoList(videoList []repository.Video, userId int64) []*message.Video {
	var err error
	FollowList := make(map[int64]struct{})
	favList := make(map[int64]struct{})
	if userId != int64(0) {
		FollowList, err = tokenFollowList(userId)
		if err != nil {
			return nil
		}
		favList, err = tokenFavList(userId)
		if err != nil {
			return nil
		}
	}
	lists := make([]*message.Video, len(videoList))
	for i, video := range videoList {
		v := &message.Video{
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Author:        messageUserInfo(video.Author),
			Title:         video.Title,
		}
		if _, ok := FollowList[video.AuthorId]; ok {
			v.Author.IsFollow = true
		}
		if _, ok := favList[video.Id]; ok {
			v.IsFavorite = true
		}
		lists[i] = v
	}
	return lists
}
