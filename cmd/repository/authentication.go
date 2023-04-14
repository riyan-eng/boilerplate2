package repository

import (
	"boilerplate/cmd/repository/reqres"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan/v2"
)

type AuthenticationQuery interface {
	RegisterAdmin(*reqres.RegisterAdminRequest) reqres.RegisterAdminResponse
	Login(*reqres.LoginRequest) reqres.LoginResponse
}

type authenticationQuery struct {
	database *sql.DB
}

func (a *authenticationQuery) RegisterAdmin(repoReq *reqres.RegisterAdminRequest) (repoRes reqres.RegisterAdminResponse) {
	tx, err := a.database.BeginTx(repoReq.Context, nil)
	if err != nil {
		repoRes.Error = err
		return
	}
	defer tx.Rollback()
	qUserData := `insert into user_datas(name) values('') returning id`
	var userDataID string
	rowUserData, err := tx.QueryContext(repoReq.Context, qUserData)
	if err != nil {
		repoRes.Error = err
		return
	}
	if err := scan.Row(&userDataID, rowUserData); err != nil {
		repoRes.Error = err
		return
	}
	qUser := fmt.Sprintf(`
		insert into users(username, user_type_code, email, password, pin, phone_number, user_data_id) values('%v', '%v', '%v', '%v', '%v', '%v', '%v') returning id
	`, repoReq.Item.UserName, repoReq.Item.UserTypeCode, repoReq.Item.Email, repoReq.Item.Password, repoReq.Item.Pin, repoReq.Item.PhoneNumber, userDataID)
	rowUser, err := tx.QueryContext(repoReq.Context, qUser)
	if err != nil {
		repoRes.Error = err
		return
	}
	if err := scan.Row(&repoRes.Item, rowUser); err != nil {
		repoRes.Error = err
		return
	}
	if err = tx.Commit(); err != nil {
		repoRes.Error = err
		return
	}
	return

}

func (a *authenticationQuery) Login(repoReq *reqres.LoginRequest) (repoRes reqres.LoginResponse) {
	q := fmt.Sprintf(`
	select u.id, u.username, u.email, u."password", coalesce(u.phone_number, '') as phone_number , ut.code as user_type_code, ut."name" as user_type_name, coalesce(ud."name", '') as name 
	from users u left join user_types ut on u.user_type_code = ut.code left join user_datas ud on u.user_data_id = ud.id where u.username = '%v' or u.email = '%v'
	`, repoReq.Item.UserName, repoReq.Item.UserName)
	row, err := a.database.QueryContext(repoReq.Context, q)
	if err != nil {
		repoRes.Error = err
		return
	}
	repoRes.Error = scan.Row(&repoRes.Item, row)
	return
}
