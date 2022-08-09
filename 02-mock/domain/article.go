package domain

type Article struct {
	ID      int64
	Author  string
	Title   string
	Tags    []string
	Content string
}

//go:generate mockgen -destination=./mock/mock_article_repository.go -package=mock github.com/rectcircle/go-test-demo/02-mock/domain ArticleRepository

type ArticleRepository interface {
	FindByID(id int64) (*Article, error)
	Create(*Article) (int64, error)
}

type ArticleService interface {
	Publish(author string, title string, tags []string, content string) (*Article, error)
	Get(id int64) (*Article, error)
}
