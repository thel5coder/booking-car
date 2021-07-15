package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"database/sql"
)

type UserCommand struct {
	db    *sql.DB
	model *models.Users
}

func NewUserCommand(db *sql.DB, model *models.Users) commands.IUserCommand {
	return &UserCommand{
		db:    db,
		model: model,
	}
}

func (c UserCommand) Add() (res string, err error) {
	statement := `INSERT INTO users (first_name,last_name,email,username,password,address,phone_number,role_id,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	err = c.db.QueryRow(statement, c.model.FirstName(), c.model.LastName(), c.model.Email(), c.model.UserName(), c.model.Password(), c.model.Address().String, c.model.PhoneNumber(),
		c.model.RoleId(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c UserCommand) Edit() (res string, err error) {
	setStatement := `first_name=$1,last_name=$2,email=$3,username=$4,address=$5,phone_number=$6,role_id=$7,updated_at=$8`
	editParams := []interface{}{c.model.FirstName(), c.model.LastName(), c.model.Email(), c.model.UserName(), c.model.Address(), c.model.PhoneNumber(), c.model.RoleId(), c.model.UpdatedAt(),
		c.model.Id()}
	if c.model.Password() != "" {
		setStatement += `,password=$10`
		editParams = append(editParams, c.model.Password())
	}

	statement := `UPDATE users SET ` + setStatement + ` WHERE id=$9 RETURNING id`
	err = c.db.QueryRow(statement, editParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c UserCommand) Delete() (res string, err error) {
	statement := `UPDATE users SET updated_at=$1,deleted_at=$2 WHERE id=$3`
	err = c.db.QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c UserCommand) EditDeposit() (res string, err error) {
	panic("implement me")
}