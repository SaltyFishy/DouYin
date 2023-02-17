package ftp

import (
	"DouYin/src/conf"
	"github.com/dutchcoders/goftp"
	"log"
	"time"
)

var Ftp *goftp.FTP

func InitFTP() {
	//获取到ftp的链接
	var err error
	Ftp, err = goftp.Connect(conf.ConnectConfig)
	if err != nil {
		log.Printf("获取到FTP链接失败！！！")
	}
	log.Printf("获取到FTP链接成功%v：", Ftp)
	//登录
	err = Ftp.Login(conf.FtpUser, conf.FtpPasssword)
	if err != nil {
		log.Printf("FTP登录失败！！！")
	}
	log.Printf("FTP登录成功！！！")
	//维持长链接
	go keepAlive()
}

func keepAlive() {
	time.Sleep(time.Duration(conf.HeartbeatTime) * time.Second)
	Ftp.Noop()
}
