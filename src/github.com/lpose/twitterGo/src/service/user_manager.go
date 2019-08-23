package service

import (
	"errors"

	"github.com/lpose/twitterGo/src/domain"
)

var userManager *UserManager

type UserManager struct {
	miUser        *domain.User
	misUsers      []*domain.User
	usersLogiados []*domain.User
}

func NewUserManager() *UserManager {
	userManager := new(UserManager)
	userManager.misUsers = make([]*domain.User, 0)
	userManager.usersLogiados = make([]*domain.User, 0)

	return userManager
}

func GetInstance() *UserManager {
	if userManager == nil {
		userManager := NewUserManager()
		return userManager
	}
	return userManager
}

func (manager *UserManager) Register(user *domain.User) error {
	manager.miUser = user
	manager.misUsers = append(manager.misUsers, user)
	return nil
}

func (manager *UserManager) GetUser() *domain.User {
	return manager.miUser
}

func (manager *UserManager) GetUsersLogin() []*domain.User {
	return manager.usersLogiados
}

func (manager *UserManager) GetUserLogiado(name string) (int, *domain.User) {
	for id, user := range manager.usersLogiados {
		if user.Nick == name {
			return id, user
		}
	}
	return -1, nil
}

func (manager *UserManager) GetUserByNick(userName string) *domain.User {
	for _, user := range manager.misUsers {
		if user.Nick == userName {
			return user
		}
	}
	return nil
}

func (manager *UserManager) Login(nick string, pass string) error {
	user := manager.GetUserByNick(nick)
	if user == nil {
		return errors.New("No hay coincidencia con el logeo\n")
	}
	if user.GetPass() != pass {
		return errors.New("No hay coincidencia con el logeo\n")
	}

	manager.usersLogiados = append(manager.usersLogiados, user)

	return nil
}

func (manager *UserManager) RemoveIndex(s []*domain.User, index int) []*domain.User {
	return append(s[:index], s[index+1:]...)
}

func (manager *UserManager) Logout(nick string, pass string) error {
	id, user := manager.GetUserLogiado(nick)
	if user == nil {
		return errors.New("El usuario no se encuentra logiado")
	}
	if user.GetPass() != pass {
		return errors.New("No hay coincidencia en los datos")
	}
	manager.usersLogiados = manager.RemoveIndex(manager.usersLogiados, id)

	return nil
}
