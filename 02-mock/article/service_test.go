package article

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rectcircle/go-test-demo/02-mock/domain"
	"github.com/rectcircle/go-test-demo/02-mock/domain/mock"
)

func Test_service_Get(t *testing.T) {
	want := domain.Article{
		ID:      1,
		Author:  "author",
		Title:   "title",
		Tags:    []string{"go"},
		Content: "content",
	}

	// 准备 Mock 控制器。
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 构造一个 Mock 的 ArticleRepository 接口的实现 m。
	//   该实现的代码由 mockgen 命令生成
	//   该实现的 mock 的函数的返回值通过 m.EXPECT() 方法构造
	m := mock.NewMockArticleRepository(ctrl)
	// 声明，使用 1 调用 m.FindByID 时，返回 want。
	m.EXPECT().FindByID(gomock.Eq(int64(1))).Return(&want, nil)
	// 声明，使用非 1 调用 m.FindByID 时，返回 没有发现错误。
	m.EXPECT().FindByID(gomock.Not(int64(1))).Return(nil, domain.ErrRecordNotFound)

	// 构造待测实例，将 mock 对象 m 传递给该实例
	s, _ := NewService(m)
	// 执行测试
	t.Run("success", func(t *testing.T) {
		got, err := s.Get(1)
		if err != nil {
			t.Fatalf("s.Get(1) err want nil, got %s", err)
		}
		if reflect.DeepEqual(got, want) {
			t.Fatalf("s.Get(1) want %+v, got %+v", want, got)
		}
	})
	t.Run("notFound", func(t *testing.T) {
		_, err := s.Get(2)
		if err == nil {
			t.Fatalf("s.Get(2) err want %s, got nil", domain.ErrRecordNotFound)
		}
	})
}

func Test_service_Publish(t *testing.T) {
	want := domain.Article{
		ID:      1,
		Author:  "author",
		Title:   "title",
		Tags:    []string{"go"},
		Content: "content",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockArticleRepository(ctrl)
	data := map[int64]domain.Article{}
	id := int64(1)
	m.EXPECT().FindByID(gomock.Any()).DoAndReturn(func(id int64) (*domain.Article, error) {
		if a, ok := data[id]; ok {
			return &a, nil
		} else {
			return nil, domain.ErrRecordNotFound
		}
	})
	m.EXPECT().Create(gomock.Any()).DoAndReturn(func(a *domain.Article) (int64, error) {
		a.ID = id
		id += 1
		data[a.ID] = *a
		return a.ID, nil
	})

	s, _ := NewService(m)
	t.Run("success", func(t *testing.T) {
		got, err := s.Publish(want.Author, want.Title, want.Tags, want.Content)
		if err != nil {
			t.Fatalf("s.Publish(1) err want nil, got %s", err)
		}
		want.ID = got.ID
		if reflect.DeepEqual(got, want) {
			t.Fatalf("s.Publish(1) want %+v, got %+v", want, got)
		}
	})
}
