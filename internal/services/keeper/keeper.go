package keeper

import (
	"fmt"
	"log/slog"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
	"github.com/al-kirpichenko/gophkeeper/internal/utils/sl"
)

type Keeper struct {
	log      *slog.Logger
	provider SecretProvider
}

// SecretProvider -
type SecretProvider interface {
	CreateSecret(secret *models.Secret) error
	ReadSecret(title string) (*models.Secret, error)
}

// NewKeeper - конструктор
func NewKeeper(
	log *slog.Logger,
	provider SecretProvider,
) *Keeper {
	return &Keeper{
		provider: provider,
		log:      log,
	}

}

// CreateSecret - создание нового секрета
func (k *Keeper) CreateSecret(secret *models.Secret) error {

	k.log.Info("create secret")

	err := k.provider.CreateSecret(secret)

	if err != nil {
		k.log.Error("Keeper.Create: ", sl.Err(err))

		return fmt.Errorf("%s: %w", "Keeper.Create: ", err)
	}
	return nil
}
