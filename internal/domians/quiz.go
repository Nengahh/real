package domians

import "time"

type (
	Quiz struct {
		ID           int       `gorm:"primary_key"`
		Judul        string    `gorm:"column:judul"`
		Deskripsi    string    `gorm:"column:deskripsi"`
		WaktuMulai   time.Time `gorm:"column:waktu_mulai"`
		WaktuSelesai time.Time `gorm:"column:waktu_selesai"`
	}
)

func (Quiz) TableName() string {
	return "quiz"
}
