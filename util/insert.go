package util

import (
	"learn-gorm/model"
)

//func Create1() {
//	user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
//	result := Db.Create(&user)
//
//	//默认返回主键
//	fmt.Println(user.ID)
//	printLine()
//	fmt.Println(result.Error)
//	printLine()
//	fmt.Println(result.RowsAffected)
//}

func Create2() {
	user := model.User{
		Name: "空阿道夫",
		//Birthday: time.Now(),
		Age: 22,
	}
	//birthday字段将被忽略，不会写入到数据库
	Db.Create(&user)
	//Db.Select("Name", "Age", "CreatedAt").Create(&user)
}

func genMockUser(n int) []*model.User {
	var res = make([]*model.User, n)
	for i := 0; i < n; i++ {
		res[i] = &model.User{
			Name: "北包包",
			Age:  int8(i),
			//Birthday: time.Now(),
		}
	}
	return res
}

func BatchInset() error {
	mockUsers := genMockUser(20)
	// 传递切片
	tx := Db.Create(&mockUsers)
	//tx := Db.CreateInBatches(&mockUsers, 5)
	return tx.Error
}
