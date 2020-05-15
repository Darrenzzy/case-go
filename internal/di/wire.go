// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"order/internal/dao"
	"order/internal/server/grpc"
	"order/internal/server/http"
	"order/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, grpc.New, http.New, NewApp))
}
