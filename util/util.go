package util

import (
	"math/rand"
	"os"
	"time"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func GetCurrentTimeForString() string {
	currentTime := time.Now()
	return currentTime.Format("200601021504")
}

func GetCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}

//随机生成字符
func RandomString() string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, 16)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Mkdir(path string) error {
	if ok, _ := PathExists(path); ok {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
