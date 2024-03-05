package repository

import (
	"context"
	"fmt"
	"real_nimi_project/internal/domians"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type (
	UserRepo struct {
		db *gorm.DB
	}

	UserRepoInterface interface {
		RegisterUser(
			ctx context.Context,
			cust *domians.User,
		) (domians.User, error)

		LoginUser(
			ctx context.Context,
			email string,
			password string,
			role string,
		) (*domians.User, error)

		GetAllUser(
			ctx context.Context,
		) ([]domians.User, error)

		DeleteUserByName(
			ctx context.Context,
			name string) error
	}

	// ctx contex.Context, user *domain.User) (*domain.User, error)
	// LoginUser(ctx context.Context, email, password string) (*domain.User, error)
	// LogoutUser(ctx context.Context, user *domain.User) error

)

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}

// func (repo UserRepo) RegisterUser(
// 	ctx context.Context,
// 	cust *domians.User,
// ) (*domians.User, error) {
// 	err := repo.db.WithContext(ctx).
// 		Create(&cust).
// 		Error
// 	return *cust, err
// }

func (repo UserRepo) RegisterUser(
	ctx context.Context,
	user *domians.User,
) (domians.User, error) {
	err := repo.db.WithContext(ctx).
		Create(&user).
		Error
	return *user, err
}

func (repo UserRepo) GetAllUser(
	ctx context.Context,
) ([]domians.User, error) {
	var cust []domians.User
	err := repo.db.WithContext(ctx).Find(&cust).
		Error
	return cust, err
}

func (repo UserRepo) LoginUser(
	ctx context.Context,
	email string,
	password string,
	role string,
) (*domians.User, error) {
	user := &domians.User{}

	// err := repo.db.WithContext(ctx).
	// 	Model(&domians.User{}).
	// 	Where("email = ? AND password = ? AND role = ?", email, password, role).
	// 	First(user).
	// 	Error

	err := repo.db.WithContext(ctx).
		Model(&domians.User{}).
		Where("email = ? AND role = ?", email, role).
		First(user).
		Error

	if err != nil {
		return nil, err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return nil, fmt.Errorf("incorrect password")
	}

	return user, err
}

func (repo UserRepo) DeleteUserByName(
	ctx context.Context,
	name string,
) error {
	user := &domians.User{}

	// Find the user by name
	err := repo.db.WithContext(ctx).
		Where("Nama = ?", name).
		First(user).
		Error

	if err != nil {
		return err
	}

	// Delete the user
	err = repo.db.WithContext(ctx).Delete(user).Error

	return err
}
