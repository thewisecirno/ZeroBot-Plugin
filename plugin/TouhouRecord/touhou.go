// Package TouhouRecord
package TouhouRecord

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var examplelimit = ctxext.NewLimiterManager(time.Second*10, 1)

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.115.217.189:6379",
		Password: "qq31415926535--",
		DB:       0,
	})
	log.Println(rdb)
	engine := control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "东方角色语录随机发送",
		Help: "- 灵梦\n" +
			"- 琪露诺\n",
	})

	engine.OnFullMatch("灵梦在吗").
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			val, err := rdb.SMembers(context.Background(), "Remei").Result()
			if err != nil {
				ctx.SendChain(message.Text(err.Error()))
			}

			size := len(val)
			r := rand.New(rand.NewSource(time.Now().Unix()))
			ctx.SendChain(message.Text(val[r.Intn(size-1)]))
		})
	engine.OnFullMatch("琪露诺在吗").
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			val, err := rdb.SMembers(context.Background(), "Cirno").Result()
			if err != nil {
				ctx.SendChain(message.Text(err.Error()))
			}

			size := len(val)
			r := rand.New(rand.NewSource(time.Now().Unix()))
			ctx.SendChain(message.Text(val[r.Intn(size-1)]))
		})

	engine.OnFullMatch("秦心").
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("忠！诚！"))
		})
}
