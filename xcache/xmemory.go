package xcache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/wskyxm/xutils/xvalue"
	"time"
)

type XMemory struct {
	mc *cache.Cache
}

func NewXMemory() *XMemory {
	obj := &XMemory{mc: cache.New(cache.NoExpiration, time.Minute * 5)}
	return obj
}

func (s *XMemory)hget(key string) map[string]string {
	// 初始化返回值
	result := make(map[string]string)

	// 获取key值
	val, ok := s.mc.Get(key)
	if !ok {return result}

	// 转换key值类型
	temp, ok := val.(map[string]string)
	if ok {return temp} else {return result}
}

func (s *XMemory)get(key string) string {
	// 获取key值
	val, ok := s.mc.Get(key)
	if !ok {return ""}

	// 转换key值类型
	result, ok := val.(string)
	if ok {return result} else {return ""}
}

func (s *XMemory)IsHash(key string) bool {
	// 获取key值
	val, ok := s.mc.Get(key)
	if !ok {return false}

	// 返回类型
	switch val.(type) {
	case map[string]string: return true
	default: return false
	}
}

func (s *XMemory)Expire(key string, exp time.Duration) {
}

func (s *XMemory)SSet(key string, val ...interface{}) error {
	panic("not supported")
	return nil
}

func (s *XMemory)SDel(key string, val ...interface{}) error {
	panic("not supported")
	return nil
}

func (s *XMemory)SCard(key string) int64 {
	panic("not supported")
	return 0
}

func (s *XMemory)SIsMember(key string, val interface{}) bool {
	panic("not supported")
	return false
}

func (s *XMemory)ZSet(key string, members ...*Z) error {
	panic("not supported")
	return nil
}

func (s *XMemory)ZDel(key string, member ...interface{}) error {
	panic("not supported")
	return nil
}

func (s *XMemory)ZCard(key string) int64 {
	panic("not supported")
	return 0
}

func (s *XMemory)ZRange(key string, start, stop int64) []string {
	panic("not supported")
	return nil
}

func (s *XMemory)ZScore(key, member string) float64 {
	panic("not supported")
	return 0
}

func (s *XMemory)ZRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRangeByScore(key string, min, max string) []string {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRangeByScoreWithScores(key string, min, max string) []Z {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRevRangeByScoreLimitWithScores(key string, min, max string, offset, count int64) []Z {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRevRangeByScoreWithScores(key string, min, max string) []Z {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRevRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	panic("not supported")
	return nil
}

func (s *XMemory)ZRevRangeByScore(key string, min, max string) []string {
	panic("not supported")
	return nil
}

func (s *XMemory)ZIncrBy(key, member string, increment float64) float64 {
	panic("not supported")
	return 0
}

func (s *XMemory)Del(keys ...string) {
	for _, key := range keys {s.mc.Delete(key)}
}

func (s *XMemory)HGetAll(key string) map[string]string {
	return s.hget(key)
}

func (s *XMemory)Exists(keys ...string) (result int64) {
	// 统计存在的key的数量
	for _, key := range keys {
		if _, ok := s.mc.Get(key); ok {result++}
	}

	return result
}

func (s *XMemory)Get(key string) string {
	return s.get(key)
}

func (s *XMemory)Set(key string, value string, exp time.Duration) error {
	s.mc.Set(key, value, exp)
	return nil
}

func (s *XMemory)HGet(key, field string) string {
	return s.HGetAll(key)[field]
}

func (s *XMemory)HSet(key string, val interface{}) error {
	// 检查参数类型
	newval, ok := val.(map[string]string)
	if !ok {return nil}

	// 检查是不是hash类型
	if !s.IsHash(key) {return nil}

	// 获取旧值
	oldval := s.HGetAll(key)

	// 增加新值
	for k, v := range newval {oldval[k] = v}

	// 更新key值
	s.mc.Set(key, oldval, cache.NoExpiration)
	return nil
}

func (s *XMemory)HDel(key string, field ...string) error {
	// 检查是不是hash类型
	if !s.IsHash(key) {return nil}

	// 删除指定的字段
	values := s.HGetAll(key)
	for _, k := range field {delete(values, k)}

	// 更新删除后的字段
	s.mc.Set(key, values, cache.NoExpiration)
	return nil
}

func (s *XMemory)HIncr(key, field string, incr int64) error {
	return s.HSet(key, map[string]string{field: fmt.Sprintf("%d", xvalue.S2I64(s.HGet(key, field)) + incr)})
}
