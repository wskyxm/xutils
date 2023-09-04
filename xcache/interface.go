package xcache

import (
	"time"
)

type ICache interface {
	Expire(key string, exp time.Duration)
	SSet(key string, val ...interface{}) error
	SDel(key string, val ...interface{}) error
	SCard(key string) int64
	SIsMember(key string, val interface{}) bool
	ZSet(key string, members ...*Z) error
	ZDel(key string, member ...interface{}) error
	ZCard(key string) int64
	ZRange(key string, start, stop int64) []string
	ZScore(key, member string) float64
	ZRangeByScoreLimit(key string, min, max string, offset, count int64) []string
	ZRangeByScore(key string, min, max string) []string
	ZRangeByScoreWithScores(key string, min, max string) []Z
	ZRevRangeByScoreLimitWithScores(key string, min, max string, offset, count int64) []Z
	ZRevRangeByScoreWithScores(key string, min, max string) []Z
	ZRevRangeByScoreLimit(key string, min, max string, offset, count int64) []string
	ZRevRangeByScore(key string, min, max string) []string
	ZIncrBy(key, member string, increment float64) float64
	Del(keys ...string)
	HGetAll(key string) map[string]string
	Exists(keys ...string) int64
	Get(key string) string
	Set(key string, value string, exp time.Duration) error
	HGet(key, field string) string
	HSet(key string, val interface{}) error
	HDel(key string, field ...string) error
	HIncr(key, field string, incr int64) error
}