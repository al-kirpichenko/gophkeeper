package auth

import (
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/al-kirpichenko/gophkeeper/internal/utils/sl"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// Auth - сервис регистрации и аутентификации
type Auth struct {
	log         *slog.Logger
	usrProvider UserProvider
	tokenTTL    time.Duration
}

// UserProvider -
type UserProvider interface {
	CreateUser(auth *models.Auth) error
	GetUser(login string) (*models.User, error)
}

// NewAuth - конструктор
func NewAuth(
	log *slog.Logger,
	userProvider UserProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrProvider: userProvider,
		log:         log,
		tokenTTL:    tokenTTL, // Время жизни возвращаемых токенов
	}
}

// CreateUser - создание нового пользователя
func (a *Auth) CreateUser(auth *models.Auth) error {

	a.log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword(auth.Password, bcrypt.DefaultCost)

	if err != nil {
		a.log.Error("failed to generate password hash", sl.Err(err))

		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser.CreatePassHash", err)
	}

	auth.Password = passHash

	err = a.usrProvider.CreateUser(auth)

	if err != nil {
		a.log.Error("failed to create new user", sl.Err(err))

		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser", err)
	}

	return nil
}

func (a *Auth) GetUser(login string) (*models.User, error) {
	return &models.User{}, nil
}
