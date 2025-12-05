package game_test

import (
	"context"
	"testing"
	"time"

	"github.com/julianoL13/game-deal-hunter/game"
	"github.com/julianoL13/game-deal-hunter/game/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateGame(t *testing.T) {

	ctx := context.Background()
	t.Run("sucess", func(t *testing.T) {
		input := game.GameInput{
			Title:       "Test Game",
			Description: "A game for testing",
			Developer:   "Test Dev",
			Publisher:   "Test Pub",
			ReleaseDate: time.Now(),
		}

		repo := new(mocks.MockRepository)
		repo.On("Insert", ctx, mock.AnythingOfType("*game.Game")).Return(nil)
		us := game.NewCreateGameUseCase(repo)
		err := us.Execute(ctx, input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("fail", func(t *testing.T) {
		input := game.GameInput{
			Title:       "Test Game",
			Description: "A game for testing",
			Developer:   "Test Dev",
			Publisher:   "Test Pub",
			ReleaseDate: time.Now(),
		}

		expectedErr := assert.AnError
		repo := new(mocks.MockRepository)
		repo.On("Insert", ctx, mock.AnythingOfType("*game.Game")).Return(expectedErr)
		us := game.NewCreateGameUseCase(repo)
		err := us.Execute(ctx, input)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

}
