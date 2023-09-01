package xcache

import (
	"fmt"
	"sync"
	"time"
	"xutils/src/xvalue"
)

type XMemory struct {
	lc	sync.Mutex
	kv	map[string]any
}

func NewXMemory() *XMemory {
	// 创建连接池
	obj := &XMemory{kv: make(map[string]any)}

	// 返回对象
	return obj
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
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 删除key
	for _, key := range keys {
		delete(s.kv, key)
	}
}

func (s *XMemory)HGetAll(key string) map[string]string {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 检查key是否存在
	if _, ok := s.kv[key]; !ok {return nil}

	// 返回key的值
	if obj, ok := s.kv[key].(map[string]string); ok {
		result := make(map[string]string)
		for k, v := range obj {result[k] = v}
		return result
	}

	return nil
}

func (s *XMemory)Exists(keys ...string) (result int64) {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 统计存在的key的数量
	for _, key := range keys {
		if _, ok := s.kv[key]; ok {result++}
	}

	return result
}

func (s *XMemory)Get(key string) string {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 检查key是否存在
	if _, ok := s.kv[key]; !ok {return ""}

	// 返回值
	if v, ok := s.kv[key].(string); ok {return v}
	return ""
}

func (s *XMemory)Set(key string, value string, exp time.Duration) error {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 设置值
	s.kv[key] = value
	return nil
}

func (s *XMemory)HGet(key, field string) string {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 检查key是否存在
	if _, ok := s.kv[key]; !ok {return ""}

	// 返回字段的值
	if obj, ok := s.kv[key].(map[string]string); ok {
		return obj[field]
	}

	return ""
}

func (s *XMemory)HSet(key string, val interface{}) error {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 检查参数类型
	data, ok := val.(map[string]string)
	if !ok {return nil}

	// key不存在就添加
	if _, ok = s.kv[key]; !ok {
		s.kv[key] = data
		return nil
	}

	// key已存在就更新
	if obj, ok := s.kv[key].(map[string]string); ok {
		for _, k := range data {delete(obj, k)}
	}

	return nil
}

func (s *XMemory)HDel(key string, field ...string) error {
	// 加锁
	s.lc.Lock(); defer s.lc.Unlock()

	// 检查key是否存在
	if _, ok := s.kv[key]; !ok {return nil}

	// 删除指定的字段
	if obj, ok := s.kv[key].(map[string]string); ok {
		for _, k := range field {delete(obj, k)}
	}

	return nil
}

func (s *XMemory)HIncr(key, field string, incr int64) error {
	return s.HSet(key, map[string]string{field: fmt.Sprintf("%d", xvalue.S2I64(s.HGet(key, field)) + incr)})
}
