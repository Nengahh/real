package domians

type (
	ParticipantAnswer struct {
		ID             int `gorm:"primary_key"`
		IDUser         int `gorm:"column:id_user"`
		IDQuiz         int `gorm:"column:id_quiz"`
		IDPertanyaan   int `gorm:"column:id_pertanyaan"`
		JawabanPeserta int `gorm:"column:jawaban_peserta"`
		Skor           int `gorm:"column:skor"`
	}
)

func (ParticipantAnswer) TableName() string {
	return "jawaban_peserta"
}
