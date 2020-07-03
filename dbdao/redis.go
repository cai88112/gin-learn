package dbdao

import (
	"context"
	"fmt"
	"ginLearn/global"
	"github.com/go-redis/redis/v8"
)
var ctx = context.Background()
func Redis() {
	addr := global.Vp.GetString("redis.addr")
	pwd := global.Vp.GetString("redis.pwd")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB : 0,
	})
	pong, err := client.Ping(ctx).Result();
	if err != nil {
		fmt.Println("redis connect fail:", err)
	} else {
		fmt.Println("redis connect ping response:", pong)
		global.Redis = client
	}
}