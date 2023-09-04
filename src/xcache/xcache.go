package xcache

import (
	"gorm.io/gorm/clause"
	"time"
	"xutils/src/xdao"
	"xutils/src/xerr"
	"xutils/src/xvalue"
)

// 缓存对象
var cacheimp ICache

// 初始化缓存，如果addr为空，则使用内存缓存
func Initialize(addr string, pass string) {
	if addr != "" {cacheimp = NewXRedis(addr, pass)
	} else {cacheimp = NewXMemory()}
}

func Expire(key string, exp time.Duration) {
	cacheimp.Expire(key, exp)
}

func SSet(key string, val ...interface{}) error {
	return cacheimp.SSet(key, val...)
}

func SDel(key string, val ...interface{}) error {
	return cacheimp.SDel(key, val...)
}

func SCard(key string) int64 {
	return cacheimp.SCard(key)
}

func SIsMember(key string, val interface{}) bool {
	return cacheimp.SIsMember(key, val)
}

func ZSet(key string, members ...*Z) error {
	return cacheimp.ZSet(key, members...)
}

func ZDel(key string, member ...interface{}) error {
	return cacheimp.ZDel(key, member...)
}

func ZCard(key string) int64 {
	return cacheimp.ZCard(key)
}

func ZRange(key string, start, stop int64) []string {
	return cacheimp.ZRange(key, start, stop)
}

func ZScore(key, member string) float64 {
	return cacheimp.ZScore(key, member)
}

func ZRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	return cacheimp.ZRangeByScoreLimit(key, min, max, offset, count)
}

func ZRangeByScore(key string, min, max string) []string {
	return cacheimp.ZRangeByScore(key, min, max)
}

func ZRangeByScoreWithScores(key string, min, max string) []Z {
	return cacheimp.ZRangeByScoreWithScores(key, min, max)
}

func ZRevRangeByScoreLimitWithScores(key string, min, max string, offset, count int64) []Z {
	return cacheimp.ZRevRangeByScoreLimitWithScores(key, min, max, offset, count)
}

func ZRevRangeByScoreWithScores(key string, min, max string) []Z {
	return cacheimp.ZRevRangeByScoreWithScores(key, min, max)
}

func ZRevRangeByScoreLimit(key string, min, max string, offset, count int64) []string {
	return cacheimp.ZRevRangeByScoreLimit(key, min, max, offset, count)
}

func ZRevRangeByScore(key string, min, max string) []string {
	return cacheimp.ZRevRangeByScore(key, min, max)
}

func ZIncrBy(key, member string, increment float64) float64 {
	return cacheimp.ZIncrBy(key, member, increment)
}

func Del(keys ...string) {
	cacheimp.Del(keys...)
}

func HGetAll(key string) map[string]string {
	return cacheimp.HGetAll(key)
}

func Exists(keys ...string) int64 {
	return cacheimp.Exists(keys...)
}

func Get(key string) string {
	return cacheimp.Get(key)
}

func Set(key string, value string, exp time.Duration) error {
	return cacheimp.Set(key, value, exp)
}

func HGet(key, field string) string {
	return cacheimp.HGet(key, field)
}

func HSet(key string, val interface{}) error {
	return cacheimp.HSet(key, val)
}

func HDel(key string, field ...string) error {
	return cacheimp.HDel(key, field...)
}

func HIncr(key, field string, incr int64) error {
	return cacheimp.HIncr(key, field, incr)
}

func DSetAlways(key, pkey string, val *xvalue.Value, field string, item interface{}) error {
	// 参数检查
	if val == nil {return xerr.InvalidParameter}

	// 更新缓存
	HSet(key, map[string]string{field: val.String()})
	if item == nil {return nil}

	// 更新数据库
	return xdao.Model(item).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: pkey}},
		DoUpdates: clause.Assignments(map[string]any{field: val.Any()}),
	}).Create(item).Error
}

func DSet(key string, val *xvalue.Value, field string, model interface{}, query interface{}, args ...interface{}) error {
	// 参数检查
	if val == nil {return xerr.InvalidParameter}

	// 更新缓存
	HSet(key, map[string]string{field: val.String()})
	if model == nil {return nil}

	// 更新数据库
	return xdao.Model(model).Where(query, args...).Update(field, val.Any()).Error
}

func DGet(key, field string, value interface{}, item interface{}, query interface{}, args ...interface{}) *xvalue.Value {
	// 返回值
	result := &xvalue.Value{}

	// 查询缓存
	if result.SetValue(HGet(key, field)); result.String() != "" {return result}
	if item == nil {return result}

	// 数据库查询
	err := xdao.Where(query, args...).First(item).Error
	if err != nil {return result} else {result.SetValue(value)}

	// 更新缓存
	HSet(key, map[string]string{field: result.String()})

	// 返回结果
	return result
}
