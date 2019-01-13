package Services

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	HandleCors(&w, req)
	var user User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	} else {
		json.Unmarshal(body, &user)
		var user_name = user.Username
		var pass_word = user.Password

		fmt.Println(user_name)
		fmt.Println(pass_word)
		if user_name == "admin" && pass_word == "admin" {

			_ = json.NewDecoder(req.Body).Decode(&user)

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"password": user.Password,
			})
			tokenString, error := token.SignedString([]byte("secret"))
			if error != nil {
				fmt.Println(error)
			}
			json.NewEncoder(w).Encode(JwtToken{Token: tokenString})

		} else {
			json.NewEncoder(w).Encode("Username or password error")
		}

	}

}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {

	HandleCors(&w, req)
	params := req.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		HandleCors(&w, req)
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 1 {
				token, error := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					fmt.Println("AUTHENTICATED USER")
					context.Set(req, "decoded", token.Claims)
					next(w, req)

				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			} else {
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func AuthenticateTest(w http.ResponseWriter, req *http.Request) {

	HandleCors(&w, req)
	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}
