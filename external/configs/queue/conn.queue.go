package queue

import (
	"sync"

	"github.com/hibiken/asynq"
)

var client *asynq.Client
var once sync.Once

func AsyncClientConn() *asynq.Client {
	once.Do(func() {
		redisConnection := asynq.RedisClientOpt{
			Addr: "localhost:6379",
		}

		client = asynq.NewClient(redisConnection)
	})
	return client
}

var server *asynq.Server
var onceServer sync.Once

func AsyncServerConn() *asynq.Server {
	onceServer.Do(func() {
		redisConnection := asynq.RedisClientOpt{
			Addr: "localhost:6379",
		}

		server = asynq.NewServer(redisConnection, asynq.Config{
			Concurrency: 10,
			// Specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6, // processed 60% of the time
				"default":  3, // processed 30% of the time
				"low":      1, // processed 10% of the time
			},
		})
	})
	return server
}
