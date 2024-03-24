package helper

import (
	"GO-AUTH/config"
	models "GO-AUTH/model"
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func checkAuthorization(r *http.Request, requireAdmin bool) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})
	if err != nil || !token.Valid {
		return false
	}

	// Mendapatkan email dari klaim token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	// Memeriksa apakah email pengguna memiliki peran admin jika diperlukan
	if requireAdmin {
		email := claims["email"].(string)
		var user models.User
		if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return false
		}
		if user.Role != "admin" {
			return false
		}
	}

	return true
}

// Fungsi untuk memeriksa apakah pengguna telah login
func IsLoggedIn(r *http.Request) bool {
	return checkAuthorization(r, false)
}

// Fungsi untuk memeriksa apakah pengguna memiliki peran admin
func IsAdmin(r *http.Request) bool {
	return checkAuthorization(r, true)
}