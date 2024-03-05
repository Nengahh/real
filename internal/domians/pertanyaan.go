package domians

type (
	Pertanyaan struct {
		ID           int    `gorm:"primary_key"`
		Pertanyaan   string `gorm:"column:pertanyaan"`
		OpsiJawaban  string `gorm:"column:opsi_jawaban"`
		JawabanBenar int    `gorm:"column:jawaban_benar"`
		IDQuiz       int    `gorm:"column:id_quiz"`
	}
)

func (Pertanyaan) TableName() string {
	return "pertanyaan"
}
