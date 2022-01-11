package taskmanager

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	taskmanager "task-manager/Models"
	"time"

	"github.com/dgrijalva/jwt-go"
)


func CreateToken(userid uint32,usertype int) (string, error){
	fmt.Println(usertype)
	atClaims := jwt.MapClaims{}
  	atClaims["authorized"] = true
  	atClaims["user_id"] = userid
	atClaims["user_type"] = usertype
	fmt.Println(usertype)
  	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
  	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil

}


func ExtractToken(r *http.Request) string{
	bearToken := r.Header.Get("Authorization")
	fmt.Println(bearToken)
	strArr := strings.Split(bearToken," ")
	if len(strArr) == 2{
		return strArr[1]
	}

	return ""
}


func VerifyToken(r *http.Request) (*jwt.Token, error){
	tokenString := ExtractToken(r)
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _ , ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil

	})
	if err != nil {
		return nil, err
	 }
	return token, nil
   
}


func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	fmt.Println(err)
	if err != nil {
	   return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	   return err
	}
	return nil
  }
  

  func ExtractTokenMetadata(r *http.Request) (*taskmanager.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
	   return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

	   userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	   userType, err := strconv.ParseInt(fmt.Sprintf("%.f",claims["user_type"]),10,64)
	   fmt.Println(err)
	   if err != nil {
		   fmt.Print("here")
		  return nil, err
	   }

	   return &taskmanager.AccessDetails{
		  UserId:   userId,
		  UserType: int(userType),
	   }, nil
	}

	return nil, err
  }
  


  
  