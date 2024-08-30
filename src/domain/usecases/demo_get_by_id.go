//go:generate mockgen -source demo_get_by_id.go -destination mock/demo_get_by_id_mock.go -package mock
package usecases

import (
	"context"

	"github.com/NeoArcanjo/testing-colibri/src/domain/models"
	"github.com/NeoArcanjo/testing-colibri/src/infra/repositories"
)

type IDemoGetByID interface {
	Execute(context.Context, int) (*models.Demo, error)
}

func NewDemoGetByID() IDemoGetByID {
	return &DemoGetByID{Repo: repositories.NewDemoMemoryRepository()}
}

type DemoGetByID struct {
	Repo repositories.DemoRepository
}

func (uc *DemoGetByID) Execute(ctx context.Context, id int) (*models.Demo, error) {
	return uc.Repo.FindById(ctx, id)
}
