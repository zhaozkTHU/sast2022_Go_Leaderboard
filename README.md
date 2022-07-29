[原代码框架](https://github.com/pyz-creeper/SAST-2022-Go-LeaderBoard)

几点疑问：

1.为什么说很多在用go重构django，优点在哪里

2.

``` go
type Config struct {
	DbUserName string `json:"db_user_name"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	DbIP       string `json:"db_ip"`
}

type Submission struct {
	ID        uint   `gorm:"not null;autoIncrement"`
	UserName  string `gorm:"type:varchar(255);"`
	Avatar    string //头像base64，也可以是一个头像链接
	CreatedAt int64  //提交时间
	Score     int    //评测成绩
	Subscore1 int    //评测小分
	Subscore2 int    //评测小分
	Subscore3 int    //评测小分
}
```

在代码的很多地方都可以见到类似的struct定义方法，每一项最后的用`''`的内容具体什么意思，如何使用