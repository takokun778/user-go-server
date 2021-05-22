package controllers

import (
	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/domains/models"
)

type UserTranslater struct {
}

/*
  ドメインモデルのユーザーをgRPC用に変換する
*/
func (UserTranslater) Translate(user *models.User) (result *pb.User) {
	result = new(pb.User)
	result.Id = user.Id()
	result.Name = user.Name()
	return
}
