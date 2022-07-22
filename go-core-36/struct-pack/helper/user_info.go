package helper

import "time"

type UserInfo struct {
	UserName string
	Age      int
	Birthday time.Time
}

func (i *UserInfo) MakeUser(name string, age int, birth time.Time) {
	i.Age = age
	i.Birthday = birth
	i.UserName = name
}

func (i *UserInfo) GetUserName() string {
	return i.UserName
}

func (i *UserInfo) GetUserAge() int {
	return i.Age
}

func (i *UserInfo) GetBirthday() time.Time {
	return i.Birthday
}

func (i *UserInfo) IsAnAudit() bool {
	return i.Age > 18
}
