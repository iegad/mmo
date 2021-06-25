package dao

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/iegad/hydra/mod"
	"github.com/iegad/hydra/pb"
	"github.com/iegad/kraken/security"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/data"
)

var (
	tUserInfoKey []byte
)

func init() {
	tUserInfoKey = security.MD5(utils.StringToBytes(data.TUserInfo))
}

// AddUserInfo 添加用户
func AddUserInfo(obj *data.UserInfo, db *sql.DB) error {
	if obj == nil {
		return ErrObj
	}

	if db == nil {
		return ErrDB
	}

	if len(obj.Email) == 0 && len(obj.PhoneNum) == 0 {
		return ErrAccount
	}

	if obj.CreateTime == 0 {
		obj.CreateTime = time.Now().Unix()
	}

	res, err := db.Exec("INSERT INTO `DB_BASIC`.`T_USER_INFO`(F_EMAIL,F_PHONE_NUM,F_CREATE_TIME,F_VER) VALUES (?,?,?,'')",
		*utils.IF_STR_EMPTY(obj.Email), *utils.IF_STR_EMPTY(obj.PhoneNum), obj.CreateTime)
	if err != nil {
		return err
	}

	obj.UserID, err = res.LastInsertId()
	return err
}

// ModifyUserInfo: 修改用户
func ModifyUserInfo(obj *data.UserInfo, db *sql.DB) error {
	// Step 1: 入参检查
	if obj == nil {
		return ErrObj
	}

	if db == nil {
		return ErrDB
	}

	if obj.UserID <= 0 {
		return ErrUserID
	}

	var (
		changed         bool
		vcode           uint32
		affected        int64
		err             error
		encode          []byte
		email, phoneNum sql.NullString
		tmp             = data.NewUserInfo()
		row             = db.QueryRow("SELECT F_AID,F_EMAIL,F_PHONE_NUM,F_CREATE_TIME,F_LAST_UPDATE,F_VER_CODE FROM `DB_BASIC`.`T_USER_INFO` WHERE F_AID=?", obj.UserID)
	)

	for dwf := true; dwf; dwf = false {
		for i := 0; i < 3; i++ {
			// Step 2: 获取原始版本
			err = row.Scan(&tmp.UserID, &email, &phoneNum, &tmp.CreateTime, &tmp.LastUpdate, &vcode)
			if err != nil {
				break
			}

			if email.Valid {
				tmp.Email = email.String
			}

			if phoneNum.Valid {
				tmp.PhoneNum = phoneNum.String
			}

			// Step 3: 备份原始版本
			encode, err = security.AES128Encode(pb.ToBytes(tmp), tUserInfoKey)
			if err != nil {
				break
			}

			tmp.Ver = hex.EncodeToString(encode)
			tmp.VerCode = vcode + 1

			// Step 4: 赋值
			if len(obj.Email) > 0 && obj.Email != tmp.Email {
				tmp.Email = obj.Email
				changed = true
			}

			if len(obj.Avator) > 0 && obj.Avator != tmp.Avator {
				tmp.Avator = obj.Avator
				changed = true
			}

			if len(obj.PhoneNum) > 0 && obj.PhoneNum != tmp.PhoneNum {
				tmp.PhoneNum = obj.PhoneNum
				changed = true
			}

			if !changed {
				break
			}

			tmp.LastUpdate = time.Now().Unix()

			// Step 5: 更新DB
			res, err := db.Exec("UPDATE `DB_BASIC`.`T_USER_INFO` SET F_EMAIL=?,F_PHONE_NUM=?,F_NICKNAME=?,F_AVATOR=?,F_LAST_UPDATE=?,F_VER=?,F_VER_CODE=? WHERE F_AID=? AND F_VER_CODE=?",
				utils.IF_STR_EMPTY(tmp.Email), utils.IF_STR_EMPTY(tmp.PhoneNum), tmp.Nickname, tmp.Avator, tmp.LastUpdate, tmp.Ver, tmp.VerCode, tmp.UserID, vcode)
			if err != nil {
				return err
			}

			affected, err = res.RowsAffected()
			if err != nil {
				break
			}

			if affected == 1 {
				data.DeleteUserInfo(obj)
				obj = tmp
				tmp = nil
				break
			}

			time.Sleep(time.Millisecond * 100)
		}
	}

	// Step 6: 清理
	data.DeleteUserInfo(tmp)

	if changed && affected != 1 {
		if err == nil {
			err = mod.ErrTimeout
		}
	}

	return err
}

func RemoveUserInfo(userID int64, db *sql.DB) error {
	if userID <= 0 {
		return ErrUserID
	}

	if db == nil {
		return mod.ErrDB
	}

	_, err := db.Exec("DELETE FROM `DB_BASIC`.`T_USER_INFO` WHERE F_AID=?", userID)
	if err != nil {
		return err
	}

	return nil
}

func QueryOneUserInfo(userID int64, db *sql.DB) (*data.UserInfo, error) {
	if userID <= 0 {
		return nil, ErrUserID
	}

	if db == nil {
		return nil, mod.ErrDB
	}

	var (
		tmp             = data.NewUserInfo()
		email, phoneNum sql.NullString
	)

	row := db.QueryRow("SELECT F_AID,F_EMAIL,F_PHONE_NUM,F_CREATE_TIME,F_LAST_UPDATE,F_VER,F_VER_CODE FROM `DB_BASIC`.`T_USER_INFO` WHERE F_AID=?", userID)
	err := row.Scan(&tmp.UserID, &email, &phoneNum, &tmp.CreateTime, &tmp.LastUpdate, &tmp.Ver, &tmp.VerCode)
	if err != nil {
		data.DeleteUserInfo(tmp)
		return nil, err
	}

	if email.Valid {
		tmp.Email = email.String
	}

	if phoneNum.Valid {
		tmp.PhoneNum = phoneNum.String
	}

	return tmp, nil
}

