package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-key-keep-it-safe")

// 🧠 تعلم: دالة مساعدة لإرسال JSON للـ Frontend
func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	// 1. تشفير الباسوورد
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	// 2. الحفظ في الداتا بيز
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", body.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "اسم المستخدم موجود مسبقاً", http.StatusBadRequest)
		return
	}
	sendJSON(w, http.StatusCreated, map[string]string{"message": "تم التسجيل بنجاح"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	var id int
	var hash string
	// 1. جلب الباسوورد المشفر من الداتا بيز
	err := db.QueryRow("SELECT id, password FROM users WHERE username = ?", body.Username).Scan(&id, &hash)
	
	// 2. مقارنة الباسوورد اللي دخله اليوزر مع المشفر
	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(body.Password)) != nil {
		http.Error(w, "بيانات الدخول خاطئة", http.StatusUnauthorized)
		return
	}

	// 3. توليد JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  id,
		"username": body.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)

	sendJSON(w, http.StatusOK, map[string]string{"token": tokenString})
}