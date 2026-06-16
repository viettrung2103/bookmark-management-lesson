package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/viettrung2103/bookmark-management-lesson/internal/app/service/mocks"
)

func TestShortenUrlHandler_Redirect(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest func(c *gin.Context)
		setupMockSvc func(ctx context.Context) *mocks.ShortenUrl

		exptectedStatus int
		expectedUrl     string
	}{
		{
			name: "normal case - success",
			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(
					http.MethodGet,
					"/shorten/1234567",
					nil,
				)
				ctx.Params = gin.Params{{Key: "code", Value: "1234567"}}
			},
			setupMockSvc: func(ctx context.Context) *mocks.ShortenUrl {
				serviceMock := mocks.NewShortenUrl(t)
				serviceMock.On("GetURL", ctx, "1234567").Return("https://google.com", nil)
				return serviceMock
			},
			exptectedStatus: http.StatusFound,
			expectedUrl:     "https://google.com",
		},
		{
			name: "normal case - success",
			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(
					http.MethodGet,
					"/shorten/1234567",
					nil,
				)
				ctx.Params = gin.Params{{Key: "code", Value: "1234567"}}
			},
			setupMockSvc: func(ctx context.Context) *mocks.ShortenUrl {
				serviceMock := mocks.NewShortenUrl(t)
				serviceMock.On("GetURL", ctx, "1234567").Return("https://google.com", nil)
				return serviceMock
			},
			exptectedStatus: http.StatusFound,
			expectedUrl:     "https://google.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			tc.setupRequest(ctx)

			mockSvc := tc.setupMockSvc(ctx)
			testHandler := NewShortenUrlHandler(mockSvc)
			testHandler.Redirect(ctx)

			assert.Equal(t, tc.exptectedStatus, rec.Code)
			assert.Equal(t, tc.expectedUrl, rec.Header().Get("Location"))
		})
	}
}
