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

	engine.OnKeywordGroup([]string{"灵梦", "赤色杀人魔", "博丽灵梦", "Remei"}).
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

	engine.OnKeywordGroup([]string{"琪露诺", "⑨", "Cirno"}).
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

	engine.OnKeywordGroup([]string{"秦心", "秦大将军", "zb", "秦心大将军"}).
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			val, err := rdb.SMembers(context.Background(), "KokoroHataZB").Result()
			if err == redis.Nil || len(val) == 0 || val == nil {
				return
			}
			if err != nil {
				ctx.SendChain(message.Text(err.Error()))
			}
			size := len(val)
			r := rand.New(rand.NewSource(time.Now().Unix()))
			ctx.SendChain(message.Text(val[r.Intn(size-1)]))
		})

	//engine.OnKeywordGroup([]string{"车万皇帝", "车万皇上"}).
	//	SetBlock(true).
	//	Limit(ctxext.LimitByGroup).
	//	Handle(func(ctx *zero.Ctx) {
	//		ctx.Event.
	//			ctx.SendChain(message.Text(val[r.Intn(size-1)]))
	//	})

	engine.OnKeywordGroup([]string{"Test zero.context"}).
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("context.Event:", ctx.Event))
			ctx.SendChain(message.Text("context.State:", ctx.State))
		})
}
