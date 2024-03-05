package quiz

import "time"

type (
	CreateQuizParam struct {
		Judul        string    `json:"judul"`
		Deskripsi    string    `json:"deskripsi"`
		WaktuMulai   time.Time `json:"waktu_mulai"`
		WaktuSelesai time.Time `json:"waktu_selesai"`
	}

	UpdateQuizParam struct {
		ID           int       `json:"id"`
		Judul        string    `json:"judul"`
		Deskripsi    string    `json:"deskripsi"`
		WaktuMulai   time.Time `json:"waktu_mulai"`
		WaktuSelesai time.Time `json:"waktu_selesai"`
	}
)
