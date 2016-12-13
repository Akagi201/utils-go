package radixs_test

import (
	"fmt"
	. "testing"

	"github.com/Akagi201/utilgo/radixs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScan(t *T) {
	keys := map[string]struct{}{}
	conf := &radixs.Config{
		RedisAddr:          "127.0.0.1:6379",
		RedisCluster:       false,
		RedisSentinel:      false,
		RedisSentinels:     []string{},
		RedisSentinelGroup: "master",
	}
	db, err := radixs.InitDB(conf)
	assert.Nil(t, err)
	for i := 0; i < 100; i++ {
		keys[fmt.Sprintf("scantest:%d", i)] = struct{}{}
	}

	for key := range keys {
		require.Nil(t, db.Cmd("SET", key, key).Err)
	}

	output := map[string]struct{}{}
	for r := range db.Scan("scantest:*") {
		output[r] = struct{}{}
	}
	assert.Equal(t, keys, output)

}