func QueryUserInfo(cond *cgi.QueryUserInfoReq, db *sql.DB) ([]*data.UserInfo, error) {
	if db == nil {
		return nil, mod.ErrDB
	}

	sb := &strings.Builder{}
	sb.WriteString("SELECT F_AID,F_EMAIL,F_PHONE_NUM,F_CREATE_TIME,F_LAST_UPDATE,F_VER,F_VER_CODE FROM `DB_BASIC`.`T_USER_INFO`")

	if cond != nil {
		wa := false
		if cond.UserID > 0 {
			sb.WriteString(fmt.Sprintf(" WHERE F_AID=%d", cond.UserID))
			wa = true
		}

		if len(cond.Email) > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_EMAIL='%s'", cond.Email))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_EMAIL='%s'", cond.Email))
				wa = true
			}
		}

		if len(cond.PhoneNum) > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_PHONE_NUM='%s'", cond.PhoneNum))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_PHONE_NUM='%s'", cond.PhoneNum))
				wa = true
			}
		}

		if cond.CreateTimeBeg > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME>=%d", cond.CreateTimeBeg))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_CREATE_TIME>=%d", cond.CreateTimeBeg))
				wa = true
			}
		}

		if cond.CreateTimeEnd > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME<%d", cond.CreateTimeEnd))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_CREATE_TIME<%d", cond.CreateTimeEnd))
				wa = true
			}
		}

		if cond.LastUpdateBeg > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE>=%d", cond.LastUpdateBeg))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_LAST_UPDATE>=%d", cond.LastUpdateBeg))
				wa = true
			}
		}

		if cond.LastUpdateEnd > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE>=%d", cond.LastUpdateEnd))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_LAST_UPDATE>=%d", cond.LastUpdateEnd))
				wa = true
			}
		}

		if len(cond.OrderBy) > 0 {
			sb.WriteString(fmt.Sprintf(" ORDER BY %s", data.TUserInfoFieldMap[cond.OrderBy]))
			if cond.Desc {
				sb.WriteString(" DESC")
			}
		}

		if cond.Limit > 0 {
			sb.WriteString(fmt.Sprintf(" LIMIT %d", cond.Limit))
			if cond.Offset > 0 {
				sb.WriteString(fmt.Sprintf(" OFFSET %d", cond.Offset))
			}
		}
	}

	rows, err := db.Query(sb.String())
	if err != nil {
		return nil, err
	}

	var (
		dataList = []*data.UserInfo{}
		email    sql.NullString
		phoneNum sql.NullString
	)

	for rows.Next() {
		item := data.NewUserInfo()
		err = rows.Scan(&item.UserID, &email, &phoneNum, &item.CreateTime, &item.LastUpdate, &item.Ver, &item.VerCode)
		if err != nil {
			break
		}

		if email.Valid {
			item.Email = email.String
		}

		if phoneNum.Valid {
			item.PhoneNum = phoneNum.String
		}

		dataList = append(dataList, item)
	}

	rows.Close()

	if err != nil {
		for i := 0; i < len(dataList); i++ {
			data.DeleteUserInfo(dataList[i])
			dataList[i] = nil
		}

		return nil, err
	}

	return dataList, nil
}

func ReleaseUserInfoDataList(dataList []*data.UserInfo) {
	for i := 0; i < len(dataList); i++ {
		data.DeleteUserInfo(dataList[i])
		dataList[i] = nil
	}
}

func CountUserInfo(cond *cgi.QueryUserInfoReq, db *sql.DB) (int64, error) {
	if db == nil {
		return -1, mod.ErrDB
	}

	sb := &strings.Builder{}
	sb.WriteString("SELECT COUNT(1) FROM `DB_BASIC`.`T_USER_INFO`")

	if cond != nil {
		wa := false
		if cond.UserID > 0 {
			sb.WriteString(fmt.Sprintf(" WHERE F_AID=%d", cond.UserID))
			wa = true
		}

		if len(cond.Email) > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_EMAIL='%s'", cond.Email))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_EMAIL='%s'", cond.Email))
				wa = true
			}
		}

		if len(cond.PhoneNum) > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_PHONE_NUM='%s'", cond.PhoneNum))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_PHONE_NUM='%s'", cond.PhoneNum))
				wa = true
			}
		}

		if cond.CreateTimeBeg > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME>=%d", cond.CreateTimeBeg))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_CREATE_TIME>=%d", cond.CreateTimeBeg))
				wa = true
			}
		}

		if cond.CreateTimeEnd > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME<%d", cond.CreateTimeEnd))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_CREATE_TIME<%d", cond.CreateTimeEnd))
				wa = true
			}
		}

		if cond.LastUpdateBeg > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE>=%d", cond.LastUpdateBeg))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_LAST_UPDATE>=%d", cond.LastUpdateBeg))
				wa = true
			}
		}

		if cond.LastUpdateEnd > 0 {
			if wa {
				sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE>=%d", cond.LastUpdateEnd))
			} else {
				sb.WriteString(fmt.Sprintf(" WHERE F_LAST_UPDATE>=%d", cond.LastUpdateEnd))
				wa = true
			}
		}
	}

	var (
		total int64
		row   = db.QueryRow(sb.String())
	)

	err := row.Scan(&total)
	if err != nil {
		return -1, err
	}

	return total, nil
}
