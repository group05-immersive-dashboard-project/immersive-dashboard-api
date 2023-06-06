package repository

import (
	"errors"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Delete implements ClassRepository.
func (cq *classQuery) Delete(classID uint) error {
	deleteOpr := cq.db.Delete(&Class{}, classID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete class")
	}

	return nil
}

// Insert implements ClassRepository.
func (cq *classQuery) Insert(class ClassEntity) (uint, error) {
	classModel := EntityToModel(class)

	createOpr := cq.db.Create(&classModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return classModel.ID, nil
}

// Select implements ClassRepository.
func (cq *classQuery) Select(classID uint) (ClassEntity, error) {
	var class Class

	queryResult := cq.db.Preload("Mentees").First(&class, classID)
	if queryResult.Error != nil {
		return ClassEntity{}, queryResult.Error
	}

	classEntity := ModelToEntity(class)

	return classEntity, nil
}

// SelectAll implements ClassRepository.
func (cq *classQuery) SelectAll() ([]ClassEntity, error) {
	var classes []Class

	queryResult := cq.db.Preload("Mentees").Find(&classes)
	if queryResult.Error != nil {
		return []ClassEntity{}, queryResult.Error
	}

	var classEntities []ClassEntity
	for _, class := range classes {
		claasEntity := ModelToEntity(class)
		classEntities = append(classEntities, claasEntity)
	}

	return classEntities, nil
}

// Update implements ClassRepository.
func (cq *classQuery) Update(classID uint, updatedClass ClassEntity) error {
	var class Class

	queryResult := cq.db.First(&class, classID)
	if queryResult.Error != nil {
		return errors.New(queryResult.Error.Error() + ", failed to get class")
	}

	classToUpdate := EntityToModel(updatedClass)
	updateOpr := cq.db.Model(&class).Updates(classToUpdate)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update class")
	}

	return nil
}

func New(db *gorm.DB) ClassRepository {
	return &classQuery{
		db: db,
	}
}
