package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	BaseHtmlTplFile     = "public/html/common/base.html"
	Error404HtmlTplFile = "public/html/common/404.html"
)

//验证用户名和密码的正确性
func CheckUserPassword(usr, psw string) bool {

	return false
}

// 获得可执行程序所在目录
func ExeDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	//Dir returns all but the last element of path,and the trailing slashes are removed
	return filepath.Dir(pathAbs), nil
}

//得到现在时刻的年，月，日，小时，并转化为字符串形式返回
func GetCurrentTime() string {
	t := time.Now()
	y, m, d := t.Date()
	h := t.Hour()
	//月，天，小时都必须为占两位数，不够则补零
	return fmt.Sprintf("%d%02d%02d%02d", y, m, d, h)
}
