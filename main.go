package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"time"
)

func init() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		return
	}
	_ = log.ReplaceLogger(logger)
	log.Info("项目启动")
}
func main() {
	for {
		// todo 改成30分钟监测一次
		time.Sleep(30 * time.Second)
		checkExpired()
	}
}

func checkExpired() {
	// 判断文件是否存在
	var limitDateFile = "/mofahezi/limit_date.txt"
	const LAYOUT = "2006-01-02 15:04:05"
	_, err := os.Stat(limitDateFile)
	if err == nil {
		log.Debug("过期文件存在")
		readFile, _ := os.ReadFile(limitDateFile)
		limitDateStr := string(readFile)
		log.Debug("过期时间为:" + limitDateStr)
		limitDate, _ := time.Parse(LAYOUT, limitDateStr)
		now := time.Now()
		log.Debug("现在时间为:" + now.Format(LAYOUT))
		if limitDate.Before(now) {
			log.Debug("已经过期")
			log.Debug("执行关机,代码没写")
		} else {
			log.Debug("没有过期")
		}
	}
	if os.IsNotExist(err) {
		fmt.Println("过期文件不存在")
		file, _ := os.OpenFile(limitDateFile, os.O_WRONLY, 0666)
		// 获取当前日期
		now := time.Now()
		nowStr := now.Format(LAYOUT)
		log.Debug("写入当前时间:" + nowStr)
		file.WriteString(nowStr)
	}
}
