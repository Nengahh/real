package repository

import (
	"context"
	"database/sql"
	"errors"
	"real_nimi_project/internal/domians"

	"gorm.io/gorm"
)

type JawabanPesertaRepo struct {
	db *gorm.DB
}

type JawabanPesertaRepoInterface interface {
	CreateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error)
	GetJawabanPesertaByID(ctx context.Context, id int) (*domians.ParticipantAnswer, error)
	UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error)
	DeleteJawabanPeserta(ctx context.Context, id int) error
}

func NewJawabanPesertaRepo(db *gorm.DB) JawabanPesertaRepo {
	return JawabanPesertaRepo{db: db}
}

func (repo JawabanPesertaRepo) CreateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error) {
	err := repo.db.WithContext(ctx).
		Create(&jawabanPeserta).
		Error
	return jawabanPeserta, err
}

func (repo JawabanPesertaRepo) GetJawabanPesertaByID(ctx context.Context, id int) (*domians.ParticipantAnswer, error) {
	jawabanPeserta := &domians.ParticipantAnswer{}
	err := repo.db.WithContext(ctx).
		Where("id = ?", id).
		First(jawabanPeserta).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, sql.ErrNoRows
	}

	return jawabanPeserta, err
}

func (repo JawabanPesertaRepo) UpdateJawabanPeserta(ctx context.Context, jawabanPeserta *domians.ParticipantAnswer) (*domians.ParticipantAnswer, error) {
	err := repo.db.WithContext(ctx).
		Model(&domians.ParticipantAnswer{}).
		Where("id = ?", jawabanPeserta.ID).
		Updates(jawabanPeserta).
		Error
	return jawabanPeserta, err
}

func (repo JawabanPesertaRepo) DeleteJawabanPeserta(ctx context.Context, id int) error {
	jawabanPeserta := &domians.ParticipantAnswer{}
	err := repo.db.WithContext(ctx).
		Where("id = ?", id).
		First(jawabanPeserta).
		Error

	if err != nil {
		return err
	}

	err = repo.db.WithContext(ctx).
		Delete(jawabanPeserta).
		Error
	return err
}
