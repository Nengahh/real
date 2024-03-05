package jawaban

import (
	"context"
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/internal/domians"
)

type JawabanPesertaUseCase struct {
	jawabanPesertaRepo repository.JawabanPesertaRepoInterface
}

type JawabanPesertaUseCaseInterface interface {
	CreateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error)
	GetJawabanPesertaByID(ctx context.Context, id int) (*domians.ParticipantAnswer, error)
	UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error)
	DeleteJawabanPeserta(ctx context.Context, id int) error
}

func NewJawabanPesertaUseCase(jawabanPesertaRepo repository.JawabanPesertaRepoInterface) JawabanPesertaUseCase {
	return JawabanPesertaUseCase{jawabanPesertaRepo: jawabanPesertaRepo}
}

func (uc JawabanPesertaUseCase) CreateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error) {
	return uc.jawabanPesertaRepo.CreateJawabanPeserta(ctx, jawabanPeserta)
}

func (uc JawabanPesertaUseCase) GetJawabanPesertaByID(ctx context.Context, id int) (*domians.ParticipantAnswer, error) {
	return uc.jawabanPesertaRepo.GetJawabanPesertaByID(ctx, id)
}

func (uc JawabanPesertaUseCase) UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error) {
	return uc.jawabanPesertaRepo.UpdateJawabanPeserta(ctx, jawabanPeserta)
}

func (uc JawabanPesertaUseCase) DeleteJawabanPeserta(ctx context.Context, id int) error {
	return uc.jawabanPesertaRepo.DeleteJawabanPeserta(ctx, id)
}
