package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	exception "github.com/titrxw/smart-home-server/app/Exception"

	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

const USER_INFO_CACHE_KEY = `user:info:%d`

type UserLogic struct {
	LogicAbstract
}

func (userLogic UserLogic) CreateUser(userName string, email string, password string) (*model.User, error) {
	db := userLogic.GetDefaultDb()
	userRepository := repository.Repository.UserRepository

	user := userRepository.GetByUserName(db, userName)
	if user != nil {
		return nil, exception.NewLogicError("该用户名已存在")
	}
	user = userRepository.GetByEmail(db, email)
	if user != nil {
		return nil, exception.NewLogicError("该email已存在")
	}

	user = &model.User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
	if userRepository.CreateUser(db, user) == false {
		return nil, exception.NewLogicError("用户创建失败")
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
			return nil, exception.NewLogicError("用户不存在")
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

func (userLogic UserLogic) GetByEmail(email string) *model.User {
	return repository.Repository.UserRepository.GetByEmail(userLogic.GetDefaultDb(), email)
}

func (userLogic UserLogic) GetByEmailAndPwd(email string, password string) (*model.User, error) {
	user := userLogic.GetByEmail(email)
	if user == nil {
		return nil, exception.NewLogicError("该email不存在")
	}
	if user.MakeHashPassword(password, user.Salt) != user.Password {
		return nil, exception.NewLogicError("email或者密码错误")
	}

	return user, nil
}

func (userLogic UserLogic) UpdateUser(user *model.User) error {
	if !repository.Repository.UserRepository.UpdateUser(userLogic.GetDefaultDb(), user) {
		return exception.NewLogicError("更新用户信息失败")
	}

	return nil
}
