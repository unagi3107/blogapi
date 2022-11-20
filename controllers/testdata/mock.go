package testdata

import "github.com/ura3107/blogapi/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return ArticleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return ArticleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return ArticleTestData[0], nil
}

func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return ArticleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}
