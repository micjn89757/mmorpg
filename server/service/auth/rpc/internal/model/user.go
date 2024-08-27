package db

import (
	"server/common/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var all_columns = []string{"id", "account", "password"}

type User struct {
	Id       uint   `gorm:"column:id;primaryKey"`
	Account  string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string { // 显示指定表名
	return "user"
}

// 根据用户名检索用户
func GetUserByAccount(account string) *User {
	db := GetMMODBConnection()
	var user User
	if err := db.Select(all_columns).Where("account=?", account).First(&user).Error; err != nil { //name不存在重复，所以用First即可
		if err != gorm.ErrRecordNotFound { // 如果是用户名不存在，不需要打错误日志
			logger.Logger.Error("get password of user failed:", zap.String("account", account), zap.String("err", err.Error())) // 系统性异常，才打错误日志
		}
		return nil
	}
	return &user
}


// 创建一个用户
func CreateAccount(account, pass string) (uint, error) {
	db := GetMMODBConnection()
	// pass = edcrypt.Md5(pass)                          //前端密码要经过md5
	user := User{Account: account, Password: pass}         //ORM
	if err := db.Create(&user).Error; err != nil { //必须传指针，因为要给user的主键赋值
		logger.Logger.Error("create user failed", zap.String("account", account), zap.String("err", err.Error()))
		return 0, err
	} else {
		logger.Logger.Info("create user id", zap.Uint("id", user.Id))
		return user.Id, nil
	}
}