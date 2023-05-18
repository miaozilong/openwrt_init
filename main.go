package main

import (
	log "github.com/cihub/seelog"
	"github.com/jordan-wright/email"
	"net/smtp"
	"os"
	"os/exec"
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
	log.Debug("进入main方法")
	for {
		log.Debug("开始休息30分")
		time.Sleep(30 * time.Minute)
		log.Debug("结束休息30分")
		checkExpired()
	}
}

func checkExpired() {
	// 判断文件是否存在
	var limitDateFile = "/mofahezi/limit_date.txt"
	_, err := os.Stat(limitDateFile)
	nowStr := getNowStr()
	if err == nil {
		log.Debug("过期文件存在")
		readFile, _ := os.ReadFile(limitDateFile)
		limitDateStr := string(readFile)
		log.Debug("过期时间为:" + limitDateStr)
		log.Debug("现在时间为:" + nowStr)
		if limitDateStr < nowStr {
			log.Debug("已经过期")
			log.Debug("执行关机")
			hostname, _ := os.Hostname()
			sendMail("魔法盒子已过期", "主机名称:"+hostname+",过期时间为:"+limitDateStr+",现在时间为:"+nowStr)
			cmd := exec.Command("halt")
			err := cmd.Run()
			if err != nil {
				_ = log.Error(err)
			}
		} else {
			log.Debug("没有过期")
		}
	}
	if os.IsNotExist(err) {
		log.Debug("过期文件不存在")
		_, _ = os.OpenFile(limitDateFile, os.O_WRONLY, 0666)
		defaultExpireDate := getDefaultExpireStr()
		log.Debug("写入默认时间:" + defaultExpireDate)
		err := os.WriteFile(limitDateFile, []byte(defaultExpireDate), 0666)
		if err != nil {
			_ = log.Error(err)
		}
	}
}

func getNowStr() string {
	const LAYOUT = "2006-01-02 15:04:05"
	var cstZone = time.FixedZone("CST", 8*3600) // 东八区
	// 获取当前日期
	now := time.Now()
	ret := now.In(cstZone).Format(LAYOUT)
	return ret
}

func getDefaultExpireStr() string {
	const LAYOUT = "2006-01-02 15:04:05"
	var cstZone = time.FixedZone("CST", 8*3600) // 东八区
	// 获取当前日期
	defaultTime := time.Now().Add(24 * (365 + 15) * time.Hour)
	ret := defaultTime.In(cstZone).Format(LAYOUT)
	return ret
}

/*
示例代码：
var sub = "hello"
var content = "hello"
sendMail(&sub, &content)
*/
func sendMail(subject string, content string) {
	e := email.NewEmail()
	e.From = "mofahezi@gmail.com"
	e.To = []string{"mofahezi@gmail.com", "miao.zilong@outlook.com"}
	e.Subject = subject
	e.Text = []byte(content)
	err2 := e.Send("smtp.gmail.com:587", smtp.PlainAuth("",
		"mofahezi@gmail.com",
		"phiwhpsymqggsotu",
		"smtp.gmail.com"))
	if err2 != nil {
		_ = log.Error("发送失败")
	}
}
