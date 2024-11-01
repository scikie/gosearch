package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gosearch/utils"
	"os"
)

var DB *sqlx.DB

// 连接到MySQL数据库,初始化。
func InitDb() {
	/* =============  创建数据库连接   =============*/

	/* 1.==> windows 上调试的使用 */
	// 定义一个数据源名称（Data Source Name），用于连接到MySQL数据库。
	// 这里包含了数据库的用户名(root)、密码(1234)、地址(127.0.0.1:3306)、数据库名(sensor_db)以及一些连接参数。
	// 默认：
	// dsn := "root:1234@tcp(127.0.0.1:3306)/sensor_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 使用配置文件：
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	// 使用sqlx.Connect函数连接到MySQL数据库。这里传递了数据库驱动名("mysql")和数据源名称(dsn)。
	// 如果连接成功，DB变量将保存数据库连接对象；如果连接失败，err变量将包含错误信息。
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		// 如果连接失败，使用fmt.Printf输出错误信息并返回。
		logrus.Fatalf("连接数据库失败, err:%v\n", err)
		os.Exit(1)

	}

	// 设置数据库连接池的最大打开连接数，这里设置为20。
	DB.SetMaxOpenConns(20)

	// 设置数据库连接池的最大空闲连接数，这里也设置为10。
	DB.SetMaxIdleConns(10)

	// 创建users表的SQL语句
	createTableSQL := `
      CREATE TABLE IF NOT EXISTS users (
		user_id INT NOT NULL AUTO_INCREMENT COMMENT '用户ID，主键',
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		phone_number VARCHAR(20) NOT NULL,
		create_time DATETIME NOT NULL DEFAULT NOW(),
		update_time DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
		delete_time datetime DEFAULT NULL,
		status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（0：异常，1：正常）,
		PRIMARY KEY (user_id),
		KEY idx_delete_time (delete_time)
	 ) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
    `
	// 执行创建表操作
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		logrus.Fatalf("创建users表失败, err:%v\n", err)
		os.Exit(1)
	}
	logrus.Infoln("users表创建成功")


	// 创建albums表的SQL语句
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS albums (
			album_id INT NOT NULL AUTO_INCREMENT COMMENT '相册的唯一标识，主键',
			user_id INT NOT NULL COMMENT '创建相册的用户ID，外键，关联用户表',
			album_name VARCHAR(255) NOT NULL COMMENT '相册名称',
			album_description TEXT COMMENT '相册描述',
			create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '相册创建时间',
			update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '相册最后更新时间',
			status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（0：异常，1：正常）',
			PRIMARY KEY (album_id),
			INDEX idx_user (user_id),
			CONSTRAINT fk_user_album FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

    `
	// 执行创建表操作
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		logrus.Fatalf("创建albums表失败, err:%v\n", err)
		os.Exit(1)
	}
	logrus.Infoln("albums表创建成功")

	// 创建photos表的SQL语句
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS photos (
			id INT NOT NULL AUTO_INCREMENT COMMENT '图片ID，主键',
			album_id INT NOT NULL COMMENT '所属相册ID',
			user_id INT NOT NULL COMMENT '上传用户ID',
			photo_name VARCHAR(511) NOT NULL COMMENT '图片名称',
			photo_description TEXT COMMENT '图片描述',
			photo_path VARCHAR(511) NOT NULL COMMENT '图片存储路径',
			photo_thumb_path VARCHAR(511) NOT NULL COMMENT '图片缩略图存储路径',
			photo_size INT NOT NULL DEFAULT 0	COMMENT '图片大小（字节）',
			photo_width INT NOT NULL DEFAULT 0 COMMENT '图片宽度（像素）',
			photo_height INT NOT NULL DEFAULT 0 COMMENT '图片高度（像素）',
			upload_time DATETIME NOT NULL DEFAULT NOW() COMMENT '上传时间',
			update_time DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW() COMMENT '更新时间',
			is_public TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否公开（0：不公开，1：公开）',
			status TINYINT(1) NOT NULL DEFAULT 1 COMMENT '图片状态（0：已删除，1：正常）',
			PRIMARY KEY (id),
			CONSTRAINT fk_user_photo FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
			CONSTRAINT fk_album_photo FOREIGN KEY (album_id) REFERENCES albums (album_id) ON DELETE CASCADE,
			INDEX idx_album_id (album_id),
			INDEX idx_user_id (user_id),
			FULLTEXT idx_fulltext_description (photo_description)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图片表';
    `
	// 执行创建表操作
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		logrus.Fatalf("创建photos表失败, err:%v\n", err)
		os.Exit(1)
	}
	logrus.Infoln("photos表创建成功")
}