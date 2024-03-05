package jawaban

type ParticipantAnswerDTO struct {
	ID             int `json:"id"`
	IDUser         int `json:"idUser"`
	IDQuiz         int `json:"idQuiz"`
	IDPertanyaan   int `json:"idPertanyaan"`
	JawabanPeserta int `json:"jawabanPeserta"`
	Skor           int `json:"skor"`
}
