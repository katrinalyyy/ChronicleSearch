package repository

import (
	"Lab1/intermal/app/ds"
	"errors"

	"gorm.io/gorm"
)

func (r *Repository) CreateUser(user ds.User) (ds.User, error) {
	if user.Name == "" {
		user.Name = "Пользователь"
	}
	err := r.db.Create(&user).Error
	return user, err
}

func (r *Repository) GetUserByID(id uint) (ds.User, error) {
	var u ds.User
	err := r.db.First(&u, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ds.User{}, errors.New("user not found")
		}
		return ds.User{}, err
	}
	return u, nil
}

func (r *Repository) GetUserByEmail(email string) (ds.User, error) {
	var u ds.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ds.User{}, errors.New("user not found")
		}
		return ds.User{}, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(id uint, user ds.User) error {
	updates := map[string]interface{}{
		"name": user.Name,
	}

	// is_moderator может обновлять только администратор
	updates["is_moderator"] = user.IsModerator

	return r.db.Model(&ds.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository) CheckCredentials(email, password string) (ds.User, error) {
	u, err := r.GetUserByEmail(email)
	if err != nil {
		return ds.User{}, err
	}
	if u.Password != password {
		return ds.User{}, errors.New("invalid credentials")
	}
	return u, nil
}
