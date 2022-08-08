package mockdemo

type articleService struct {
	repository ArticleRepository
}

func NewArticleService(r ArticleRepository) (ArticleService, error) {
	return &articleService{
		repository: r,
	}, nil
}

func (s *articleService) Get(id int64) (*Article, error) {
	return s.repository.FindByID(id)
}

func (s *articleService) Publish(author string, title string, tags []string, content string) (*Article, error) {
	id, err := s.repository.Create(&Article{
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
