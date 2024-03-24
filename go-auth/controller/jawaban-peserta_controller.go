package controller

import (
	"GO-AUTH/helper"
	models "GO-AUTH/model"
	"encoding/json"
	"net/http"
)

func CreateJawabanPeserta(w http.ResponseWriter, r *http.Request) { 
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
    var Inputjawaban models.Jawaban_Peserta

    // Decode data JSON dari body request ke dalam struktur Quiz
    err := json.NewDecoder(r.Body).Decode(&Inputjawaban)
    if err != nil {
        response := map[string]string{"message": "Gagal memproses data jawaba"+err.Error()}
        helper.ResponseJSON(w, http.StatusBadRequest, response)
        return
    }

    // Simpan data Quiz ke dalam database
    if err := models.DB.Create(&Inputjawaban).Error; err != nil {
        response := map[string]string{"message": "Gagal membuat jawaban"+ err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // Kirim respons bahwa quiz berhasil dibuat
    response := map[string]string{"message": "Jawaban berhasil disimpan"}
    helper.ResponseJSON(w, http.StatusCreated, response)
}
