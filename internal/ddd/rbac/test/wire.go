//go:build wireinject
// +build wireinject

package test

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/LyricTian/gin-admin/v10/internal/wirex"
	"github.com/google/wire"
)

func BuildRepo(ctx context.Context) (*RepoTest, func(), error) {
	wire.Build(
		wirex.InitDB,
		//entity.EntitySet,
		repo.RepoSet,
		wire.Struct(new(RepoTest), "*"),
	) // end
	return &RepoTest{}, nil, nil
}
