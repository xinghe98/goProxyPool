package storgeMoudle

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/xinghe98/goProxyPool/common/db"
)

type storgeMoudle struct {
	redis   *db.Redis
	setname string
}

func NewStorge() *storgeMoudle {
	redis := db.NewRedisClient()
	setname := viper.GetString("DB.set")
	return &storgeMoudle{
		redis,
		setname,
	}
}

func (s *storgeMoudle) SaveIp(value string, score float64) bool {
	if score < 0 { // 传入的score<0 说明要减分了，顺便删除为0的
		// 先获取之前的score
		originScore, _ := s.redis.Rdb.ZScore(s.redis.Ctx, s.setname, value).Result()
		if originScore <= 0 {
			// 当分数已经为0时,则直接删除
			_, err := s.redis.Rdb.ZRem(s.redis.Ctx, s.setname, value).Result()
			if err != nil {
				log.Println("[❌] 数据库错误")
				return false
			}
		} else {
			// 做减分操作
			err := s.redis.Rdb.ZAdd(s.redis.Ctx, s.setname, redis.Z{Score: originScore + score, Member: value}).Err()
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
		return true
	} else { // 如何传入的score>0说明是加分:第一次存或者是检测后给100
		err := s.redis.Rdb.ZAdd(s.redis.Ctx, s.setname, redis.Z{Score: score, Member: value}).Err()
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
}

// api获取代理
func (s *storgeMoudle) GetIp() (string, error) {
	ips, err := s.redis.Rdb.ZRevRangeByScore(s.redis.Ctx, s.setname, &redis.ZRangeBy{Min: "100", Max: "100"}).Result()
	if err != nil {
		return "", err
	}
	if len(ips) == 0 {
		return "", err
	}
	ip := ips[rand.Intn(len(ips))]
	return ip, nil
}

// 获取任意数量的ip
func (s *storgeMoudle) GetSomeIp(start int64, end int64) []string {
	res, err := s.redis.Rdb.ZRange(s.redis.Ctx, s.setname, start, end).Result()
	if err != nil {
		log.Println(err)
		return []string{"获取错误,请检查数据库"}
	}
	return res
}

// 获取总数
func (s *storgeMoudle) GetCount() int {
	count, _ := s.redis.Rdb.ZCard(s.redis.Ctx, s.setname).Result()
	return int(count)
}
