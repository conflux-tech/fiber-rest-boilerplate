package users

// Repository interface allows us to access the CRUD operations
type Repository interface {
	Create(user *User) (*User, error)
	List(page, pageSize int) ([]*User, error)
	Get(id int) (*User, error)
	Update(id int, user *User) (*User, error)
	Delete(id int) (bool, error)
}
