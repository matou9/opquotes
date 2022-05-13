package pub

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"opquotes/config"
	_ "opquotes/config"
	"time"
)

var Rds *redis.Client
var Ctx context.Context
func init(){
	Rds,_=NewRedis()
}

func NewRedis()(rds *redis.Client,err error){
	Ctx =context.Background()
	host:=config.Getstring("redis.host")
	pass:=config.Getstring("redis.password")
	db:=config.Getint("redis.db")
	rds= redis.NewClient(&redis.Options{
		Addr:         host,
		Password:		pass,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		PoolSize:     10,
		PoolTimeout:  10 * time.Second,
		DB:db,
	})
	if _,err = rds.Ping(Ctx).Result();err!=nil{
		return nil,err
	}
	return
}

/*
err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
*/