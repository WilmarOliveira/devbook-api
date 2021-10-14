package entities

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.Format()

	return nil
}

func (user *User) validate() error {

	if user.Name == "" {
		return errors.New("O campo name não pode ficar em branco")
	}

	if user.Nick == "" {
		return errors.New("O campo nick não pode ficar em branco")
	}

	if user.Email == "" {
		return errors.New("O campo email não pode ficar em branco")
	}

	if user.Password == "" {
		return errors.New("O campo password não pode ficar em branco")
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
