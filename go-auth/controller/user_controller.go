package controller

import (
	"GO-AUTH/config"
	"GO-AUTH/helper"
	models "GO-AUTH/model"
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var userInput models.User

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&userInput); err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusBadRequest, response)
        return
    }
    defer r.Body.Close()

    var user models.User
    if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
        switch err {
        case gorm.ErrRecordNotFound:
            response := map[string]string{"message": "Email atau Password salah"}
            helper.ResponseJSON(w, http.StatusUnauthorized, response)
            return
        default:
            response := map[string]string{"message": err.Error()}
            helper.ResponseJSON(w, http.StatusInternalServerError, response)
            return
        }
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
        response := map[string]string{"message": "Email atau Password salah"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }
	if user.Role != "admin" {
        response := map[string]string{"message": "Anda tidak memiliki izin untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }

    //proses pembuatan jwt
    expTime := time.Now().Add(time.Minute * 15)
    token, err := createJWTToken(user.Email, expTime)
    if err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    //mengembalikan token ke cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Path:     "/",
        Value:    token,
        HttpOnly: true,
    })
    response := map[string]string{"message": "login berhasil"}
    helper.ResponseJSON(w, http.StatusOK, response)
}

// createJWTToken digunakan untuk membuat token JWT
func createJWTToken(email string, expTime time.Time) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = email
    claims["exp"] = expTime.Unix()

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte(config.JWT_KEY))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if userInput.Role == "" {
        userInput.Role = "user" // Tetapkan peran default jika tidak ada yang diberikan
    }

	var existingUser models.User
	if err := models.DB.Where("email = ?", userInput.Email).First(&existingUser).Error; err == nil {
		// Jika email sudah ada, respons gagal karna email sudah ada
		response := map[string]string{"message": "Email sudah terdaftar"}
		helper.ResponseJSON(w, http.StatusConflict, response)
		return
	}
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func Logout(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Path:     "/",
        Value:    "",
        HttpOnly: true,
		MaxAge: -1,
    })
    response := map[string]string{"message": "Berhasil Logout"}
    helper.ResponseJSON(w, http.StatusOK, response)
}