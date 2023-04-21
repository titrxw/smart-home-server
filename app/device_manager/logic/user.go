package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/device_manager/repository"
)

const UserInfoCacheKey = `user:info:%d`

type User struct {
	Abstract
}

func (l User) CreateUser(userName string, email string, password string) (*model.User, error) {
	db := l.GetDefaultDb()
	userRepository := repository.Repository.User

	user := userRepository.GetByUserName(db, userName)
	if user != nil {
		return nil, exception.NewResponseError("该用户名已存在")
	}
	user = userRepository.GetByEmail(db, email)
	if user != nil {
		return nil, exception.NewResponseError("该email已存在")
	}

	user = &model.User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
	if userRepository.CreateUser(db, user) == false {
		return nil, exception.NewResponseError("用户创建失败")
	}

	return user, nil
}

func (l User) ResetUserCache(ctx context.Context, user *model.User) error {
	cacheKey := fmt.Sprintf(UserInfoCacheKey, user.ID)
	data := l.GetDefaultRedis().Del(ctx, cacheKey)
	if data.Err() != nil && data.Err() != redis.Nil {
		return data.Err()
	}

	return nil
}

func (l User) GetUserById(ctx context.Context, userId model.UID) (*model.User, error) {
	cacheKey := fmt.Sprintf(UserInfoCacheKey, userId)
	data := l.GetDefaultRedis().Get(ctx, cacheKey)
	if data.Err() != nil && data.Err() != redis.Nil {
		return nil, data.Err()
	}

	if data.Val() == "" {
		user := repository.Repository.User.GetById(l.GetDefaultDb(), userId)
		if user == nil {
			return nil, exception.NewResponseError("用户不存在")
		}

		encodeData, err := helper.JsonEncode(user)
		if err != nil {
			return nil, err
		}
		result := l.GetDefaultRedis().Set(ctx, cacheKey, encodeData, 0)
		if result.Err() != nil {
			return nil, result.Err()
		}

		return user, nil
	} else {
		user := new(model.User)
		err := helper.JsonDecode(data.Val(), &user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func (l User) GetByEmail(email string) *model.User {
	return repository.Repository.User.GetByEmail(l.GetDefaultDb(), email)
}

func (l User) GetByEmailAndPwd(email string, password string) (*model.User, error) {
	user := l.GetByEmail(email)
	if user == nil {
		return nil, exception.NewResponseError("该email不存在")
	}
	if user.MakeHashPassword(password, user.Salt) != user.Password {
		return nil, exception.NewResponseError("email或者密码错误")
	}

	return user, nil
}

func (l User) UpdateUser(user *model.User) error {
	if !repository.Repository.User.UpdateUser(l.GetDefaultDb(), user) {
		return exception.NewResponseError("更新用户信息失败")
	}

	return nil
}
