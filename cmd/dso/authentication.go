package dso

type RegisterAdmin struct {
	ID           string
	UserName     string
	UserTypeCode string
	Email        string
	Password     string
	Pin          string
	PhoneNumber  string
	UserDataID   string
	CreatedBy    string
}

type Login struct {
	ID           string `db:"id"`
	UserName     string `db:"username"`
	UserTypeCode string `db:"user_type_code"`
	UserTypeName string `db:"user_type_name"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	PhoneNumber  string `db:"phone_number"`
	Name         string `db:"name"`
}
