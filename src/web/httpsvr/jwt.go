package httpsvr

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
}

func GenerateToken(claims *AuthClaims, secret []byte) (string, error) {
	// 生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secret)
	if err != nil {return "", err}

	// 返回Token字符串
	return result, nil
}

func ParseToken(token string, secret []byte) (*AuthClaims, error) {
	claims := &AuthClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token)(interface{}, error){
		return []byte(secret), nil
	})

	return claims, err
}

