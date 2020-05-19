package database

import "github.com/Aoang/haiju/pkg/model"

func GetAllCommunity() []*model.Community {
	var i []*model.Community
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *communityService) Get(i *model.Community) error {
	if err := db.First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *communityService) Add(i *model.Community) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *communityService) Update(i *model.Community) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *communityService) Delete(i *model.Community) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetAllRoom() []*model.Room {
	var i []*model.Room
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *roomService) Get(i *model.Room) error {
	if err := db.First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *roomService) Add(i *model.Room) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *roomService) Update(i *model.Room) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *roomService) Delete(i *model.Room) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetAllUser() []*model.User {
	var i []*model.User
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *userService) Get(i *model.User) error {
	if i.ID != 0 {
		return db.First(i).Error
	}
	if err := db.Where("username = ?", i.Username).First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *userService) Add(i *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *userService) Update(i *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *userService) Delete(i *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetLimitBuilding(sum int) ([]*model.Building, int) {
	var i []*model.Building
	j := 0
	if err := db.Offset(sum).Limit(10).Find(&i).Error; nil != err {
		return nil, 0
	}
	db.Model(&model.Building{}).Count(&j)

	return i, j
}

func GetAllBuilding() []*model.Building {
	var i []*model.Building
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *buildingService) Get(i *model.Building) error {
	if err := db.First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *buildingService) Add(i *model.Building) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *buildingService) Update(i *model.Building) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *buildingService) Delete(i *model.Building) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetAllLocation() []*model.Location {
	var i []*model.Location
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *locationService) Get(i *model.Location) error {
	if err := db.First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *locationService) Add(i *model.Location) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *locationService) Update(i *model.Location) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *locationService) Delete(i *model.Location) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetAllFile() []*model.File {
	var i []*model.File
	if err := db.Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func GetFile(id uint64) []*model.File {
	var i []*model.File
	if err := db.Where("oid = ?", id).Find(&i).Error; nil != err {
		return nil
	}
	return i
}

func (srv *fileService) Get(i *model.File) error {
	if err := db.Where("oid = ?", i.OID).First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *fileService) Add(i *model.File) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *fileService) Update(i *model.File) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *fileService) Delete(i *model.File) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func GetAllProperty() []*model.Property {
	var property []*model.Property
	if err := db.Find(&property).Error; nil != err {
		return nil
	}
	return property
}

func (srv *propertyService) Get(i *model.Property) error {
	if err := db.First(i).Error; nil != err {
		return err
	}
	return nil
}

func (srv *propertyService) Add(i *model.Property) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *propertyService) Update(i *model.Property) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (srv *propertyService) Delete(i *model.Property) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Delete(i).Error; nil != err {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
