package tokencache

import (
	"errors"

	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/pkg/contracts"
)

type cacheManager struct {
	storageManager IStorageManager
	authParameters *msalbase.AuthParametersInternal
}

func CreateCacheManager(storageManager IStorageManager, authParameters *msalbase.AuthParametersInternal) ICacheManager {
	cache := &cacheManager{storageManager, authParameters}
	return cache
}

func (m *cacheManager) TryReadCache() *ReadCacheResponse {
	return nil
}

func (m *cacheManager) CacheTokenResponse(tokenResponse *msalbase.TokenResponse) (contracts.IAccount, error) {
	return nil, errors.New("not implemented")
}

func (m *cacheManager) DeleteCachedRefreshToken() error {
	return nil
}
