package pertanyaan

type (
	CreateQuestionParam struct {
		Pertanyaan   string `json:"pertanyaan"`
		OpsiJawaban  string `json:"opsi_jawaban"`
		JawabanBenar int    `json:"jawaban_benar"`
		IDQuiz       int    `json:"id_quiz"`
	}

	UpdateQuestionParam struct {
		ID           int    `json:"id"`
		Pertanyaan   string `json:"pertanyaan"`
		OpsiJawaban  string `json:"opsi_jawaban"`
		JawabanBenar int    `json:"jawaban_benar"`
		IDQuiz       int    `json:"id_quiz"`
	}
)
