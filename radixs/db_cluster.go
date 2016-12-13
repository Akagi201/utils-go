package radixs

import (
	"log"

	"github.com/mediocregopher/radix.v2/cluster"
	"github.com/mediocregopher/radix.v2/redis"
)

type clusterDB struct {
	*cluster.Cluster
}

func newClusterDB() (DBer, error) {
	log.Printf("connecting to redis cluster at %s", dbConf.RedisAddr)
	c, err := cluster.New(dbConf.RedisAddr)
	if err != nil {
		log.Fatalln(err)
	}
	return &clusterDB{c}, err
}

func (d *clusterDB) Cmd(cmd string, args ...interface{}) *redis.Resp {
	return d.Cluster.Cmd(cmd, args...)
}

func (d *clusterDB) Pipe(p ...*PipePart) ([]*redis.Resp, error) {
	// We can't really pipe with cluster, just do Cmd in a loop
	ret := make([]*redis.Resp, 0, len(p))
	for i := range p {
		r := d.Cmd(p[i].cmd, p[i].args...)
		if r.Err != nil {
			return nil, r.Err
		}
		ret = append(ret, r)
	}
	return ret, nil
}

func (d *clusterDB) Scan(pattern string) <-chan string {
	retCh := make(chan string)
	go func() {
		defer close(retCh)

		redisClients, err := d.GetEvery()
		if err != nil {
			log.Printf("clusterScan(%s) ClientPerMaster(): %s", pattern, err)
			return
		}

		// Make sure we return all clients no matter what
		for i := range redisClients {
			defer d.Put(redisClients[i])
		}

		for _, redisClient := range redisClients {
			if err := scanHelper(redisClient, pattern, retCh); err != nil {
				log.Printf("clusterScan(%s) scanHelper(): %s", pattern, err)
				return
			}
		}
	}()
	return retCh
}

func (d *clusterDB) GetAddr() (string, error) {
	return d.GetAddrForKey(""), nil
}
