package storgeMoudle

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xinghe98/goProxyPool/common/db"
)

type storgeMoudle struct {
	redis *db.Redis
}

func NewStorge() *storgeMoudle {
	redis := db.NewRedisClient()
	return &storgeMoudle{
		redis,
	}
}

func (s *storgeMoudle) SaveIp(value string, score float64) bool {
	err := s.redis.Rdb.ZAdd(s.redis.Ctx, "zset", redis.Z{Score: score, Member: value}).Err()
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
