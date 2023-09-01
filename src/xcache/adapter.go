package xcache

import (
	"github.com/go-redis/redis/v8"
	"xutils/src/xlog"
)

type Z struct {
	S float64
	M interface{}
}

func checkerr(err error) error {
	if err != nil && err.Error() == "redis: nil" {return nil}
	if err != nil {xlog.ErrorDepth("redis error(%s)", 2, err.Error())}
	return err
}

func z2redisz(z []Z) []redis.Z {
	result := make([]redis.Z, len(z))
	for i := 0; i < len(z); i++ {result[i] = redis.Z{Member: z[i].M, Score: z[i].S}}
	return result
}

func redisz2z(z []redis.Z) []Z {
	result := make([]Z, len(z))
	for i := 0; i < len(z); i++ {result[i] = Z{M: z[i].Member, S: z[i].Score}}
	return result
}

func z2rediszptr(z []*Z) []*redis.Z {
	result := make([]*redis.Z, len(z))
	for i := 0; i < len(z); i++ {result[i] = &redis.Z{Member: z[i].M, Score: z[i].S}}
	return result
}

func redisz2zptr(z []*redis.Z) []*Z {
	result := make([]*Z, len(z))
	for i := 0; i < len(z); i++ {result[i] = &Z{M: z[i].Member, S: z[i].Score}}
	return result
}
