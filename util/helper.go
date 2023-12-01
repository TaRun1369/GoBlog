package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKet = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{ // claims is the payload of the token
		// in easier language claims is the data we want to send in the token
		Issuer : issuer,// issuer is the person who is sending the token
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // token will expire after 24 hours can be changed
	})

	return claims.SignedString([]byte(SecretKet)) // signedstring does the signing of the token
	// []byte(SecretKet) is the secret key 
}

func ParseJwt(cookie string)(string,error){
	token,err:= jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func(t *jwt.Token)(interface{},error){
		return []byte(SecretKet),nil // []byte does that it converts the string to byte
		// byte is the datatype in golang which is used to store the data in binary format
		// we are storing in binary format because it is more secure
	})

	if err != nil || !token.Valid{
		return "",err
	}
	claims:= token.Claims.(*jwt.StandardClaims) // standardclaims is the type of the claims it does the typecasting of the claims 
	// typecasting is needed for getting the data from the claims
	return claims.Issuer,nil 


}