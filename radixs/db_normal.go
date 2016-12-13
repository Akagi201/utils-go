package radixs

import (
	"log"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

type normalDB struct {
	*pool.Pool
}

func newNormalDB() (DBer, error) {
	log.Printf("connecting to redis at %s", Conf.RedisAddr)
	p, err := pool.New("tcp", Conf.RedisAddr, 200)
	if err != nil {
		log.Fatal(err)
	}
	return &normalDB{p}, err
}

func (d *normalDB) Cmd(cmd string, args ...interface{}) *redis.Resp {
	c, err := d.Get()
	if err != nil {
		return redis.NewResp(err)
	}

	r := c.Cmd(cmd, args...)
	d.Put(c)
	return r
}

func (d normalDB) Pipe(p ...*PipePart) ([]*redis.Resp, error) {
	c, err := d.Get()
	if err != nil {
		return nil, err
	}
	defer d.Put(c)

	for i := range p {
		c.PipeAppend(p[i].cmd, p[i].args...)
	}

	rs := make([]*redis.Resp, len(p))
	for i := range rs {
		rs[i] = c.PipeResp()
		if err = rs[i].Err; err != nil {
			return nil, err
		}
	}

	return rs, nil
}

func (d normalDB) Scan(pattern string) <-chan string {
	retCh := make(chan string)
	go func() {
		defer close(retCh)

		redisClient, err := d.Get()
		if err != nil {
			log.Printf("normalScan(%s) Get(): %s", pattern, err)
			return
		}
		defer d.Put(redisClient)

		if err = scanHelper(redisClient, pattern, retCh); err != nil {
			log.Printf("normalScan(%s) scanHelper: %s", pattern, err)
			return
		}
	}()

	return retCh
}

func (d normalDB) GetAddr() (string, error) {
	return Conf.RedisAddr, nil
}
