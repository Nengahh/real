package user

import (
	"context"
	"real_nimi_project/internal/adapter/repository"
	"real_nimi_project/internal/domians"
	helper "real_nimi_project/utils"
	// "real_nimi_project/utils/helper"
)

type (
	UserUseCase struct {
		userRepo repository.UserRepoInterface
	}

	UserUseCaseInterface interface {
		RegisterUser(
			ctx context.Context,
			payload RegisterUserParam,
		) (result UseCaseRegisterResult, err error)

		GetAll(
			ctx context.Context,
		) ([]domians.User, error)

		LoginUser(
			ctx context.Context,
			email, password, role string,
		) (*domians.User, string, error)

		DeleteUserByName(
			ctx context.Context,
			name string,
		) error
	}
)

func (uc UserUseCase) RegisterUser(
	ctx context.Context,
	payload RegisterUserParam,
) (result UseCaseRegisterResult, err error) {

	// helper.HasPass(payload.Password)
	hashPass := helper.HasPass(payload.Password)
	user, err := uc.userRepo.RegisterUser(
		ctx,
		&domians.User{
			Nama:     payload.Nama,
			Email:    payload.Email,
			Role:     "student",
			Password: hashPass,
		})
	result.User = RegisterUserParam{
		Nama:  user.Nama,
		Email: user.Email,
		Role:  user.Role,
	}
	return result, err
}

func (uc UserUseCase) GetAll(
	ctx context.Context,
) ([]domians.User, error) {
	return uc.userRepo.GetAllUser(ctx)
}

func (uc UserUseCase) LoginUser(
	ctx context.Context,
	email, password, role string,
) (*domians.User, string, error) {
	user, err := uc.userRepo.LoginUser(ctx, email, password, role)
	if err != nil {
		return nil, "", err
	}

	// verify hashed password
	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	//Generate token JWT
	tokenString, errToken := helper.GenerateToken(user.ID, email, role)
	if errToken != nil {
		return nil, "", err
	}

	// c := ctx.Value("gin_context").(*gin.Context)
	// c.SetCookie(
	// 	"token",
	// 	tokenString,
	// 	3600,
	// 	"/",
	// 	"",
	// 	false,
	// 	true,
	// )

	return user, tokenString, nil

}

// func (uc UserUseCase) DeleteUserByName(
//     ctx context.Context,
//     name string,
// ) error {
//     return uc.userRepo.DeleteUserByName(ctx, name)
// }

func (uc UserUseCase) DeleteUserByName(
	ctx context.Context,
	name string,
) error {
	return uc.userRepo.DeleteUserByName(ctx, name)
}

func (uc UserUseCase) LogoutUser(ctx context.Context) error {
	// Karena kita tidak memiliki akses langsung ke cookie di sisi server,
	// kita hanya perlu mengembalikan respons yang menginstruksikan
	// browser untuk menghapus atau mengatur ulang cookie.
	return nil
}
