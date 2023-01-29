package builder

import (
	"time"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	domainservice "github.com/juliocesarscheidt/go-orm-api/domain/service"
)

type UserBuilder struct {
	PasswordService domainservice.PasswordService
}

func (builder UserBuilder) NewUser(name string, email string, password string) (*entity.User, error) {
	err := entity.ValidateUserFields(map[string]string{"Name": name, "Email": email, "Password": password})
	if err != nil {
		return nil, err
	}
	hashedPassword, _ := builder.PasswordService.EncryptPassword(password)
	user := &entity.User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	return user, nil
}

func (builder UserBuilder) AlterUser(name string, password string) (*entity.User, error) {
	err := entity.ValidateUserFields(map[string]string{"Name": name, "Password": password})
	if err != nil {
		return nil, err
	}
	hashedPassword, _ := builder.PasswordService.EncryptPassword(password)
	user := &entity.User{
		Name:      name,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}
	return user, nil
}