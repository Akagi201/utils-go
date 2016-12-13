package radixs

import (
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

// DBer is implemented and used by the rest of okq to interact with whatever
// backend has been chosen
type DBer interface {

	// Cmd is a function which will perform the given cmd/args in redis and
	// returns the resp. It automatically handles using redis cluster, if that
	// is enabled
	Cmd(string, ...interface{}) *redis.Resp

	// Pipe runs a set of commands (given by p) one after the other. It is *not*
	// guaranteed that all the commands will be run on the same client. If any
	// commands return an error the pipeline will stop and return that error.
	// Otherwise the Resp from each command is returned in a slice
	//
	//	r, err := db.Pipe(
	//		db.PP("SET", "foo", "bar"),
	//		db.PP("GET", "foo"),
	//	)
	Pipe(...*PipePart) ([]*redis.Resp, error)

	// Scan is a function which returns a channel to which keys matching the
	// given pattern are written to. The channel must be read from until it is
	// closed, which occurs when there are no more keys or when an error has
	// occurred (this error will be logged)
	//
	// This should not be used in any critical paths
	Scan(string) <-chan string

	// GetAddr returns any valid address of a redis instance. Useful for cases
	// where we want to create redis connections external to this db package
	GetAddr() (string, error)
}

// PipePart is a single command to be run in a pipe. See Pipe for an example on
// usage
type PipePart struct {
	cmd  string
	args []interface{}
}

// PP should be called NewPipePart, but that's pretty verbose. It simple returns
// a new PipePart, to be used in a call to Pipe
func PP(cmd string, args ...interface{}) *PipePart {
	return &PipePart{cmd, args}
}

// ConfigMeta contains DB configs
type ConfigMeta struct {
	// RedisAddr Address redis is listening on
	RedisAddr string `json:"redis_addr"`
	// RedisSentinels A sentinel address to connect to through to the client - overrides other options
	RedisSentinels []string `json:"redis_sentinel_addr"`
	// RedisSentinelGroup A redis sentinel group name for selecting which redis masters to connect
	RedisSentinelGroup string `json:"redis_sentinel_group"`
	// RedisCluster Whether or not to treat the redis address as a node in a larger cluster
	RedisCluster bool `json:"redis_cluster"`
	// RedisSentinel Whether or not to redis sentinel mode
	RedisSentinel bool `json:"redis_sentinel"`
}

var (
	// Conf global DB configs
	Conf *ConfigMeta
	// Inst is an instance of DBer which is automatically initialized and which is
	// what should be used by the rest of okq
	Inst DBer
)

// InitDB init DB with configs, if config is nil, then default configs will be used.
func InitDB(config *ConfigMeta) {
	var err error

	if config == nil {
		Conf = &ConfigMeta{
			RedisAddr:          "127.0.0.1:6379",
			RedisCluster:       false,
			RedisSentinel:      false,
			RedisSentinels:     []string{},
			RedisSentinelGroup: "master",
		}
	} else {
		Conf = config
	}

	if Conf.RedisSentinel {
		Inst, err = newSentinelDB()
	} else if Conf.RedisCluster {
		Inst, err = newClusterDB()
	} else {
		Inst, err = newNormalDB()
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func scanHelper(redisClient *redis.Client, pattern string, retCh chan string) error {
	cursor := "0"
	for {
		r := redisClient.Cmd("SCAN", cursor, "MATCH", pattern)
		if r.Err != nil {
			return r.Err
		}
		elems, err := r.Array()
		if err != nil {
			return err
		}

		results, err := elems[1].List()
		if err != nil {
			return err
		}

		for i := range results {
			retCh <- results[i]
		}

		if cursor, err = elems[0].Str(); err != nil {
			return err
		} else if cursor == "0" {
			return nil
		}
	}
}
