package service

import (
	"server/mdl"
)

func WorkCreate(loginUser model.User, params model.CreateWorkParam) (model.Work, error) {
	work := model.Work {
		Title: params.Title,
		Description: params.Description,
		StartTime: params.StartTime,
		Drinks: params.Drinks,
	}
	err := Db.Model(&loginUser).Association("Work").Append(work)
	return work, err
}

func WorkRead(id uint) (model.Work, error) {
	var work model.Work
	result := Db.First(&work, id)
	return work, result.Error
}

func WorkUpdate(loginUser model.User, id uint, params model.UpdateWorkParam) (model.Work, error) {
	var work model.Work
	err := Db.Model(&loginUser).Association("Work").Find(&work, id)
	if err != nil {
		return model.Work{}, err
	}
	work.Title = params.Title
	work.Description = params.Description
	work.EndTime = params.EndTime
	result := Db.Save(&work)
	return work, result.Error
}

func WorkDelete(loginUser model.User, id uint) (model.Work, error) {
	var work model.Work
	err := Db.Model(&loginUser).Association("Work").Find(&work, id)
	if err != nil {
		return model.Work{}, err
	}
	result := Db.Delete(&work)
	return work, result.Error
}

func WorkList() ([]model.Work, error) {
	var works []model.Work
	result := Db.Find(&works)
	return works, result.Error
}

func DrinkCreate(loginUser model.User, workId uint, params model.CreateDrinkParam) (model.Drink, error) {
	var work model.Work
	drink := model.Drink {
		Name: params.Name,
		Description: params.Description,
		Image: params.Image,
		Alcohol: params.Alcohol,
		Amount: params.Amount,
	}
	err := Db.Model(&loginUser).Association("Work").Find(&work, workId)
	Db.Model(&work).Association("Drink").Append(&drink)
	return drink, err
}

func DrinkRead(workId uint, id uint) (model.Drink, error) {
	var work model.Work
	var drink model.Drink
	err := Db.Find(&work, workId).Association("Drink").Find(&drink, id)
	return drink, err
}

func DrinkUpdate(loginUser model.User, workId uint, id uint, params model.UpdateDrinkParam) (model.Drink, error) {
	var work model.Work
	var drink model.Drink
	err := Db.Model(&loginUser).Association("Work").Find(&work, workId)
	if err != nil {
		return model.Drink{}, err
	}
	Db.Model(&work).Association("Drink").Find(&drink, id)
	drink.Name = params.Name
	drink.Description = params.Description
	drink.Image = params.Image
	drink.Alcohol = params.Alcohol
	drink.Amount = params.Amount
	result := Db.Save(&drink)
	return drink, result.Error
}

func DrinkDelete(loginUser model.User, workId uint, id uint) (model.Drink, error) {
	var work model.Work
	var drink model.Drink
	err := Db.Model(&loginUser).Association("Work").Find(&work, workId)
	if err != nil {
		return model.Drink{}, err
	}
	Db.Model(&work).Association("Drink").Delete(&drink, id)
	return drink, err
}

func DrinkList(workId uint) ([]model.Drink, error) {
	var work model.Work
	var drinks []model.Drink
	err := Db.Find(&work, workId).Association("Drink").Find(&drinks)
	return drinks, err
}
