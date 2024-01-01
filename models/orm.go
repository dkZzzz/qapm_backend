package models

func CreateOptimize(optimize *Optimize) (int, error) {
	// 将新的 Optimize 对象保存到数据库
	result := DB.Create(optimize)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return optimize.ID, nil
	}
}

func SelectOptimizeByUser(user int) ([]Optimize, error) {
	var optimize []Optimize
	result := DB.Where("user = ?", user).Find(&optimize)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return optimize, nil
	}
}

func SelectOptimizeById(id int) (*Optimize, error) {
	var optimize Optimize
	result := DB.Where("id = ?", id).First(&optimize)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &optimize, nil
	}
}
