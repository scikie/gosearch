package model

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

// 定义用户表结构
type User struct {
	UserID      int            `json:"user_id,omitempty" db:"user_id" form:"user_id"`
	Username    string         `json:"username" db:"username" binding:"required" form:"username"`
	Password    string         `json:"password" db:"password" binding:"required" form:"password"`
	Email       string         `json:"email" db:"email" binding:"required" form:"email"`
	PhoneNumber string         `json:"phone_number" db:"phone_number" form:"phone_number"`
	CreateTime  time.Time      `json:"create_time,omitempty" db:"create_time" form:"create_time"`
	UpdateTime  time.Time      `json:"update_time,omitempty" db:"update_time" form:"update_time"`
	DeleteTime  DeleteNullTime `json:"delete_time,omitempty" db:"delete_time" form:"delete_time"`
	Status      uint8          `json:"status,omitempty" db:"status" form:"status"`
}

// 通过名字查询用户名称是否已经存在
func CheckUser(name string) (code int, err error) {
	var user User
	sqlStr := "select * from users where username=?"
	err = DB.Get(&user, sqlStr, name)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到记录，返回成功状态码，表示该用户名未被使用
			logrus.Debugf("[%s]用户名未被使用\n", name)
			return 200, nil
		} else {
			// 真正的数据库错误
			logrus.Errorf("func=model.CheckUser():%v\n", err)
			return 500, err
		}
	}
	// 如果查询成功，表示该用户名已经被使用，返回已经使用的状态码
	if user.UserID > 0 {
		return 1001, err //1001
	}
	return 200, err
}

// CreateUser 新增用户
func CreateUser(data *User) (code int, err error) {
	data.CreateTime=time.Now()
	logrus.Debugf("新增user:%v\n", data)
	query := `INSERT INTO users (
	user_id,username, password,email,phone_number,create_time,update_time,delete_time,status
	)
	VALUES (
	:user_id,:username,:password,:email,:phone_number,:create_time,:update_time,:delete_time,:status
	)`
	result, err := DB.NamedExec(query, data)
	if err != nil {
		logrus.Errorln("func=model.CreateUser.DB.NamedExec():", err)
	}

	// 获取插入后的ID
	_, err = result.LastInsertId()
	if err != nil {
		logrus.Errorln("func=model.CreateUser.result.LastInsertId():", err)
		return 500, err // 500
	}
	return 200, err
}

// GetUser 查询用户
func GetUser(id int) (user User, code int, err error) {
	sqlStr := `SELECT * from users where user_id = ?`
	err = DB.Get(&user, sqlStr, id)
	if err != nil {
		logrus.Errorf("func=model.GetUser.DB.Get():%v\n", err)
		return user, 500, err
	}
	return user, 200, err
}

// GetUser 查询用户
func GetUserByName(username string) (user User, code int, err error) {
	sqlStr := `SELECT * from users where username = ?`
	err = DB.Get(&user, sqlStr, username)
	if err != nil {
		if err==sql.ErrNoRows{
			logrus.Errorf("no rows in result set\n")
		return user, 1001, err
		}else{
			logrus.Errorf("func=model.GetUserByName.DB.Get():%v\n", err)
			return user, 500, err
		}
		
	}
	return user, 200, err
}

// 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) (users []User, total,code int, err error) {
	total=0
	if username != "" {
		sqlStr := `SELECT * FROM users where username LIKE ? LIMIT ? OFFSET ?;`
		err = DB.Select(&users, sqlStr, username, pageSize, (pageNum-1)*pageSize)
		if err != nil {
			logrus.Errorf("func=model.GetUsers.1.DB.Select():%v\n", err)
			return users, total,500,err
		}
		sqlStr = `SELECT COUNt(*) FROM users where username LIKE ?;`
		err = DB.Get(&total, sqlStr,username)
		if err != nil {
			logrus.Errorf("func=model.GetUsers.2.DB.Select():%v\n", err)
			return users, total,500,err
		}
		return users,total, 200, nil
		
	} else {
		sqlStr := `SELECT * FROM users LIMIT ? OFFSET ?;`
		err = DB.Select(&users, sqlStr, pageSize, (pageNum-1)*pageSize)
		if err != nil {
			logrus.Errorf("func=model.GetUsers.3.DB.Select():%v\n", err)
			return users,total, 500, err
		}
		sqlStr = `SELECT COUNt(*) FROM users;`
		err = DB.Get(&total, sqlStr)
		if err != nil {
			logrus.Errorf("func=model.GetUsers.4.DB.Select():%v\n", err)
			return users, total,500,err
		}
		return users,total, 200, nil
	}
	

}
