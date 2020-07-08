package common

import (
	"fmt"
	"ginLearn/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Claim struct {
	Name string
	Pwd string
	Num string
	jwt.StandardClaims
}
var tokenKey = viper.GetString("token.secret")

func GetToken(device model.Device) (string,error){
		expireTime := time.Now().Add(7 * 24 * time.Hour)
		claims := &Claim{
			Name:device.Num,
			Pwd:device.Num,
			Num:device.Num,
			StandardClaims:jwt.StandardClaims{
				ExpiresAt:expireTime.Unix(),
				IssuedAt:time.Now().Unix(),
				Issuer:"jiayouba",
				Subject:"device token",
			},

		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
		tokenString,err := token.SignedString([]byte(tokenKey))
		return tokenString,err
}

func CheckToken(tokenString string) (bool, *Claim){
	claims := &Claim{}
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenKey),nil
	})
	if err != nil{
		return false,nil
	}
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		fmt.Printf("%v %v", claims.Num, claims.StandardClaims.ExpiresAt)
		return true,claims
	} else {
		fmt.Println(err)
		return false,nil
	}
}
