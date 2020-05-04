//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-order/api/routers"
	"github.com/FernandoCagale/c4-order/event"
	"github.com/FernandoCagale/c4-order/internal/datastore"
	"github.com/FernandoCagale/c4-order/pkg"
	"github.com/google/wire"
	"gopkg.in/mgo.v2"
)

func SetupApplication(*mgo.Session) (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}

func SetupMongoDB() (*mgo.Session, error) {
	wire.Build(datastore.Set)
	return nil, nil
}

func SetupEvents(session *mgo.Session) (*event.OrderEvent, error) {
	wire.Build(event.Set)
	return nil, nil
}
