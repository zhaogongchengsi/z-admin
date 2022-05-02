package global

import (
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	ConfigName string = "config"    // 配置文件名字
	ConfigType string = "yaml"      // 配置文件类型
	ConfigPath string = "./configs" // 配置文件路径
)

var (
	Server      *Service
	Db          *DataBase
	DBEngine    *gorm.DB
	CaptchaConf *Captcha
	// 这里可以用radis缓存
	VerifyStore = base64Captcha.DefaultMemStore
)

func InitGlobal() error {
	s, err := NewSetting(ConfigName, ConfigPath, ConfigType)
	if err != nil {
		return err
	}

	err = s.ReadSetting("Service", &Server)

	if err != nil {
		return err
	}

	err = s.ReadSetting("DateBase", &Db)

	if err != nil {
		return err
	}

	err = s.ReadSetting("Captcha", &CaptchaConf)

	if err != nil {
		fmt.Printf("读取验证码配置失败 %v", err)
		return err
	}

	err = ConnectDb()

	if err != nil {
		return err
	}

	return nil
}

func ConnectDb() error {

	dns := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local`,
		Db.UserName,
		Db.PassWord,
		Db.Url,
		Db.Port,
		Db.DbName,
		Db.Charset,
	)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Db.TablePrefix,   // 表名前缀，
			SingularTable: Db.SingularTable, // 使用单数表名，启用该选项后，`User` 表将是`user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 不使用物理外键
	})

	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()

	if err != nil {
		return err
	}

	DBEngine = DB

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(Db.MaxIdleConns)

	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(Db.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
