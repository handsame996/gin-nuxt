package casbinEnf

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
	"sync"
	"time"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func GetCasbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		// 创建Casbin执行器
		a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/", true)
		if err != nil {
			log.Fatalf("初始化适配器失败: %v", err)
		}

		syncedCachedEnforcer, err := casbin.NewSyncedCachedEnforcer("configs/model.conf", a)
		if err != nil {
			log.Fatalf("创建执行器失败: %v", err)
		}

		// 启用缓存
		syncedCachedEnforcer.EnableCache(true)

		// 每10秒自动刷新策略
		syncedCachedEnforcer.StartAutoLoadPolicy(10 * time.Second)

		// 加载策略
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			log.Fatalf("加载策略失败: %v", err)
		}
	})

	return syncedCachedEnforcer
}
