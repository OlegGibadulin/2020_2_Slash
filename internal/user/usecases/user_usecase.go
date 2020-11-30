package usecases

import (
	"context"
	. "github.com/go-park-mail-ru/2020_2_Slash/internal/consts"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/user"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

type UserUsecase struct {
	userBlockClient user.UserBlockClient
}

func NewUserUsecase(client user.UserBlockClient) *UserUsecase {
	return &UserUsecase{
		userBlockClient: client,
	}
}

func (uu *UserUsecase) Create(modelUser *models.User) *errors.Error {
	grpcUser, err := uu.userBlockClient.Create(context.Background(),
		user.ModelUserToGrpc(modelUser))
	if err != nil {
		customErr := errors.Get(ErrorCode(status.Code(err)))
		return customErr
	}

	err = copier.Copy(modelUser, user.GrpcUserToModel(grpcUser))
	if err != nil {
		return errors.New(CodeInternalError, err)
	}

	return nil
}

func (uu *UserUsecase) GetByEmail(email string) (*models.User, *errors.Error) {
	grpcUser, err := uu.userBlockClient.GetByEmail(context.Background(),
		&user.Email{Email: email})
	if err != nil {
		customErr := errors.Get(ErrorCode(status.Code(err)))
		return nil, customErr
	}

	return user.GrpcUserToModel(grpcUser), nil
}

func (uu *UserUsecase) GetByID(userID uint64) (*models.User, *errors.Error) {
	grpcUser, err := uu.userBlockClient.GetByID(context.Background(),
		&user.ID{ID: userID})
	if err != nil {
		customErr := errors.Get(ErrorCode(status.Code(err)))
		return nil, customErr
	}

	return user.GrpcUserToModel(grpcUser), nil
}

func (uu *UserUsecase) UpdateProfile(newUserData *models.User) (*models.User, *errors.Error) {
	grpcUser, err := uu.userBlockClient.UpdateProfile(context.Background(),
		user.ModelUserToGrpc(newUserData))
	if err != nil {
		customErr := errors.Get(ErrorCode(status.Code(err)))
		return nil, customErr
	}

	return user.GrpcUserToModel(grpcUser), nil
}

func (uu *UserUsecase) UpdateAvatar(userID uint64, newAvatar string) (*models.User, *errors.Error) {
	grpcUser, err := uu.userBlockClient.UpdateAvatar(context.Background(),
		&user.IdAvatar{
			Id:     &user.ID{ID: userID},
			Avatar: &user.Avatar{Avatar: newAvatar},
		})
	if err != nil {
		customErr := errors.New(ErrorCode(status.Code(err)), err)
		return nil, customErr
	}

	return user.GrpcUserToModel(grpcUser), nil
}

func (uu *UserUsecase) CheckPassword(user *models.User, password string) *errors.Error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(password)); err != nil {
		return errors.Get(CodeWrongPassword)
	}
	return nil
}

func (uu *UserUsecase) IsAdmin(userID uint64) (bool, *errors.Error) {
	user, err := uu.GetByID(userID)
	if err != nil {
		return false, err
	}

	return user.Role == Admin, nil
}
