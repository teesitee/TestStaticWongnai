package api_test

import (
	"net/http/httptest"
	"testing"
	"teststaticwongnai/api"
	"teststaticwongnai/api/mock_api"
	"teststaticwongnai/duplicate"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestAPIGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock_api.NewMockTestGetinfo(ctrl)
	h := api.Handler{
		ServiceRes: mock,
	}
	mock.EXPECT().ResGet().Return(nil, duplicate.Age{})
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	h.Getinfo(ginContext)
}
