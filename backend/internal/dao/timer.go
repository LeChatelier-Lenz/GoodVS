package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"goodvs/internal/service"
	"time"
)

type TimerService struct {
	ticker *time.Ticker
	done   chan bool
	db     DBMS
}

func NewTimerService(interval time.Duration, db DBMS) *TimerService {
	return &TimerService{
		ticker: time.NewTicker(interval),
		done:   make(chan bool),
		db:     db,
	}
}

func (ts *TimerService) Start() {
	go func() {
		for {
			select {
			case <-ts.ticker.C:
				// 定时执行查询操作
				fmt.Println("Time to check...")
				ts.QueryProductFollowed()
			case <-ts.done:
				// 如果收到 done 信号，退出定时任务
				ts.ticker.Stop()
				return
			}
		}
	}()
}

func (ts *TimerService) Stop() {
	ts.done <- true
}

// QueryProductFollowed 查询所有被关注的商品
func (ts *TimerService) QueryProductFollowed() {
	queries, err := ts.db.GetCheckingList()
	if err != nil {
		logrus.Error("GetCheckingList failed", err)
		return
	}
	if len(queries) == 0 {
		logrus.Info("[Normal] No product to check")
		return
	}
	// 调用爬虫脚本检查价格
	result, err := service.CheckByCrawler(queries)
	if err != nil {
		logrus.Error("CheckByCrawler failed", err)
		return
	}
	// 更新数据库中的价格
	for _, v := range result {
		tx := ts.db.Begin()
		err = ts.db.UpdateProductPrice(v)
		if err != nil {
			logrus.Error("UpdateProductPrice failed", err)
			tx.Rollback()
			continue
		}
		tx.Commit()
		for _, q := range queries {
			if q.ProductId == v.ProductId {
				if v.Price <= q.Price && v.Price != 0 {
					// 价格降低，发送邮件
					logrus.Info("[Normal] Product price dropped")
					reqs, err := ts.db.GatherEmailInfo(v.ProductId, q.Price, v.Price)
					if err != nil {
						logrus.Error("GatherEmailInfo failed", err)
						continue
					}
					for _, req := range reqs {
						err = service.ES.SendEmail(req)
						if err != nil {
							logrus.Error("SendEmail failed", err)
							continue
						}
					}
					logrus.Info("[Normal] Email sent")
				}
			}
		}
	}
	logrus.Info("[Normal] Product price updated")
}
