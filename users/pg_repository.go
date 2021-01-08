package users

import "gorm.io/gorm"

// PGRepository defines the Postgres repository struct
type PGRepository struct {
	DB *gorm.DB
}

// NewPGRepo initializes PGRepo
func NewPGRepo(db *gorm.DB) Repository {
	return &PGRepository{
		DB: db,
	}
}

// Create func creats a database record in Postgres
func (r *PGRepository) Create(user *User) (*User, error) {
	if res := r.DB.Create(user); res.Error != nil {
		return &User{}, res.Error
	}
	return user, nil
}

// List func paginates the list of users
func (r *PGRepository) List(page, pageSize int) ([]*User, error) {
	var users []*User
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	if res := r.DB.Offset(offset).Limit(pageSize).Find(&users); res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

// Get func gets the detail of the user
func (r *PGRepository) Get(id int) (*User, error) {
	user := User{}
	if res := r.DB.Find(&user, id); res.Error != nil {
		return &user, res.Error
	}
	return &user, nil
}

// Update func updates the user record
func (r *PGRepository) Update(id int, user *User) (*User, error) {
	u := User{
		Name: user.Name,
	}
	if res := r.DB.Save(&u); res.Error != nil {
		return &User{}, res.Error
	}
	return &u, nil
}

// Delete func deletes the user record from the database
func (r *PGRepository) Delete(id int) (bool, error) {
	var user User
	r.DB.Find(&user, id)
	if res := r.DB.Delete(&user); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}
