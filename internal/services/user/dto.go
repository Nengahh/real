package user

type (
	RegisterUserParam struct {
		Nama     string `json:"nama" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Role     string `json:"role"`
		Password string `json:"password" binding:"required"`
	}

	UseCaseRegisterResult struct {
		User RegisterUserParam `json:"user"`
	}

	SuccessLoginUser struct {
		Email       string `json:"email"`
		AccessToken string `json:"accessToken"`
	}

	LoginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
)
