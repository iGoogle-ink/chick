package orm

import "github.com/go-redis/redis/v7"

// Redis redis config.
type Redis struct {
	Addr     []string
	Password string
	DB       int
}

func InitRedisCluster(c *Redis) (r *redis.ClusterClient) {
	r = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addr,
		Password: c.Password,
	})
	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}
	return r

}
