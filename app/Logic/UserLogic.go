package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"

	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

const USER_INFO_CACHE_KEY = `user:info:%d`

type UserLogic struct {
	LogicAbstract
}

func (userLogic UserLogic) CreateUser(userName string, mobile string, password string) (*model.User, error) {
	db := userLogic.GetDefaultDb()
	userRepository := repository.Repository.UserRepository

	user := userRepository.GetByUserName(db, userName)
	if user != nil {
		return nil, errors.New("该用户名已存在")
	}
	user = userRepository.GetByMobile(db, mobile)
	if user != nil {
		return nil, errors.New("该手机号已存在")
	}

	user = &model.User{
		UserName: userName,
		Mobile:   mobile,
		Password: password,
	}
	if userRepository.CreateUser(db, user) == false {
		return nil, errors.New("用户创建失败")
	}

	return user, nil
}

func (userLogic UserLogic) ResetUserCache(ctx context.Context, user *model.User) error {
	cacheKey := fmt.Sprintf(USER_INFO_CACHE_KEY, user.ID)
	data := userLogic.GetDefaultRedis().Del(ctx, cacheKey)
	if data.Err() != nil && data.Err() != redis.Nil {
		return data.Err()
	}

	return nil
}

func (userLogic UserLogic) GetUserById(ctx context.Context, userId model.UID) (*model.User, error) {
	cacheKey := fmt.Sprintf(USER_INFO_CACHE_KEY, userId)
	data := userLogic.GetDefaultRedis().Get(ctx, cacheKey)
	if data.Err() != nil && data.Err() != redis.Nil {
		return nil, data.Err()
	}

	if data.Val() == "" {
		user := repository.Repository.UserRepository.GetById(userLogic.GetDefaultDb(), userId)
		if user == nil {
			return nil, errors.New("用户不存在")
		}

		encodeData, err := helper.JsonEncode(user)
		if err != nil {
			return nil, err
		}
		result := userLogic.GetDefaultRedis().Set(ctx, cacheKey, encodeData, 0)
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

func (userLogic UserLogic) GetByMobileAndPwd(mobile string, password string) (*model.User, error) {
	user := repository.Repository.UserRepository.GetByMobile(userLogic.GetDefaultDb(), mobile)
	if user == nil {
		return nil, errors.New("该手机号不存在")
	}
	if user.MakeHashPassword(password, user.Salt) != user.Password {
		return nil, errors.New("手机号或者密码错误")
	}

	return user, nil
}

func (userLogic UserLogic) UpdateUser(user *model.User) error {
	if !repository.Repository.UserRepository.UpdateUser(userLogic.GetDefaultDb(), user) {
		return errors.New("更新用户信息失败")
	}

	return nil
}
