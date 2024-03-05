package domians

type (
	User struct {
		ID       uint   `gorm:"primary_key"`
		Nama     string `gorm:"column:nama"`
		Email    string `gorm:"column:email"`
		Password string `gorm:"password"`
		Role     string `gorm:"role"`
	}
)

func (User) TableName() string {
	return "user"
}
