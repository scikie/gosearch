package model

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

// 定义albums表结构
type Album struct {
	AlbumID          int       `json:"album_id,omitempty" db:"album_id" form:"album_id"`
	UserID           int       `json:"user_id,omitempty" db:"user_id" form:"user_id"`
	Albumname        string    `json:"albumname" db:"albumname" binding:"required" form:"albumname"`
	AlbumDescription string    `json:"album_description,omitempty" db:"album_description" form:"album_description"`
	CreateTime       time.Time `json:"create_time,omitempty" db:"create_time" form:"create_time"`
	UpdateTime       time.Time `json:"update_time,omitempty" db:"update_time" form:"update_time"`
	Status           uint8     `json:"status,omitempty" db:"status" form:"status"`
}

// 通过名字查询album名称是否已经存在
func CheckAlbum(name string) (code int, err error) {
	var data Album
	sqlStr := "select * from albums where albumname=?"
	err = DB.Get(&data, sqlStr, name)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到记录，返回成功状态码，表示该相册名未被使用
			logrus.Debugf("[%s]相册名未被使用\n", name)
			return 200, nil
		} else {
			// 真正的数据库错误
			logrus.Errorf("func=model.CheckAlbum():%v\n", err)
			return 500, err
		}
	}
	// 如果查询成功，表示该相册名已经被使用，返回已经使用的状态码
	if data.AlbumID > 0 {
		return 1001, err //1001
	}
	return 200, err
}

// CreateUser 新增相册
func CreateAlbum(data *Album) (code int, err error) {
	data.CreateTime=time.Now()
	logrus.Debugf("新增album:%v\n", data)
	query := `INSERT INTO albums (
	album_id,user_id,albumname,album_description,create_time,update_time,status
	)
	VALUES (
	:album_id,:user_id,:albumname,:album_description,:create_time,:update_time,:status
	)`
	result, err := DB.NamedExec(query, data)
	if err != nil {
		logrus.Errorln("func=model.CreateAlbum.DB.NamedExec():", err)
	}

	// 获取插入后的ID
	_, err = result.LastInsertId()
	if err != nil {
		logrus.Errorln("func=model.CreateAlbum.result.LastInsertId():", err)
		return 500, err // 500
	}
	return 200, err
}


func GetAlbum(id int) (data Album, code int, err error) {
	sqlStr := `SELECT * from albums where album_id = ?`
	err = DB.Get(&data, sqlStr, id)
	if err != nil {
		logrus.Errorf("func=model.GetAlbum.DB.Get():%v\n", err)
		return data, 500, err
	}
	return data, 200, err
}


func GetAlbumByUser(userid int) (album Album, code int, err error) {
	sqlStr := `SELECT * from albums where user_id = ?`
	err = DB.Get(&album, sqlStr, userid)
	if err != nil {
		if err==sql.ErrNoRows{
			logrus.Errorf("no rows in result set\n")
		return album, 1001, err
		}else{
			logrus.Errorf("func=model.GetAlbumByUser.DB.Get():%v\n", err)
			return album, 500, err
		}
		
	}
	return album, 200, err
}

