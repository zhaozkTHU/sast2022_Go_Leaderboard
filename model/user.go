package model

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"not null;autoIncrement"`    //用户id
	UserName string `gorm:"type:varchar(255);unique;"` //用户名字
	Votes    uint   `gorm:"default:0"`                 //用户所得票数
}

func CreateUser(name string) (error, uint) {
	user := User{
		UserName: name,
		Votes:    0,
	}
	tx := DB.Create(&user)
	return tx.Error, user.ID
}

func GetUserByName(name string) (error, User) {
	user := User{}
	tx := DB.Model(&User{}).Where("user_name=?", name).First(&user)
	return tx.Error, user
}

func AddVoteForUser(name string) error {
	err, user := GetUserByName(name)
	if err != nil {
		return err
	}
	tx := DB.Model(&User{}).Where("id = ?", user.ID).Update("votes", gorm.Expr("votes+ ?", 1))
	return tx.Error
}
