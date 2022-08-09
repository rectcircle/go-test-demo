package article

import "github.com/rectcircle/go-test-demo/02-mock/domain"

type service struct {
	repository domain.ArticleRepository
}

func NewService(r domain.ArticleRepository) (domain.ArticleService, error) {
	return &service{
		repository: r,
	}, nil
}

func (s *service) Get(id int64) (*domain.Article, error) {
	return s.repository.FindByID(id)
}

func (s *service) Publish(author string, title string, tags []string, content string) (*domain.Article, error) {
	id, err := s.repository.Create(&domain.Article{
		ID:      0,
		Author:  author,
		Title:   title,
		Tags:    tags,
		Content: content,
	})
	if err != nil {
		return nil, err
	}
	return s.Get(id)
}
