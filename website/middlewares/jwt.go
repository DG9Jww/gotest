package middlewares

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	SecretKey = []byte("1EC9074ACEAD1326E5068C4FB81CB221")	//密钥
	AccessTokenExpireDuration = time.Hour * 3    		//AccessToken过期时间
	RefreshTokenExpireDuration = time.Hour * 3    		//RefreshToken过期时间
)

type MyClaim struct{
	UserID		int
	UserRight	int 
	jwt.StandardClaims
}

//创建Token
func CreateToken(userID int,user_right int) (string,error) {
	claim := MyClaim{
		UserID: userID,
		UserRight: user_right,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),   //过期时间,必须填
			Issuer: "dachuang",   //签发人
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	//使用SignedString返回JWT字符串
	return token.SignedString(SecretKey)
}

//
func GetToken(userID int,user_right int) (atoken,rtoken string,err error) {
	claim := MyClaim{
		UserID: userID,
		UserRight: user_right,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),   //过期时间,必须填
			Issuer: "dachuang",   //签发人
		},
	}
	//这里为AccessToken
	atoken,err = jwt.NewWithClaims(jwt.SigningMethodHS256,claim).SignedString(SecretKey)
	if err != nil{
		return atoken,rtoken,err
	}
	//Refresh Token,不需要自定义字段，有个过期时间就好
	rtoken,err = jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
		Issuer: "dachuang",
	}).SignedString(SecretKey)
	return atoken,rtoken,err
}

//解析Token
func ParseToken(tokenString string) (*MyClaim,error) {
	var claim = new(MyClaim)
	token,err := jwt.ParseWithClaims(tokenString,claim, func(t *jwt.Token) (interface{}, error) {
		return SecretKey,nil
	})
	if err != nil{
		return nil,err
	}
	if token.Valid{
		return claim,nil
	}
	return nil,errors.New("invalid Token")
}

//验证Token是否合法
func AuthenticateToken(c *gin.Context){
	//1.先获取请求头中Authorization字段(看前端发起请求时token放哪)
	auth := c.Request.Header.Get("Authorization")
	if auth == ""{
		c.JSON(200,gin.H{
			"status":false,
			"msg":"请先登录",
		})
		c.Abort()
		return
	}

	//2.获取Token,具体看请求时Authorization字段怎么写的
	token := auth
	//3.解析token
	claim,err := ParseToken(token)

	if err != nil{
		c.JSON(200,gin.H{
			"status":false,
			"msg":"非法请求",
		})
		c.Abort()
		return
	}
	//4.将UserID信息保存到当前的c里面
	c.Set("userID",claim.UserID)
	c.Set("userRight",claim.UserRight)
	//后续操作都可通过c.Get("userID")
	c.Next()
}

//刷新Token
func RefreshToken(){

}