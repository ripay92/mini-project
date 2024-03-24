package controller

import (
	"GO-AUTH/helper"
	models "GO-AUTH/model"
	"encoding/json"
	"net/http"
	"time"
)
func GetQuiz(w http.ResponseWriter, r *http.Request) {
	if !helper.IsLoggedIn(r) {
        response := map[string]string{"message": "Anda harus login untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }

    if !helper.IsAdmin(r) {
        response := map[string]string{"message": "Anda tidak memiliki izin admin untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusForbidden, response)
        return
    }

	var quizzes []models.Quiz
	if err := models.DB.Find(&quizzes).Error; err != nil {
		response := map[string]string{"message": "Gagal mengambil data quiz"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, quizzes)
}
func CreateQuiz(w http.ResponseWriter, r *http.Request) {
    if !helper.IsLoggedIn(r) {
        response := map[string]string{"message": "Anda harus login untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }

    if !helper.IsAdmin(r) {
        response := map[string]string{"message": "Anda tidak memiliki izin admin untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusForbidden, response)
        return
    }

    // proses pembuatan quiz
    var Inputquis models.Quiz

    // Decode data JSON dari body request ke dalam struktur Quiz
    err := json.NewDecoder(r.Body).Decode(&Inputquis)
    if err != nil {
        response := map[string]string{"message": "Gagal memproses data quiz"}
        helper.ResponseJSON(w, http.StatusBadRequest, response)
        return
    }

    // Simpan data Quiz ke dalam database
    if err := models.DB.Create(&Inputquis).Error; err != nil {
        response := map[string]string{"message": "Gagal membuat quiz"+ err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // Kirim respons bahwa quiz berhasil dibuat
    response := map[string]string{"message": "Quiz berhasil dibuat"}
    helper.ResponseJSON(w, http.StatusCreated, response)
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
    // Validasi apakah pengguna telah login dgn role admin
    if !helper.IsLoggedIn(r) {
        response := map[string]string{"message": "Anda harus login untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }

    if !helper.IsAdmin(r) {
        response := map[string]string{"message": "Anda tidak memiliki izin admin untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusForbidden, response)
        return
    }

    // Mendapatkan ID quiz dari URL
    quizID := r.URL.Query().Get("id")

    // Hapus data Quiz dari database berdasarkan ID
    if err := models.DB.Where("id = ?", quizID).Delete(&models.Quiz{}).Error; err != nil {
        response := map[string]string{"message": "Gagal menghapus quiz"}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // Kirim respons bahwa quiz berhasil dihapus
    response := map[string]string{"message": "Quiz berhasil dihapus"}
    helper.ResponseJSON(w, http.StatusOK, response)
}


func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
    // Validasi apakah pengguna telah login dgn role admin
    if !helper.IsLoggedIn(r) {
        response := map[string]string{"message": "Anda harus login untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusUnauthorized, response)
        return
    }

    if !helper.IsAdmin(r) {
        response := map[string]string{"message": "Anda tidak memiliki izin admin untuk mengakses ini"}
        helper.ResponseJSON(w, http.StatusForbidden, response)
        return
    }

    // Mendapatkan ID quiz dari URL
    quizID := r.URL.Query().Get("id")

    // Decode data JSON dari body request ke dalam struktur Quiz
    var updatedQuiz models.Quiz
    err := json.NewDecoder(r.Body).Decode(&updatedQuiz)
    if err != nil {
        response := map[string]string{"message": "Gagal memproses data quiz"}
        helper.ResponseJSON(w, http.StatusBadRequest, response)
        return
    }

    // Konversi string waktu_mulai dan waktu_selesai ke time.Time
    updatedQuiz.WaktuMulai, _ = time.Parse(time.RFC3339, updatedQuiz.WaktuMulai.Format(time.RFC3339))
    updatedQuiz.WaktuSelesai, _ = time.Parse(time.RFC3339, updatedQuiz.WaktuSelesai.Format(time.RFC3339))

    // Update data Quiz di database berdasarkan ID
    if err := models.DB.Model(&models.Quiz{}).Where("id = ?", quizID).Updates(updatedQuiz).Error; err != nil {
        response := map[string]string{"message": "Gagal memperbarui quiz"}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // Kirim respons bahwa quiz berhasil diperbarui
    response := map[string]string{"message": "Quiz berhasil diperbarui"}
    helper.ResponseJSON(w, http.StatusOK, response)
}

