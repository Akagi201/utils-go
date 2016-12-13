package radixs

import (
	"errors"
	"strings"
	"time"

	"github.com/grooveshark/golib/agg"
	"github.com/mediocregopher/radix.v2/redis"
)

// DefaultTimeout is used when reading from socket
const DefaultTimeout = 30 * time.Second

// Debug If true turns on debug logging and agg support (see
// https://github.com/grooveshark/golib)
var Debug bool

// Client is a client for the redis. It can talk to a pool of redis
// instances and failover from one to the other if one loses connectivity
type Client struct {
	clients map[string]*redis.Client

	// Timeout to use for reads/writes to redis. This defaults to DefaultTimeout,
	// but can be overwritten immediately after NewClient is called
	Timeout time.Duration
}

// New return a *Client for Cmd
func New(addr ...string) *Client {
	c := Client{
		clients: map[string]*redis.Client{},
		Timeout: DefaultTimeout,
	}

	for i := range addr {
		c.clients[addr[i]] = nil
	}

	return &c
}

func (c *Client) getConn() (string, *redis.Client, error) {
	for addr, rclient := range c.clients {
		if rclient != nil {
			return addr, rclient, nil
		}
	}

	for addr := range c.clients {
		rclient, err := redis.DialTimeout("tcp", addr, c.Timeout)
		if err == nil {
			c.clients[addr] = rclient
			return addr, rclient, nil
		}
	}

	return "", nil, errors.New("no connectable endpoints")
}

func doCmd(rclient *redis.Client, cmd string, args ...interface{}) *redis.Resp {
	start := time.Now()
	r := rclient.Cmd(cmd, args...)
	if Debug && r.Err == nil {
		agg.Agg(strings.ToUpper(cmd), time.Since(start).Seconds())
	}
	return r
}

// Cmd do the redis cmd directly
func (c *Client) Cmd(cmd string, args ...interface{}) *redis.Resp {
	for i := 0; i < 3; i++ {
		addr, rclient, err := c.getConn()
		if err != nil {
			return redis.NewResp(err)
		}

		r := doCmd(rclient, cmd, args...)
		if err := r.Err; err != nil {
			if r.IsType(redis.IOErr) {
				err := rclient.Close()
				_ = err
				rclient = nil
				c.clients[addr] = nil
				continue
			}
		}
		return r
	}

	return redis.NewResp(errors.New("could not find usable endpoint"))
}
