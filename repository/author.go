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

func (repository *InMemoryAuthorRepository) Read(id int) (*model.Author, error) {
	author, ok := repository.authors[id]
	if !ok {
		return nil, ErrAuthorNotFound
	}
	return author, nil
}

func (repository *InMemoryAuthorRepository) Create(author *model.Author) (*model.Author, error) {
	author.ID = len(repository.authors) + 1
	repository.authors[author.ID] = author
	return author, nil
}

func (repository *InMemoryAuthorRepository) Update(id int, author *model.Author) error {
	if _, ok := repository.authors[id]; !ok {
		return ErrAuthorNotFound
	}
	author.ID = id
	repository.authors[id] = author
	return nil
}

func (repository *InMemoryAuthorRepository) Delete(id int) (*model.Author, error) {
	author, ok := repository.authors[id]
	if !ok {
		return nil, ErrAuthorNotFound
	}
	delete(repository.authors, id)
	return author, nil
}
