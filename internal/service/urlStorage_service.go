package service

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/viettrung2103/bookmark-management-lesson/internal/repository"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/utils"
)

//go:generate mockery --name=ShortenUrl --filename=shortenurl.go

// ShortenUrl represents the shorten url service
type ShortenUrl interface {
	ShortenUrl(ctx context.Context, url string, exp int64) (string, error)
	GetLinkFromCode(ctx context.Context, urlCode string) (string, error)
	GetURL(ctx context.Context, urlCode string) (string, error)
}

type shortenUrlService struct {
	repo   repository.UrlStorage
	keygen utils.KeyGenerator
}

// NewShortenUrl returns a new ShortenUrl
func NewShortenUrl(repo repository.UrlStorage, keygen utils.KeyGenerator) ShortenUrl {
	return &shortenUrlService{
		repo:   repo,
		keygen: keygen,
	}
}

const (
	urlCodeLength = 7
)

// ShortenUrl shortens a url
func (s *shortenUrlService) ShortenUrl(ctx context.Context, url string, exp int64) (string, error) {
	// tao key
	key := s.keygen.GenerateKey(urlCodeLength)

	res, err := s.repo.GetURL(ctx, key)

	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}

	if res != "" {
		return s.ShortenUrl(ctx, url, exp)
	}

	// put key into redis
	err = s.repo.StoreURL(ctx, key, url, time.Duration(exp)*time.Second)
	if err != nil {
		return "", err
	}
	return key, nil
}

var ErrCodeDoesNotExist = errors.New("code does not exist")

func (s *shortenUrlService) GetLinkFromCode(ctx context.Context, urlCode string) (string, error) {
	url, err := s.repo.GetURL(ctx, urlCode)
	if errors.Is(err, redis.Nil) {
		return "", ErrCodeDoesNotExist
	}

	return url, err
}

var (
	UrlNotFound = errors.New("url not found")
)

func (s *shortenUrlService) GetURL(ctx context.Context, urlCode string) (string, error) {
	// call repo to get url
	url, err := s.repo.GetURL(ctx, urlCode)
	if err == nil {
		if errors.Is(err, redis.Nil) {
			return "", UrlNotFound
		}
	}

	return url, err

}
