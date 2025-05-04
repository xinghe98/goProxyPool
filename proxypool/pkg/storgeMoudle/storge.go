package storgeMoudle

import (
	"fmt"
	"log"

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
	err := s.redis.Rdb.ZAdd(s.redis.Ctx, s.setname, redis.Z{Score: score, Member: value}).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/**
 *
 * @param member
 * 当分数降为0时，删除该元素
 */
func (s *storgeMoudle) DeleteIp(ip string) {
}

// api获取代理
func (s *storgeMoudle) GetIp() string {
	return ""
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
