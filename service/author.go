package service

import "bootcamp-homework/model"

type AuthorRepository interface {
	All() []*model.Author
	Read(authorID int) (*model.Author, error)
	Create(author *model.Author) (*model.Author, error)
	Update(authorID int, author *model.Author) error
	Delete(authorID int) error
}

type AuthorService struct {
	authorRepo AuthorRepository
}

func NewAuthorService(authorRepo AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (service *AuthorService) GetAllAuthors() []*model.Author {
	return service.authorRepo.All()
}

func (service *AuthorService) CreateAuthor(author *model.Author) (*model.Author, error) {
	return service.authorRepo.Create(author)
}

func (service *AuthorService) GetAuthorByID(authorID int) (*model.Author, error) {
	return service.authorRepo.Read(authorID)
}

func (service *AuthorService) UpdateAuthor(authorID int, author *model.Author) error {
	return service.authorRepo.Update(authorID, author)
}

func (service *AuthorService) DeleteAuthor(authorID int) (*model.Author, error) {
	author, err := service.authorRepo.Read(authorID)
	if err != nil {
		return nil, err
	}

	err = service.authorRepo.Delete(authorID)
	if err != nil {
		return nil, err
	}

	return author, nil
}

