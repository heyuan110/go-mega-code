package model

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

// SetPassword func: set password to hash string
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword func
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

// GetUserByUsername func
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username=?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	return db.Create(&user).Error
}
