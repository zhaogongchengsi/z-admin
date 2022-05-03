package utils

import (
	"time"
	"z-admin/global"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Uid      uuid.UUID `json:"uid"`
	UserName string    `json:"username"`
	jwt.StandardClaims
}

func NewJWT(signingKey string) *JWT {
	return &JWT{
		SigningKey: []byte(signingKey),
	}
}

func CreateJwt(uid uuid.UUID, username string) (string, error) {
	nowTime := time.Now()                                                          // token 生效时间
	expireTime := nowTime.Add(time.Duration(global.JwtConf.ExpiresAt) * time.Hour) // token 有效时间 小时为单位

	claims := Claims{
		Uid:      uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: global.JwtConf.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString([]byte(global.JwtConf.SigningKey))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JwtConf.SigningKey), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
