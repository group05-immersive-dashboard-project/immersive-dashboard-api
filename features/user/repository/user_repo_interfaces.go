package repository

type UserRepository interface {
	Insert(user UserEntity) error
	Select(userID uint) (UserEntity, error)
	SelectAll(userID uint) ([]UserEntity, error)
	Update(userID uint, updatedUser UserEntity) error
	Delete(userID uint) error
	Login(email, password string) (UserEntity, string, error)
}
