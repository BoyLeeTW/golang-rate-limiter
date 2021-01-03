package ratelimiter

import (
	"golang-rate-limiter/internal/mock"
	"golang-rate-limiter/internal/models"
	"reflect"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/mock/gomock"
)

func TestRateLimiter_CheckExceed(t *testing.T) {
	t.Run("should return error if received error", func(t *testing.T) {
		want := models.RateLimiterResponse{}
		wantErr := memcache.ErrServerError

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheService := mock.NewMockMemcachedServiceInterface(ctrl)

		rateLimiter := NewRateLimiter(RateLimiterConfig{
			PeriodSecond: 60,
			LimitCount:   60,
		}, mockedMemcacheService)

		mockedMemcacheService.
			EXPECT().
			IncrementOrAdd("ip", uint64(1), int32(60)).
			Return(uint64(0), memcache.ErrServerError).
			Times(1)

		got, err := rateLimiter.CheckExceed("ip")

		if err != wantErr {
			t.Errorf("RateLimiter.CheckExceed() error = %v, wantErr %v", err, wantErr)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RateLimiter.CheckExceed() = %v, want %v", got, want)
		}
	})

	t.Run("should return not exceed with current count < limit and received no error", func(t *testing.T) {
		want := models.RateLimiterResponse{
			IsExceed:     false,
			CurrentCount: 30,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheService := mock.NewMockMemcachedServiceInterface(ctrl)

		rateLimiter := NewRateLimiter(RateLimiterConfig{
			PeriodSecond: 60,
			LimitCount:   60,
		}, mockedMemcacheService)

		mockedMemcacheService.
			EXPECT().
			IncrementOrAdd("ip", uint64(1), int32(60)).
			Return(uint64(30), nil).
			Times(1)

		got, err := rateLimiter.CheckExceed("ip")

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RateLimiter.CheckExceed() = %v, want %v", got, want)
		}
	})

	t.Run("should return not exceed with current count > limit and received no error", func(t *testing.T) {
		want := models.RateLimiterResponse{
			IsExceed:     true,
			CurrentCount: 61,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheService := mock.NewMockMemcachedServiceInterface(ctrl)

		rateLimiter := NewRateLimiter(RateLimiterConfig{
			PeriodSecond: 60,
			LimitCount:   60,
		}, mockedMemcacheService)

		mockedMemcacheService.
			EXPECT().
			IncrementOrAdd("ip", uint64(1), int32(60)).
			Return(uint64(61), nil).
			Times(1)

		got, err := rateLimiter.CheckExceed("ip")

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RateLimiter.CheckExceed() = %v, want %v", got, want)
		}
	})

	t.Run("should return not exceed with current count = limit and received no error", func(t *testing.T) {
		want := models.RateLimiterResponse{
			IsExceed:     false,
			CurrentCount: 60,
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheService := mock.NewMockMemcachedServiceInterface(ctrl)

		rateLimiter := NewRateLimiter(RateLimiterConfig{
			PeriodSecond: 60,
			LimitCount:   60,
		}, mockedMemcacheService)

		mockedMemcacheService.
			EXPECT().
			IncrementOrAdd("ip", uint64(1), int32(60)).
			Return(uint64(60), nil).
			Times(1)

		got, err := rateLimiter.CheckExceed("ip")

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RateLimiter.CheckExceed() = %v, want %v", got, want)
		}
	})
}
