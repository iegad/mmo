package basic

import (
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/iegad/kraken/security"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao/archive_log"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func ModBasic(obj *ds.Basic, db *sql.DB, es *elastic.Client) error {
	utils.Assert(obj != nil && &obj.Entry != nil && db != nil && es != nil, "ModBasic in params is invalid")

	var (
		err error
		raw *ds.Basic
	)

	for i := 0; i < 3; i++ {
		raw, err = GetBasicByID(obj.Entry.UserID, db)
		if err != nil {
			return err
		}

		err = updateBasic(obj, raw, db, es)
		if err == nil {
			break
		}
	}

	if err != nil {
		return err
	}

	rawData := utils.PbToBytes(raw)
	objData := utils.PbToBytes(raw)

	// 当数据有改变时才需要记录修改日志
	if hex.EncodeToString(security.MD5(rawData)) != hex.EncodeToString(security.MD5(objData)) {
		err = archive_log.AddArchiveLog(&ds.ArchiveLog{
			Type:        ds.ArchiveLogType_ArchiveLogType_Basic,
			VerCode:     obj.VerCode,
			UserID:      raw.Entry.UserID,
			ArchiveTime: time.Now().Unix(),
			Raw:         rawData,
			Changed:     objData,
		}, es)
	}

	return err
}

func updateBasic(obj, raw *user.Basic, db *sql.DB, es *elastic.Client) error {
	n := 0

	if len(obj.Entry.Email) == 0 {
		obj.Entry.Email = raw.Entry.Email
		n++
	}

	if len(obj.Entry.PhoneNum) == 0 {
		obj.Entry.PhoneNum = raw.Entry.PhoneNum
		n++
	}

	if len(obj.Entry.Nickname) == 0 {
		obj.Entry.Nickname = raw.Entry.Nickname
		n++
	}

	if obj.Entry.Gender == 0 {
		obj.Entry.Gender = raw.Entry.Gender
		n++
	}

	if len(obj.Entry.Avator) == 0 {
		obj.Entry.Avator = raw.Entry.Avator
		n++
	}

	// 表示没有任何数据需要修改
	if n == 5 {
		return nil
	}

	obj.LastUpdate = time.Now().Unix()
	obj.CreateTime = raw.CreateTime
	obj.VerCode = raw.VerCode + 1

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE `DB_USER`.`T_BASIC` SET F_EMAIL=?,F_PHONE_NUM=?,F_GENDER=?,F_NICKNAME=?,F_AVATOR=?,F_VER_CODE=? WHERE F_USER_ID=? AND F_VER_CODE=?",
		utils.IF_STR_EMPTY(obj.Entry.Email), utils.IF_STR_EMPTY(obj.Entry.PhoneNum), obj.Entry.Gender, obj.Entry.Nickname, obj.Entry.Avator, obj.VerCode, obj.Entry.UserID, raw.VerCode)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = entry.ModEntry(obj.Entry, es)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
