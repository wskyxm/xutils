package xcache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"runtime"
	"time"
)

type XRedis struct {
	cli	 *redis.Client
	ctx	 context.Context
}

func NewXRedis(addr, pass string) *XRedis {
	// 创建连接池
	obj := &XRedis{}
	obj.ctx = context.Background()
	obj.cli = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: pass,
		DB: 0,
		PoolSize: 8 * runtime.NumCPU(),
		MinIdleConns: 4 * runtime.NumCPU(),
	})

	// 检查redis通讯是否正常
	_, err := obj.cli.Ping(obj.ctx).Result()
	if err != nil {panic(err)}

	// 返回对象
	return obj
}

func (s *XRedis)Expire(key string, exp time.Duration) {
	if exp > 0 {s.cli.Expire(s.ctx, key, exp)
	} else {s.cli.Persist(s.ctx, key)}
}

func (s *XRedis)SSet(key string, val ...interface{}) error {
	if len(val) == 0 {return nil}
	_, err := s.cli.SAdd(s.ctx, key, val...).Result()
	return checkerr(err)
}

func (s *XRedis)SDel(key string, val ...interface{}) error {
	if len(val) == 0 {return nil}
	_, err := s.cli.SRem(s.ctx, key, val...).Result()
	return checkerr(err)
}

func (s *XRedis)SCard(key string) int64 {
	result, err := s.cli.SCard(s.ctx, key).Result()
	checkerr(err); return result
}

func (s *XRedis)SIsMember(key string, val interface{}) bool {
	result, err := s.cli.SIsMember(s.ctx, key, val).Result()
	checkerr(err); return result
}

func (s *XRedis)ZSet(key string, members ...*Z) error {
	if len(members) == 0 {return nil}
	return checkerr(s.cli.ZAdd(s.ctx, key, z2rediszptr(members)...).Err())
}

func (s *XRedis)ZDel(key string, member ...interface{}) error {
	if len(member) == 0 {return nil}
	return checkerr(s.cli.ZRem(s.ctx, key, member...).Err())
}

func (s *XRedis)ZCard(key string) int64 {
	result, err := s.cli.ZCard(s.ctx, key).Result()
	checkerr(err); return result
}

func (s *XRedis)ZRange(key string, start, stop int64) []string {
	result, err := s.cli.ZRange(s.ctx, key, start, stop).Result()
	checkerr(err); return result
}

func (s *XRedis)ZScore(key, member string) float64 {
	result, err := s.cli.ZScore(s.ctx, key, member).Result()
	checkerr(err); return result
}

func (s *XRedis)ZRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	opt := &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}
	result, err := s.cli.ZRangeByScore(s.ctx, key, opt).Result()
	checkerr(err); return result
}

func (s *XRedis)ZRangeByScore(key string, min, max string) []string {
	result, err := s.cli.ZRangeByScore(s.ctx, key, &redis.ZRangeBy{Min: min, Max: max}).Result()
	checkerr(err); return result
}

func (s *XRedis)ZRangeByScoreWithScores(key string, min, max string) []Z {
	result, err := s.cli.ZRangeByScoreWithScores(s.ctx, key, &redis.ZRangeBy{Min: min, Max: max}).Result()
	checkerr(err); return redisz2z(result)
}

func (s *XRedis)ZRevRangeByScoreLimitWithScores(key string, min, max string, offset, count int64) []Z {
	opt := &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}
	result, err := s.cli.ZRevRangeByScoreWithScores(s.ctx, key, opt).Result()
	checkerr(err); return redisz2z(result)
}

func (s *XRedis)ZRevRangeByScoreWithScores(key string, min, max string) []Z {
	result, err := s.cli.ZRevRangeByScoreWithScores(s.ctx, key, &redis.ZRangeBy{Min: min, Max: max}).Result()
	checkerr(err); return redisz2z(result)
}

func (s *XRedis)ZRevRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	opt := &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}
	result, err := s.cli.ZRevRangeByScore(s.ctx, key, opt).Result()
	checkerr(err); return result
}

func (s *XRedis)ZRevRangeByScore(key string, min, max string) []string {
	result, err := s.cli.ZRevRangeByScore(s.ctx, key, &redis.ZRangeBy{Min: min, Max: max}).Result()
	checkerr(err); return result
}

func (s *XRedis)ZIncrBy(key, member string, increment float64) float64 {
	result, err := s.cli.ZIncrBy(s.ctx, key, increment, member).Result()
	checkerr(err); return result
}

func (s *XRedis)Del(keys ...string) {
	if len(keys) == 0 {return}
	checkerr(s.cli.Del(s.ctx, keys...).Err())
}

func (s *XRedis)HGetAll(key string) map[string]string {
	result, err := s.cli.HGetAll(s.ctx, key).Result()
	checkerr(err); return result
}

func (s *XRedis)Exists(keys ...string) int64 {
	if len(keys) == 0 {return 0}
	result, err := s.cli.Exists(s.ctx, keys...).Result()
	checkerr(err); return result
}

func (s *XRedis)Get(key string) string {
	result, err := s.cli.Get(s.ctx, key).Result()
	checkerr(err); return result
}

func (s *XRedis)Set(key string, value string, exp time.Duration) error {
	return checkerr(s.cli.Set(s.ctx, key, value, exp).Err())
}

func (s *XRedis)HGet(key, field string) string {
	result, err := s.cli.HGet(s.ctx, key, field).Result()
	checkerr(err); return result
}

func (s *XRedis)HSet(key string, val interface{}) error {
	return checkerr(s.cli.HSet(s.ctx, key, val).Err())
}

func (s *XRedis)HDel(key string, field ...string) error {
	if len(field) == 0 {return nil}
	return checkerr(s.cli.HDel(s.ctx, key, field...).Err())
}

func (s *XRedis)HIncr(key, field string, incr int64) error {
	return checkerr(s.cli.HIncrBy(s.ctx, key, field, incr).Err())
}
