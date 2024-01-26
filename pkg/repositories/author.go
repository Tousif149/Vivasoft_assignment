package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"

	"gorm.io/gorm"
)

type authorRepo struct {
	db *gorm.DB
}

func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}
func (repo *authorRepo) GetAuthors(authorID uint) []models.AuthorDetail {
	var author []models.AuthorDetail
	var err error

	if authorID != 0 {
		err = repo.db.Where("id = ?", authorID).Find(&author).Error
	} else {
		err = repo.db.Find(&author).Error
	}
	if err != nil {
		return []models.AuthorDetail{}
	}
	return author
}
func (repo *authorRepo) CreateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}
func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}
func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	var Author models.AuthorDetail
	if err := repo.db.Where("id = ?", authorID).Delete(&Author).Error; err != nil {
		return err
	}
	return nil
}
