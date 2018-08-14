package tools

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"time"
)

var db *redis.Database

func RedisInit() {
	db = redis.New(service.Config{
		Network:     "",
		Addr:        "127.0.0.1:6379",
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   10,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      ""})
	// optionally configure the bridge between your redis server

	// use go routines to query the database
	//db.Async(true)
	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	// fmt.Printf("sess type:%T\n", sess)
}

func GetSess(key int) (Sess *sessions.Sessions) {
	Sess = sessions.New(sessions.Config{Cookie: "sessionscookieid",
		Expires: time.Duration(key) * time.Minute})
	Sess.UseDatabase(db)
	return Sess
}
