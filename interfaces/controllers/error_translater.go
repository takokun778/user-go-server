package controllers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/takokun778/user-go-server/domains/models"
)

/*
  ドメインモデルのエラーをgRPC用のエラーに変換する
*/
func ErrorTranslate(error models.IBaseError) error {
	// エラーログはここで記述
	switch error.Code() {
	case 403:
		return status.Errorf(codes.Unauthenticated, error.Message())
	case 404:
		return status.Errorf(codes.NotFound, error.Message())
	default:
		return status.Errorf(codes.Unknown, error.Message())
	}
}
