package jawaban

import (
	"fmt"
	"log"
	"real_nimi_project/internal/adapter/dto"
	"real_nimi_project/internal/domians"
	"time"

	"golang.org/x/net/context"
)

type Controller struct {
	Uc JawabanPesertaUseCaseInterface
}

type ControllerInterface interface {
	CreateJawabanPeserta(ctx context.Context, payload ParticipantAnswerDTO) (*dto.Response, error)
	GetJawabanPesertaByID(ctx context.Context, id int) (*dto.Response, error)
	UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*dto.Response, error)
	DeleteJawabanPeserta(ctx context.Context, id int) (*dto.Response, error)
}

func (ctrl Controller) CreateJawabanPeserta(ctx context.Context, payload ParticipantAnswerDTO) (*dto.Response, error) {
	start := time.Now()
	jawabanPeserta := &domians.ParticipantAnswer{
		IDUser:         payload.IDUser,
		IDQuiz:         payload.IDQuiz,
		IDPertanyaan:   payload.IDPertanyaan,
		JawabanPeserta: payload.JawabanPeserta,
		Skor:           payload.Skor,
	}
	result, err := ctrl.Uc.CreateJawabanPeserta(ctx, jawabanPeserta)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Jawaban peserta created successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) GetJawabanPesertaByID(ctx context.Context, id int) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.Uc.GetJawabanPesertaByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Jawaban peserta retrieved successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*dto.Response, error) {
	start := time.Now()
	updatedJawabanPeserta, err := ctrl.Uc.UpdateJawabanPeserta(ctx, jawabanPeserta)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		updatedJawabanPeserta,
		"Jawaban peserta updated successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) DeleteJawabanPeserta(ctx context.Context, id int) (*dto.Response, error) {
	start := time.Now()
	err := ctrl.Uc.DeleteJawabanPeserta(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		nil,
		"Jawaban peserta deleted successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}
