package repository

type ClassRepository interface {
	Insert(class ClassEntity) (uint, error)
	Select(classID uint) (ClassEntity, error)
	SelectAll() ([]ClassEntity, error)
	Update(classID uint, updatedClass ClassEntity) error
	Delete(classID uint) error
}
