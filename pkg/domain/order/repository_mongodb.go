package order

import (
	"github.com/FernandoCagale/c4-order/internal/errors"
	"github.com/FernandoCagale/c4-order/pkg/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "order"
	DATABASE   = "c4-order-database"
)

type MongodbRepository struct {
	session *mgo.Session
}

func NewMongodbRepository(session *mgo.Session) *MongodbRepository {
	return &MongodbRepository{session}
}

func (repo *MongodbRepository) FindAll() (orders []*entity.Customer, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.Find(nil).All(&orders)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return orders, nil
}

func (repo *MongodbRepository) FindById(ID string) (order *entity.Customer, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.FindId(ID).One(&order)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return order, nil
}

func (repo *MongodbRepository) Create(e *entity.Customer) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	change, err := repo.FindById(e.Code)
	switch err {
	case errors.ErrNotFound:
		err = coll.Insert(e)
	case nil:
		err = coll.Update(change, bson.M{"$push": bson.M{"orders": bson.M{"$each": e.Orders}}})
	default:
		return errors.ErrInternalServer
	}

	if err != nil {
		if mgo.IsDup(err) {
			return errors.ErrConflict
		}
		return errors.ErrInternalServer
	}
	return nil
}

func (repo *MongodbRepository) DeleteById(ID string) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.RemoveId(ID)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return errors.ErrNotFound
		default:
			return errors.ErrInternalServer
		}
	}

	return nil
}
