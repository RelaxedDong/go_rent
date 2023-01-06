package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	OpenId string
	jwt.StandardClaims
}

var SecretKey = "\x7d\xef\x87\xd5\xf8\xbb\xff\xfc\x80\x91\x06\x91\xfd\xfc\xed\x69"

func GenerateToken(openid string) (string, error) {
	// https://www.jianshu.com/p/83f8b338683a
	nowTime := time.Now()
	expireTime := nowTime.Add(432000 * time.Second) //  5天过期
	// 有效载荷就是存放有效信息的地方
	claims := Claims{
		OpenId: openid, // 公共的声明，公共的声明可以添加任何的信息，一般添加用户的相关信息或其他业务需要的必要信息。
		StandardClaims: jwt.StandardClaims{ // 标准中注册的声明（Reserved claims）
			ExpiresAt: expireTime.Unix(), // jwt的过期时间，这个过期时间必须要大于签发时间
		},
	}
	//header： {
	//	"typ": "JWT", 声明类型，固定为
	//	jwt
	//	"alg": "HS256"
	//	声明签名算法，通常为HMAC SHA256
	//}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//func main() {
//	tokenId, err := GenerateToken("donghao")
//	fmt.Println(tokenId, err)
//
//	resp, e := ParseToken(tokenId)
//	fmt.Println(resp.OpenId, e)
//}
