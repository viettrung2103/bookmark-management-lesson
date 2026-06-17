package jwtutils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJWTGenerator(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		keyPath string

		expectedErrStr string
	}{
		{
			name:           "normal case",
			keyPath:        filepath.FromSlash("./test.private.pem"),
			expectedErrStr: "",
		},
		{
			name:           "err case - file not found",
			keyPath:        filepath.FromSlash("./non-exist.pem"),
			expectedErrStr: "open",
		},
		{
			name:           "err case - not a private key",
			keyPath:        filepath.FromSlash("./test.public.pem"),
			expectedErrStr: "structure error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewJWTGenerator(tc.keyPath)
			if err != nil {
				assert.ErrorContains(t, err, tc.expectedErrStr)
			}
		})
	}
}

func TestGenerateJWT(t *testing.T) {
	expectedToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0In0.Y2TzSZIwUbQy-xVpaWMmZCoToqbAWODyUhkpEMse3gFLnxmG71Y8jc7sZqNpb0VSaswJ7tGIhhKiVxIhM0OG8EAE81ZNRTap3aI8_cBGsKYkVsWyqU35JZd75pOF5455nNYLqu1L1X0dpCgXGEUgbi4SrZY_h2c6An4NPyC484aXyNRPTSPi5OKxUyYcuQ5EWjHlsQLx6SqA-5ntuYQzApWRN7VKGdUdN5_bwrfkkwYwHrenOs-yYZCyp74c5Ih4tS1ESmH83AcE9cvPnqBd9OURYNGTS2-_P_OhhDwTKEabPqwG0L6neY3OL7VUxAoRywbDFsHgQT1YR5a9ddMjgw"
	t.Parallel()
	gen, err := NewJWTGenerator(filepath.FromSlash("./test.private.pem"))
	if err != nil {
		t.Fatal("should not fail")
	}

	token, err := gen.GenerateJWT(map[string]any{
		"sub": "1234",
	})

	if err != nil {
		t.Fatal("should not fail")
	}
	assert.Equal(t, token, expectedToken)

}
