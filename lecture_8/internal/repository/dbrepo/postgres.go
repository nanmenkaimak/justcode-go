package dbrepo

import (
	"errors"
	"github.com/nanmenkaimak/justcode-go/lecture_8/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDBRepo) CreateUser(newUser models.User) (int, error) {
	var userID int
	password, err := hashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	err = m.DB.Get(&userID,
		`insert into users (first_name, last_name, email, password)
				values ($1, $2, $3, $4) returning id`,
		newUser.FirstName, newUser.LastName, newUser.Email, password)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (m *postgresDBRepo) Authenticate(email string, password string) (int, error) {
	var user models.User

	err := m.DB.Get(&user,
		`select email, password from users where email = $1`, email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return 0, errors.New("incorrect password")
	} else if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (m *postgresDBRepo) UpdateUser(user models.User) (bool, error) {
	password, err := hashPassword(user.Password)
	if err != nil {
		return false, err
	}
	res, err := m.DB.Exec(`update users set first_name = $1, last_name = $2, email = $3, password = $4 where id = $5`,
		user.FirstName, user.LastName, user.Email, password, user.ID)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (m *postgresDBRepo) DeleteUserByID(id int) (bool, error) {
	res, err := m.DB.Exec(`delete from users where id = $1`, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
