package service

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"TikTokLite/minioStore"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func PublishVideo(userId int64, saveFile, title string) (*message.DouyinPublishActionResponse, error) {
	client := minioStore.GetMinio()
	videourl, err := client.UploadFile("video", saveFile, strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, err
	}
	imageFile, err := GetImageFile(saveFile)

	if err != nil {
		return nil, err
	}

	log.Debugf("imageFile %v\n", imageFile)

	picurl, err := client.UploadFile("pic", imageFile, strconv.FormatInt(userId, 10))
	if err != nil {
		picurl = "https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7909abe413ec4a1e82032d2beb810157~tplv-k3u1fbpfcp-zoom-in-crop-mark:1304:0:0:0.awebp?"
	}

	err = repository.InsertVideo(userId, videourl, picurl, title)
	if err != nil {
		return nil, err
	}
	return &message.DouyinPublishActionResponse{}, nil
}

func PublishList(tokenUserId, userId int64) (*message.DouyinPublishListResponse, error) {
	videos, err := repository.GetVideoList(userId)
	if err != nil {
		return nil, err
	}
	list := &message.DouyinPublishListResponse{
		VideoList: VideoList(videos, tokenUserId),
	}

	return list, nil
}

func GetImageFile(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "/")
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-3]) + "jpg"
	picpath := config.GetConfig().Path.Picfile
	picName := filepath.Join(picpath, videoName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picName)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	log.Debugf(picName)
	return picName, nil
}
