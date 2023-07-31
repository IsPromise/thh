package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "md_5",
		Short: "",
		Run:   runMd5,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runMd5(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")
	fmt.Println("success")
	fmt.Println(md5String("/forgelogin/api/forge/getUserList" + "ifang_php_authorization_session=fguu0PoDCx2QicizgadheBRC4wyYADdQNR99hkQx; PHPSESSID=ST-9808-93oKzmBrk31bOmJyWROd-passport-cloud-58v5-cn"))
	fmt.Println(md5String("/forgelogin/api/forge/getUserList" + "10.249.146.114:51136"))
}

func md5String(input string) string {

	// 创建一个MD5哈希对象
	hash := md5.New()

	// 将字符串转换为字节数组并计算哈希值
	hash.Write([]byte(input))
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashValue)
	return hashString
}
