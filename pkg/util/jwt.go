package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-gin-smallproject/pkg/setting"
	"time"
)

//指定加密密钥
//指定被保存在token中的实体对象，Claims 结构体。需要内嵌jwt.StandardClaims。这个结构体是用来保存信息的。
//根据数据产生token：根据传入的信息，组装成一个Claims结构体对象，再从对象中获取token
//根据token解析数据：解析出token所对应的interface{}，再使用断言解析出Claims对象，取数据

var jwtSecret = []byte(setting.JwtSecret)

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	exprieTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: exprieTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err

}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
