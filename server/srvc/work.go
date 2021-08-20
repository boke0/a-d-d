package service

import (
	. "server/mdl"
)

func WorkCreate(loginUser User, params CreateWorkParam) (Work, error) {
	work := Work {
		Title: params.Title,
		Description: params.Description,
		StartTime: params.StartTime,
		EndTime: nil,
		Drinks: params.Drinks,
		UserId: loginUser.ID,
	}
	Db.Create(&work)
	result := Db.Save(&work)
	return work, result.Error
}

func WorkRead(id uint) (Work, error) {
	var work Work
	result := Db.First(&work, id)
	return work, result.Error
}

func WorkInProgressRead(loginUser User) (Work, error) {
	var work Work
	var drinks []Drink
	result := Db.Where("end_time IS NULL AND user_id = ?", loginUser.ID).First(&work)
	Db.Model(&work).Association("Drinks").Find(&drinks)
	work.Drinks = drinks
	return work, result.Error
}

func WorkUpdate(loginUser User, id uint, params UpdateWorkParam) (Work, error) {
	var work Work
	result := Db.Where("user_id = ? and id = ?", loginUser.ID, id).First(&work)
	if result.Error != nil {
		return Work{}, result.Error
	}
	work.Title = params.Title
	work.Description = params.Description
	work.EndTime = &params.EndTime
	result = Db.Save(&work)
	return work, result.Error
}

func WorkDelete(loginUser User, id uint) (Work, error) {
	var work Work
	result := Db.Where("user_id = ? and id = ?", loginUser.ID, id).First(&work)
	if result.Error != nil {
		return Work{}, result.Error
	}
	result = Db.Delete(&work)
	return work, result.Error
}

func WorkList() ([]Work, error) {
	var works []Work
	result := Db.Find(&works)
	return works, result.Error
}

func DrinkCreate(loginUser User, workId uint, params CreateDrinkParam) (Drink, error) {
	var work Work
	drink := Drink {
		Name: params.Name,
		Alcohol: params.Alcohol,
		Amount: params.Amount,
	}
	result := Db.Where("user_id = ? and id = ?", loginUser.ID, workId).First(&work)
	Db.Model(&work).Association("Drinks").Append(&drink)
	return drink, result.Error
}

func DrinkRead(workId uint, id uint) (Drink, error) {
	var drink Drink
	result := Db.Where("work_id = ? and id = ?", workId, id).Find(&drink)
	return drink, result.Error
}

func DrinkUpdate(loginUser User, workId uint, id uint, params UpdateDrinkParam) (Drink, error) {
	var drink Drink
	result := Db.Where("user_id = ? and work_id = ? and id = ?", loginUser.ID, workId, id).First(&drink)
	if result.Error != nil {
		return Drink{}, result.Error
	}
	drink.Name = params.Name
	drink.Alcohol = params.Alcohol
	drink.Amount = params.Amount
	result = Db.Save(&drink)
	return drink, result.Error
}

func DrinkDelete(loginUser User, workId uint, id uint) (Drink, error) {
	var drink Drink
	result := Db.Where("user_id = ? and work_id = ? and id = ?", loginUser.ID, workId, id).Delete(&drink)
	return drink, result.Error
}

func DrinkList(workId uint) ([]Drink, error) {
	var work Work
	var drinks []Drink
	err := Db.Find(&work, workId).Association("Drinks").Find(&drinks)
	return drinks, err
}
