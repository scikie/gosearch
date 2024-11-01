package model

import (
	"time"
)

// Photo 定义照片的结构体
type Photo struct {
	ID               int       `db:"id" json:"id" form:"id"`
	AlbumID          int       `db:"album_id" json:"album_id" binding:"required" form:"album_id"`
	UserID           int       `db:"user_id" json:"user_id" binding:"required" form:"user_id"`
	PhotoName        string    `db:"photo_name" json:"photo_name" binding:"required" form:"photo_name"`
	PhotoDescription string    `db:"photo_description" json:"photo_description" form:"photo_description"`
	PhotoPath        string    `db:"photo_path" json:"photo_path" binding:"required" form:"photo_path"`
	PhotoThumbPath   string    `db:"photo_thumb_path" json:"photo_thumb_path" binding:"required" form:"photo_thumb_path"`
	PhotoSize        int       `db:"photo_size" json:"photo_size" form:"photo_size"`
	PhotoWidth       int       `db:"photo_width" json:"photo_width" form:"photo_width"`
	PhotoHeight      int       `db:"photo_height" json:"photo_height" form:"photo_height"`
	UploadTime       time.Time `db:"upload_time" json:"upload_time" form:"upload_time"`
	UpdateTime       time.Time `db:"update_time" json:"update_time" form:"update_time"`
	IsPublic         bool      `db:"is_public" json:"is_public" form:"is_public"`
	Status           bool      `db:"status" json:"status" form:"status"`
}

// Photo 定义照片的结构体
type RespPhoto struct {
	PhotoName        string    `db:"photo_name" json:"photo_name" binding:"required" form:"photo_name"`
	PhotoDescription string    `db:"photo_description" json:"photo_description" form:"photo_description"`
	PhotoPath        string    `db:"photo_path" json:"photo_path" binding:"required" form:"photo_path"`
	PhotoSize        int       `db:"photo_size" json:"photo_size,omitempty" binding:"required" form:"photo_size"`
	PhotoWidth       int       `db:"photo_width" json:"photo_width,omitempty" binding:"required" form:"photo_width"`
	PhotoHeight      int       `db:"photo_height" json:"photo_height,omitempty" binding:"required" form:"photo_height"`
}

func InsertPhoto(photo Photo) error {
    // 定义插入SQL语句，使用命名参数
    insertSQL := `
        INSERT INTO photos (id, album_id, user_id, photo_name, photo_description, photo_path, 
                            photo_thumb_path, photo_size, photo_width, photo_height, upload_time, 
                            update_time, is_public, status)
        VALUES (:id, :album_id, :user_id, :photo_name, :photo_description, :photo_path, 
                :photo_thumb_path, :photo_size, :photo_width, :photo_height, :upload_time, 
                :update_time, :is_public, :status)
    `
    // 使用NamedExec执行插入操作
    _, err := DB.NamedExec(insertSQL, photo)
    return err
}

func SelectPhoto(search string) ([]RespPhoto,error) {
    // 定义插入SQL语句，使用命名参数
	//  MATCH(photo_description) AGAINST('+关键词' IN BOOLEAN MODE)
    selectSQL := `
		SELECT photo_name,photo_description,photo_path,photo_size,photo_width,photo_height
			FROM photos
		WHERE MATCH(photo_description) AGAINST('+?' IN BOOLEAN MODE);
    `
	var photos []RespPhoto
    // 使用NamedExec执行插入操作
    err := DB.Select(&photos,selectSQL,search)
    return photos,err
}



