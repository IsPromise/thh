package cmd

import (
	"fmt"
	"github.com/leancodebox/goose/preferences"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tool:send_email",
		Short: "邮件发送",
		Run:   runSendEmail,
	})
}

func runSendEmail(_ *cobra.Command, _ []string) {
	preferences.GetString("MAIL_DRIVER")
	host := preferences.GetString("MAIL_HOST")
	port := preferences.GetInt("MAIL_PORT")
	user := preferences.GetString("MAIL_USERNAME")
	password := preferences.GetString("MAIL_PASSWORD")
	preferences.GetString("MAIL_FROM_ADDRESS")
	preferences.GetString("MAIL_ENCRYPTION")
	if len(host) == 0 {
		fmt.Println("未设置邮箱地址，无法发送")
		return
	}
	d := gomail.NewDialer(host, port, user, password)
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(user, "test")) //这种方式可以添加别名，即“XX官方”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	mailTo := []string{"email@email.com"}
	m.SetHeader("To", mailTo...)   //发送给多个用户
	m.SetHeader("Subject", "测试数据") //设置邮件主题
	m.SetBody("text/html", "测试数据") //设置邮件正文
	// 如果需要发送文件
	//m.Attach()
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("发送成功")
	}
}
