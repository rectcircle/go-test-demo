package article

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/mattn/go-sqlite3"
	"github.com/rectcircle/go-test-demo/02-mock/domain"
)

type repository struct {
	db           *sql.DB
	insertStmt   *sql.Stmt
	findByIDStmt *sql.Stmt
}

func NewRepository(db *sql.DB) (domain.ArticleRepository, error) {
	insertStmt, err := db.Prepare(`insert into articles(author, title, tags, content) values(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	findByIDStmt, err := db.Prepare(`select author, title, tags, content from articles where id = ?`)
	if err != nil {
		return nil, err
	}
	return &repository{
		db:           db,
		insertStmt:   insertStmt,
		findByIDStmt: findByIDStmt,
	}, nil
}

func (a *repository) FindByID(id int64) (*domain.Article, error) {
	r := a.findByIDStmt.QueryRow(id)
	if err := r.Err(); err != nil {
		return nil, err
	}
	var author, title, _tags, content string
	if err := r.Scan(&author, &title, &_tags, &content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrRecordNotFound
		}
		return nil, err
	}
	tags := strings.Split(_tags, ",")
	return &domain.Article{
		ID:      id,
		Author:  author,
		Title:   title,
		Tags:    tags,
		Content: content,
	}, nil
}

func (r *repository) Create(a *domain.Article) (int64, error) {
	_tags := strings.Join(a.Tags, ",")
	res, err := r.insertStmt.Exec(a.Author, a.Title, _tags, a.Content)
	if err != nil {
		if errors.Is(sqlite3.ErrConstraintUnique, err) {
			return 0, domain.ErrConstraintUnique
		}
		return 0, err
	}
	return res.LastInsertId()
}
