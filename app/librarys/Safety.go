package librarys

import (
	"regexp"
	"webmis/app/config"
	"webmis/core"

	"github.com/golang-jwt/jwt/v5"
)

/* 验证类 */
type Safety struct {
	core.Base
}

/* 正则-公共 */
func (s *Safety) Config(name string, value string) bool {
	switch name {
	case "uname":
		return s.Test("^[a-zA-Z][a-zA-Z0-9\\_\\@\\-\\*\\&]{3,15}$", value)
	case "passwd":
		return s.Test("^[a-zA-Z0-9|_|@|-|*|&]{6,16}$", value)
	case "tel":
		return s.Test("^1\\d{10}$", value)
	case "email":
		return s.Test("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$", value)
	case "idcard":
		return s.Test("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", value)
	default:
		return false
	}
}

/* 正则-验证 */
func (s *Safety) Test(reg string, value string) bool {
	res, _ := regexp.MatchString(reg, value)
	return res
}

/* Base64-加密 */
func (s *Safety) Encode(param map[string]interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"param": param,
	})
	str, err := token.SignedString([]byte(config.Env().Key))
	if err != nil {
		return ""
	}
	return str
}

/* Base64-解密 */
func (s *Safety) Decode(token string) map[string]interface{} {
	str, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env().Key), nil
	})
	if claims, ok := str.Claims.(jwt.MapClaims); ok && str.Valid {
		return claims["param"].(map[string]interface{})
	} else if err != nil {
		return nil
	}
	return nil
}
