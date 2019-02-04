package tokencache

import (
	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/pkg/contracts"
)

type ReadCacheResponse struct {
	TokenResponse *msalbase.TokenResponse
	Account       contracts.IAccount
}

type ICacheManager interface {
	TryReadCache() *ReadCacheResponse
	CacheTokenResponse(tokenResponse *msalbase.TokenResponse) (contracts.IAccount, error)
	DeleteCachedRefreshToken() error
}
