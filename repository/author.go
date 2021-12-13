package repository

import "bootcamp-homework/model"

// InMemoryAuthorRepository TODO: make goroutine safe
type InMemoryAuthorRepository struct {
	authors map[int]*model.Author
}

func NewInMemoryAuthorRepository() *InMemoryAuthorRepository {
	return &InMemoryAuthorRepository{make(map[int]*model.Author)}
}

func (repository *InMemoryAuthorRepository) All() []*model.Author {
	authors := make([]*model.Author, len(repository.authors))

	index := 0
	for _, author := range repository.authors {
		authors[index] = author
		index++
	}

	return authors
}

func (repository *InMemoryAuthorRepository) Read(authorID int) (*model.Author, error) {
	author, ok := repository.authors[authorID]
	if !ok {
		return nil, model.ErrAuthorNotFound
	}
	return author, nil
}

func (repository *InMemoryAuthorRepository) Create(author *model.Author) (*model.Author, error) {
	author.ID = len(repository.authors) + 1
	repository.authors[author.ID] = author
	return author, nil
}

func (repository *InMemoryAuthorRepository) Update(authorID int, author *model.Author) error {
	if _, ok := repository.authors[authorID]; !ok {
		return model.ErrAuthorNotFound
	}
	author.ID = authorID
	repository.authors[authorID] = author
	return nil
}

func (repository *InMemoryAuthorRepository) Delete(authorID int) error {
	delete(repository.authors, authorID)
	return nil
}
