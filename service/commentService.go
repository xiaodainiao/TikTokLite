package service

import (
	"TikTokLite/log"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
)

func CommentAction(commentId, videoId, userId int64, comment_text, actionType string) (*message.DouyinCommentActionResponse, error) {

	if actionType == "1" {
		//commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)
		commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)

		if err != nil {
			return nil, err
		}
		user, _ := repository.GetUserInfo(userId)
		commentResponse := &message.DouyinCommentActionResponse{
			Comment: &message.Comment{
				Id:         commentInfo.CommentId,
				User:       messageUserInfo(user),
				Content:    comment_text,
				CreateDate: commentInfo.Time,
			},
		}

		return commentResponse, nil
	} else {
		err := repository.CommentDelete(videoId, commentId)
		if err != nil {
			return nil, err
		}
		commentResponse := &message.DouyinCommentActionResponse{
			Comment: nil,
		}
		return commentResponse, nil
	}

}

//用户评论
func CommentList(videoId int64) (*message.DouyinCommentListResponse, error) {
	comments, err := repository.CommentList(videoId)
	if err != nil {
		return nil, err
	}
	log.Debugf("comments:%v\n", comments)

	list := &message.DouyinCommentListResponse{
		CommentList: make([]*message.Comment, len(comments)),
	}

	for i, comment := range comments {
		//为了找到video_id所对应的user_id，在通过user_id找到user_name.传递给前端
		userID := comment.UserId
		user, _ := repository.GetUserInfo(userID)
		userinfo := messageUserInfo(user)

		v := &message.Comment{
			Id:         comment.CommentId,
			User:       userinfo,
			Content:    comment.Comment,
			CreateDate: comment.Time,
		}
		list.CommentList[i] = v
	}

	return list, nil

}
