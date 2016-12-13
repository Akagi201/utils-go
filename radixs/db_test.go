package radixs_test

import (
	"fmt"
	. "testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/Akagi201/utilgo/radixs"
)

func TestScan(t *T) {
	keys := map[string]struct{}{}
	conf := &radixs.ConfigMeta{
		RedisAddr:          "127.0.0.1:6379",
		RedisCluster:       false,
		RedisSentinel:      false,
		RedisSentinels:     []string{},
		RedisSentinelGroup: "master",
	}
	radixs.InitDB(conf)
	for i := 0; i < 100; i++ {
		keys[fmt.Sprintf("scantest:%d", i)] = struct{}{}
	}

	for key := range keys {
		require.Nil(t, radixs.Inst.Cmd("SET", key, key).Err)
	}

	output := map[string]struct{}{}
	for r := range radixs.Inst.Scan("scantest:*") {
		output[r] = struct{}{}
	}
	assert.Equal(t, keys, output)

}
