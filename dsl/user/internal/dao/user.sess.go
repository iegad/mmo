package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/data"
)

/**
 * 需要作到单点登录.
 * 如果用户已经在A设备登录成功, 此时B设备也发出了登录请求.
 * B设备会被通知用户已登录.
 * 此时B设备第二次尝试登录, 则允许B设备登录, 踢下A设备登录状态.
 *
 * 实现方式:
 *  5 分钟之内两次登录 并且DeviceInfo不一至则允许登录
 *  ActiveTime > time.Now().Add(-time.Minutes * 5) && DeviceInfo <> current.DeviceInfo
 */

const (
	USER_SESS_TTL = time.Hour * 24 * 7
)

func userSessKey(userID int64) string {
	return fmt.Sprintf("UserSess_%d", userID)
}

func SetUserSess(ss *data.UserSession, rconn *redis.Client) error {
	if len(ss.Token) != 16 {
		return errors.New("token is invalid")
	}

	if len(ss.MountAddr) >= 36 {
		return errors.New("mountAddr is invalid")
	}

	if len(ss.DeviceCode) != 16 {
		return errors.New("deviceCode is invalid")
	}

	ss.ActiveTime = time.Now().Unix()
	data, err := json.Marshal(ss)
	if err != nil {
		return err
	}
	return rconn.Set(context.TODO(), userSessKey(ss.UserInfo.UserID), utils.BytesToString(data), USER_SESS_TTL).Err()
}

func GetUserSess(userID int64, rconn *redis.Client) (*data.UserSession, error) {
	if userID <= 0 {
		return nil, errors.New("userID is invalid")
	}

	jd, err := rconn.Get(context.TODO(), userSessKey(userID)).Result()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ss := data.NewUserSession()
	err = json.Unmarshal(utils.StringToBytes(jd), ss)
	if err != nil {
		return nil, err
	}

	return ss, nil
}

func RemoveUserSess(userID int64, rconn *redis.Client) error {
	return rconn.Del(context.TODO(), userSessKey(userID)).Err()
}
