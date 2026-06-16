package intergration_test

//
//func TestGenPassEndpoint(t *testing.T) {
//	t.Parallel()
//
//	testCases := []struct {
//		name string
//
//		setupTestHTTP func(api api.Engine) *httptest.ResponseRecorder
//
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name: "normal case",
//			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
//				req, _ := http.NewRequest("GET", "/genpass", nil)
//				respRecorder := httptest.NewRecorder()
//				api.ServeHTTP(respRecorder, req)
//				return respRecorder
//			},
//
//			expectedStatusCode:   http.StatusOK,
//			expectedResponseBody: `{"password":`,
//		},
//		{
//			name: "wrong endpoint",
//			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
//				req, _ := http.NewRequest("POST", "/genpass", nil)
//				respRecorder := httptest.NewRecorder()
//				api.ServeHTTP(respRecorder, req)
//				return respRecorder
//			},
//
//			expectedStatusCode:   http.StatusNotFound,
//			expectedResponseBody: ``,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//			mockRedis := redis.InitMockRedis(t)
//			mockDB := sqldb.CreateTestDb(t)
//			mockEngine := gin.Default()
//
//			testApi := api.NewEngine(mockEngine, &api.Config{}, mockRedis, mockDB)
//			recorder := tc.setupTestHTTP(testApi)
//
//			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
//			assert.Contains(t, recorder.Body.String(), tc.expectedResponseBody)
//		})
//	}
//}
