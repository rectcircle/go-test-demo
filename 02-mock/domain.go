package mockdemo

import "errors"

var (
	ErrRecordNotFound   = errors.New("record not found")
	ErrConstraintUnique = errors.New("constraint unique")
)

type Article struct {
	ID      int64
	Author  string
	Title   string
	Tags    []string
	Content string
}

type ArticleRepository interface {
	FindByID(id int64) (*Article, error)
	Create(*Article) (int64, error)
}

type ArticleService interface {
	Publish(author string, title string, tags []string, content string) (*Article, error)
	Get(id int64) (*Article, error)
}
