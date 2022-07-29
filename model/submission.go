package model

import (
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

// Submission hint: 如果你想直接返回结构体，可以考虑在这里加上`json`的tag
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

// ReturnSub 这里提供返回的submission的示例结构
type ReturnSub struct {
	UserName  string `json:"user"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"time"`
	Score     int    `json:"score"`
	UserVotes int    `json:"votes"`
	Subscore1 int
	Subscore2 int
	Subscore3 int
}

func Judge(content string) ([4]int, error) {
	//处理答案
	f, _ := ioutil.ReadFile("ground_truth.txt")
	groundTruth := string(f)
	answerStrings := strings.Split(groundTruth, "\n") //行分割
	var answerDict map[string][3]string
	for i := 1; i < len(answerStrings); i++ {
		a := strings.Split(answerStrings[i], ",")
		answerDict[a[0]] = [3]string{a[1], a[2], a[3]}
	}

	//处理提交结果
	submit := strings.Split(content, "\n")
	var submit_dict map[string][3]string
	for i := 1; i < len(submit); i++ {
		a := strings.Split(submit[i], ",")
		submit_dict[a[0]] = [3]string{a[1], a[2], a[3]}
	}

	//判断分数
	var score [4]int
	for key, value := range submit_dict {
		tmp := 0
		for i := 0; i < 3; i++ {
			if value[i] == answerDict[key][i] {
				score[i+1]++
				tmp++
			}
		}
		if tmp == 3 {
			score[0]++
		}
	}

	return score, nil
}

func CreateSubmission(name string, avatar string, content string) (error, uint) {
	err, _ := GetUserByName(name)
	if err != nil {
		er, _ := CreateUser(name)
		if er != nil {
			return er, uint(0)
		}
	}

	score, err := Judge(content)
	if err != nil {
		return err, uint(0)
	}
	submission := Submission{
		UserName:  name,
		Avatar:    avatar,
		CreatedAt: time.Now().Unix(),
		Score:     score[0],
		Subscore1: score[1],
		Subscore2: score[2],
		Subscore3: score[3],
	}
	tx := DB.Create(&submission)
	return tx.Error, submission.ID
}

func GetUserSubmissions(name string) (error, []ReturnSub) {
	var AllSub []Submission
	var ReSub []ReturnSub
	//返回某一用户的所有提交
	//在查询时可以使用.Order()来控制结果的顺序，详见https://gorm.io/zh_CN/docs/query.html#Order
	//当然，也可以查询后在这个函数里手动完成排序
	tx := DB.Model(&Submission{}).Where("UserName=?", name).Order("CreatedAt").Find(&AllSub)

	var user User
	DB.Model(&User{}).Where("UserName=?", name).Find(&user)
	for _, sub := range AllSub {
		ReSub = append(ReSub, ReturnSub{
			UserName:  sub.UserName,
			Avatar:    sub.Avatar,
			CreatedAt: sub.CreatedAt,
			Score:     sub.Score,
			UserVotes: int(user.Votes),
			Subscore1: sub.Subscore1,
			Subscore2: sub.Subscore2,
			Subscore3: sub.Subscore3,
		})
	}
	return tx.Error, ReSub
}

func GetLeaderBoard() []ReturnSub {
	//一个可行的思路，先全部选出submission，然后手动选出每个用户的最后一次提交
	var AllSub []Submission
	DB.Model(&Submission{}).Where("1=1").Find(&AllSub)
	//在这里添加逻辑！
	result := make(map[string]Submission)
	for i := 0; i < len(AllSub); i++ {
		last, err := result[AllSub[i].UserName]
		if !err {
			result[AllSub[i].UserName] = AllSub[i]
		} else {
			if last.CreatedAt < AllSub[i].CreatedAt {
				result[last.UserName] = AllSub[i]
			}
		}
	}
	var ReturnSlice = []ReturnSub{}
	for _, k := range result {
		var votes User
		DB.Model(&User{}).Where("user_name=?", k.UserName).First(&votes)
		ReturnSlice = append(ReturnSlice, ReturnSub{
			UserName:  k.UserName,
			Avatar:    k.Avatar,
			CreatedAt: k.CreatedAt,
			Score:     k.Score,
			UserVotes: int(votes.Votes),
			Subscore1: k.Subscore1,
			Subscore2: k.Subscore2,
			Subscore3: k.Subscore3,
		})
	}
	sort.Slice(ReturnSlice, func(i, j int) bool { return ReturnSlice[i].Score > ReturnSlice[j].Score })

	return ReturnSlice
}
