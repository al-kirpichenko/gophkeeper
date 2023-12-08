package auth

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/al-kirpichenko/gophkeeper/internal/utils/jwt"
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
	CreateUser(user *models.User) error
	GetUser(login string) (*models.User, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

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

	passHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if err != nil {
		a.log.Error("failed to generate password hash", sl.Err(err))

		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser.CreatePassHash", err)
	}

	user := &models.User{
		Login:    auth.Login,
		Password: passHash,
	}

	err = a.usrProvider.CreateUser(user)

	if err != nil {
		a.log.Error("Auth.Register: ", sl.Err(err))

		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser", err)
	}

	return nil
}

// Login - аутентификация пользователя
func (a *Auth) Login(auth *models.Auth) (string, error) {

	a.log.Info("login user")

	// Достаём пользователя из БД
	user, err := a.usrProvider.GetUser(auth.Login)

	if err != nil {
		a.log.Error("Auth.Login: ", sl.Err(err))

		return "", fmt.Errorf("%s: %w", "Auth.Login: ", err)
	}

	// Проверяем корректность полученного пароля
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(auth.Password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", "Auth.Login: ", ErrInvalidCredentials)
	}

	a.log.Info("user logged in successfully")

	// Создаём токен авторизации
	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", "Auth.Login: ", err)
	}

	return token, nil
}
