package mongo

import (
	"github.com/hnzhang/social-network/pkg"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mgo.Session, config *root.MongoConfig) *UserService {
	collection := session.DB(config.DbName).C("User")
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (p *UserService) CreateUser(u *root.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}
	return p.collection.Insert(&user)
}

func (p *UserService) GetUserByEmail(email string) (error, root.User) {
	model := userModel{}
	err := p.collection.find(bson.M{"email": email}).One(&model)
	return err, root.User{
		Id:       model.Id.Hex(),
		Username: model.Username,
		Password: "-",
	}
}
