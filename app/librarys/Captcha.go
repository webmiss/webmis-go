package librarys

import (
	"encoding/base64"
	"math/rand"
	"time"
	"webmis/app/util"
	"webmis/core"

	"github.com/mojocn/base64Captcha"
)

const txtChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var store = base64Captcha.DefaultMemStore

/* 验证码 */
type Captcha struct {
	core.Base
}

/* 获取字符 */
func (c *Captcha) GetCode(num int) string {
	code := make([]byte, num)
	for i := 0; i < num; i++ {
		code[i] = txtChars[rnd.Intn(len(txtChars))]
	}
	return string(code)
}

/* 获取数字 */
func (c *Captcha) GetNum(num int) string {
	code := make([]byte, num)
	for i := 0; i < num; i++ {
		code[i] = txtChars[rnd.Intn(10)+52]
	}
	return string(code)
}

/* 图形验证码 */
func (c *Captcha) Vcode(num int) (string, []byte) {
	driver := &base64Captcha.DriverString{
		Height:          40,
		Width:           140,
		NoiseCount:      1,   // 干扰线
		ShowLineOptions: 3,   // 干扰线类型 0:不显示，1：显示横线，2：显示斜线，3：显示曲线
		Length:          num, // 验证码长度
		Source:          txtChars,
	}
	cp := base64Captcha.NewCaptcha(driver, store)
	_, b64s, code, _ := cp.Generate()
	arr := util.Explode(",", b64s)
	img, _ := base64.StdEncoding.DecodeString(arr[1])
	return code, img
}
