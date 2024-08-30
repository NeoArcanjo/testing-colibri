package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/NeoArcanjo/testing-colibri/src/domain/models"
	"github.com/NeoArcanjo/testing-colibri/src/domain/usecases"
	"github.com/NeoArcanjo/testing-colibri/src/infra/repositories/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewDemoGetByID(t *testing.T) {
	uc := usecases.NewDemoGetByID()
	assert.NotNil(t, uc)
}

func TestDemoGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockDemoRepository(ctrl)
	uc := usecases.DemoGetByID{Repo: repo}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		id := 1
		expected := &models.Demo{ID: id, Name: "1"}
		repo.EXPECT().FindById(ctx, 1).Return(expected, nil)

		result, err := uc.Execute(context.Background(), id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		id := 1
		repo.EXPECT().FindById(ctx, id).Return(nil, fmt.Errorf("could not find by id"))

		result, err := uc.Execute(context.Background(), id)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
