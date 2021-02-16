package syntax

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
	"gopkg.in/ini.v1"
)

// goroutineの実行制御数設定
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess2(ctx context.Context) {
	// 待たずに処理をキャンセルする
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}

	// AcquireでLockする。
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		// MustXXXは必須。
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func ThirdParty() {
	// go get golang.org/x/sync/semaphore
	ctx76 := context.TODO()
	go longProcess2(ctx76)
	go longProcess2(ctx76)
	go longProcess2(ctx76)
	time.Sleep(5 * time.Second)

	// go get gopkg.in/ini.v1
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
