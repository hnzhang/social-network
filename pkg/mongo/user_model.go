package mongo

import (
	"github.com/google/uuid"
	"github.com/hnzhang/social-network/pkg"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userModel struct {
	Id           bson.ObjectId `bson:"_id, omitempty`
	Username     string
	Email        string
	PasswordHash string
	Salt         string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *root.User) (*userModel, error) {
	user := userModel{Username: u.Username}
	err := user.setSaltedPassword(u.Password)
	return &user, nil
}

func (u *userModel) comparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *userModel) setSaltedPassword(password string) error {
	salt := uuid.New().String()
	passwordBytes := []byte(password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash[:])
	u.Salt = salt

	return nil
}
