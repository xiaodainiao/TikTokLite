package common

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
)

var (
	redisClient *redis.Pool
	rs          *redsync.Redsync
	//锁过期时间
	lockExpiry = 2 * time.Second
	//获取锁失败重试时间间隔
	retryDelay = 500 * time.Millisecond
	//值过期时间
	valueExpire  = 86400
	ErrMissCache = errors.New("miss Cache")
	//锁设置
	option = []redsync.Option{
		redsync.WithExpiry(lockExpiry),
		redsync.WithRetryDelay(retryDelay),
	}
)

func RedisInit() {
	conf := config.GetConfig()
	network := conf.Redis.Network
	address := conf.Redis.Host
	port := conf.Redis.Port
	auth := conf.Redis.Auth
	host := fmt.Sprintf("%s:%s", address, port)
	redisClient = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, host,
				redis.DialPassword(auth),
				redis.DialDatabase(2),
			)
			if err != nil {
				log.Error("conn redis failed,", err)
				return nil, err
			}

			return c, err
		},
	}
	redisClient.Get().Do("flushdb")
	sync := redigo.NewPool(redisClient)
	rs = redsync.New(sync)
	log.Info("redis conn success")

}

func GetRedis() redis.Conn {
	return redisClient.Get()
}
func CloseRedis() {
	redisClient.Close()
}

func Exists(key string) bool {
	conn := redisClient.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func GetLock(key string) (*redsync.Mutex, error) {
	mutex := rs.NewMutex(key+"_lock", option...)
	if err := mutex.Lock(); err != nil {
		return mutex, err
	}
	return mutex, nil
}

func UnLock(mutex *redsync.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		return err
	}
	return nil
}

/////////////////////////String类型接口////////////////////////////////////////

func CacheSet(key string, data interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX", valueExpire)
	if err != nil {
		return err
	}
	return nil
}

func CacheGet(key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	if len(reply) == 0 {
		return nil, ErrMissCache
	}

	return reply, nil
}

///////////////////////////List类型接口////////////////////////////////////////

func CacheLPush(key string, value ...interface{}) error {
	return listPush("LPUSH", key, value)
}

func CacheRPush(key string, value ...interface{}) error {
	return listPush("RPUSH", key, value)
}

func CacheLPop(key string) ([]byte, error) {
	return listPop("LPOP", key)
}

func CacheRPop(key string) ([]byte, error) {
	return listPop("RPOP", key)
}

func CacheLGetAll(key string) ([][]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.ByteSlices(conn.Do("LRANGE", key, "0", "-1"))
	if err != nil {
		return [][]byte{}, err
	}
	return data, nil
}

func listPush(op, key string, data ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	for _, d := range data {
		value, err := json.Marshal(d)
		if err != nil {
			return err
		}
		_, err = conn.Do(op, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func listPop(op, key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do(op, key))
	if err != nil {
		return reply, err
	}

	return reply, nil
}

/////////////////////////Hash类型接口///////////////////////////////////////////

func CacheHSet(key, mkey string, value ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()

	for _, d := range value {
		data, err := json.Marshal(d)
		if err != nil {
			return nil
		}

		_, err = conn.Do("HSET", key, mkey, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func CacheHGet(key, mkey string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", key, mkey))

	//fmt.Printf("data:%v", data)
	if err != nil {
		return []byte{}, err
	}
	if len(data) == 0 {
		return []byte{}, ErrMissCache
	}
	return data, nil
}

func CacheDelHash(key, mkey string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey)
	if err != nil {
		return err
	}
	return nil
}

func CacheDelHash2(key, mkey, comment_id string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey, comment_id)
	if err != nil {
		return err
	}
	return nil
}
