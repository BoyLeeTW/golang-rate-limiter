package pkg

import (
	"golang-rate-limiter/internal/mock"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/mock/gomock"
)

func TestMemcachedService_IncrementOrAdd(t *testing.T) {
	t.Run("should return error if received error other than ErrCacheMiss", func(t *testing.T) {
		wantValue := uint64(0)
		wantErr := memcache.ErrServerError

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Increment("key", uint64(1)).
			Return(uint64(0), memcache.ErrServerError).
			Times(1)

		got, err := mc.IncrementOrAdd("key", 1, 60)

		if err != wantErr {
			t.Errorf("MemcachedService.IncrementOrAdd() error = %v, wantErr %v", err, wantErr)
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})

	t.Run("should call addOrIncrement if receive ErrCacheMiss when increment", func(t *testing.T) {
		wantValue := uint64(1)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Increment("key", uint64(1)).
			Return(uint64(0), memcache.ErrCacheMiss).
			Times(1)
		mockedMemcacheClient.
			EXPECT().
			Add(&memcache.Item{Key: "key", Value: []byte("1"), Expiration: 60}).
			Return(nil).
			Times(1)

		got, err := mc.IncrementOrAdd("key", 1, 60)

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})

	t.Run("test increment successfully", func(t *testing.T) {
		wantValue := uint64(1)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Increment("key", uint64(1)).
			Return(uint64(1), nil).
			Times(1)
		mockedMemcacheClient.
			EXPECT().
			Add(gomock.Any()).
			Times(0)

		got, err := mc.IncrementOrAdd("key", 1, 60)

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})
}

func TestMemcachedService_AddOrIncrement(t *testing.T) {
	t.Run("should return error if received error other than ErrNotStored", func(t *testing.T) {
		wantValue := uint64(0)
		wantErr := memcache.ErrServerError

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Add(&memcache.Item{Key: "key", Value: []byte("1"), Expiration: 60}).
			Return(memcache.ErrServerError).
			Times(1)
		mockedMemcacheClient.
			EXPECT().
			Increment(gomock.Any(), gomock.Any()).
			Times(0)

		got, err := mc.addOrIncrement("key", 1, 60)

		if err != wantErr {
			t.Errorf("MemcachedService.IncrementOrAdd() error = %v, wantErr %v", err, wantErr)
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})

	t.Run("should call Increment if received error ErrNotStored", func(t *testing.T) {
		wantValue := uint64(1)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Add(&memcache.Item{Key: "key", Value: []byte("1"), Expiration: 60}).
			Return(memcache.ErrNotStored).
			Times(1)
		mockedMemcacheClient.
			EXPECT().
			Increment("key", uint64(1)).
			Return(uint64(1), nil).
			Times(1)

		got, err := mc.addOrIncrement("key", 1, 60)

		if err != nil {
			t.Errorf("should be nil")
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})

	t.Run("should return error if received error after calling Increment", func(t *testing.T) {
		wantValue := uint64(0)
		wantErr := memcache.ErrServerError

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

		mc := NewMemcachedService(mockedMemcacheClient)

		mockedMemcacheClient.
			EXPECT().
			Add(&memcache.Item{Key: "key", Value: []byte("1"), Expiration: 60}).
			Return(memcache.ErrNotStored).
			Times(1)
		mockedMemcacheClient.
			EXPECT().
			Increment("key", uint64(1)).
			Return(uint64(0), memcache.ErrServerError).
			Times(1)

		got, err := mc.addOrIncrement("key", 1, 60)

		if err != wantErr {
			t.Errorf("MemcachedService.IncrementOrAdd() error = %v, wantErr %v", err, wantErr)
			return
		}

		if got != wantValue {
			t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
		}
	})

	// t.Run("test increment successfully", func(t *testing.T) {
	// 	wantValue := uint64(1)

	// 	ctrl := gomock.NewController(t)
	// 	defer ctrl.Finish()

	// 	mockedMemcacheClient := mock.NewMockMemcachedClientInterface(ctrl)

	// 	mc := NewMemcachedService(mockedMemcacheClient)

	// 	mockedMemcacheClient.
	// 		EXPECT().
	// 		Increment("key", uint64(1)).
	// 		Return(uint64(1), nil).
	// 		Times(1)
	// 	mockedMemcacheClient.
	// 		EXPECT().
	// 		Add(gomock.Any()).
	// 		Times(0)

	// 	got, err := mc.IncrementOrAdd("key", 1, 60)

	// 	if err != nil {
	// 		t.Errorf("should be nil")
	// 		return
	// 	}

	// 	if got != wantValue {
	// 		t.Errorf("MemcachedService.IncrementOrAdd() = %v, want %v", got, wantValue)
	// 	}
	// })
}
