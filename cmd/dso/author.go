package dso

type Author struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
}
