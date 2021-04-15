package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Secret     = ""  // 加盐
	ExpireTime = 600 // token有效期 单位秒
)

func Jwtinit(secret string, expireTime int) {
	if expireTime <= 0 {
		panic("过期时间不能小于0")
	}
	if len(secret) == 0 {
		panic("密钥不能为空")
	}
	Secret = secret
	ExpireTime = expireTime
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
	UserCode  string `json:"user_code"`
}

func (c *JWTClaims) setExpiredAt(expiredAt int64) {
	c.ExpiresAt = expiredAt
}

func TokenObtain(claims *JWTClaims) (string, error) {
	claims.IssuedAt = time.Now().Unix()
	claims.setExpiredAt(time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func TokenVerify(strToken string) error {
	c := &JWTClaims{}
	token, err := jwt.ParseWithClaims(strToken, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return err
	}
	if err := token.Claims.Valid(); err != nil {
		return err
	}

	return nil
}

func GetTokenClaims(strToken string) (JWTClaims, error) {
	c := &JWTClaims{}
	token, err := jwt.ParseWithClaims(strToken, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return *c, err
	}
	if err := token.Claims.Valid(); err != nil {
		return *c, err
	}

	return *c, nil
}
