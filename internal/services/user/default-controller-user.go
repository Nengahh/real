package user

import (
	"fmt"
	"log"
	"net/http"
	"real_nimi_project/internal/adapter/dto"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type (
	Controller struct {
		Uc UserUseCaseInterface
	}

	ContrllerInterface interface {
		Register(
			ctx context.Context,
			payload RegisterUserParam,
		) (*dto.Response, error)

		GetAll(
			ctx context.Context,
		) (*dto.Response, error)

		LoginUser(
			ctx context.Context,
			email, password, role string,
		) (SuccessLoginUser, error)

		DeleteUserByName(
			ctx context.Context,
			name string,
		) (*dto.Response, error)
	}
)

func (ctrl Controller) Register(
	ctx context.Context,
	payload RegisterUserParam,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.Uc.RegisterUser(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Register is success",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) GetAll(
	ctx context.Context,
) (*dto.Response, error) {
	start := time.Now()
	res, err := ctrl.Uc.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		res,
		"list of users",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) LoginUser(
	ctx context.Context,
	email, password, role string,
) (SuccessLoginUser, error) {
	user, tokenString, err := ctrl.Uc.LoginUser(ctx, email, password, role)
	if err != nil {
		return SuccessLoginUser{}, err
	}
	response := SuccessLoginUser{
		Email:       user.Email,
		AccessToken: tokenString,
	}

	return response, nil
}

func (ctrl Controller) DeleteUserByName(
	ctx context.Context,
	name string,
) (*dto.Response, error) {
	start := time.Now()

	err := ctrl.Uc.DeleteUserByName(ctx, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.NewSuccessResponse(
		nil,
		"User deleted successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

// func (ctrl Controller) CheckAdminExists() (bool, error) {
//     // Lakukan query ke database untuk mencari pengguna dengan peran admin
//     // Misalkan kita menggunakan ORM (Object-Relational Mapping) seperti GORM

//     // Lakukan query untuk mencari pengguna dengan peran admin
//     var user User
//     err := ctrl.Db.Model(&User{}).Where("role = ?", "admin").First(&user).Error
//     if err != nil {
//         if err == gorm.ErrRecordNotFound {
//             // Jika tidak ditemukan pengguna dengan peran admin, kembalikan nilai false
//             return false, nil
//         }
//         // Jika terjadi kesalahan lain saat query, kembalikan error
//         return false, err
//     }

//     // Jika ditemukan pengguna dengan peran admin, kembalikan nilai true
//     return true, nil
// }

// func (ctrl Controller) DeleteUserByName(
// 	ctx context.Context,
// 	name string,
// ) (*dto.Response, error){
// 	start := time.Now()

// 	err := ctrl.Uc.DeleteUserByName(ctx,name)
// 	if err != nil{
// 		log.Println(err)
// 		return nil,err
// 	}

// 	return dto.NewSuccessResponse(
// 		nil,
// 		"user berhasil  di delet",
// 		fmt.Sprint(time.Since(start).Milliseconds(),"ms."),
// 	), nil

// }

func Logout(c *gin.Context) {
	c.SetCookie(
		"token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
