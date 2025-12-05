package game_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/julianoL13/app-game-deal-hunter/game"
	"github.com/julianoL13/app-game-deal-hunter/game/mocks"
)

func TestCreateGame(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		g, _ := game.New(
			"Test Game",
			"A game for testing",
			"Test Dev",
			"Test Pub",
			time.Now(),
		)

		writer := new(mocks.MockWriter)
		writer.On("Insert", ctx, g).Return(nil)
		uc := game.NewCreateGameUseCase(writer)
		err := uc.Execute(ctx, g)
		assert.NoError(t, err)
		writer.AssertExpectations(t)
	})
	t.Run("fail", func(t *testing.T) {
		g, _ := game.New(
			"Test Game",
			"A game for testing",
			"Test Dev",
			"Test Pub",
			time.Now(),
		)

		expectedErr := assert.AnError
		writer := new(mocks.MockWriter)
		writer.On("Insert", ctx, g).Return(expectedErr)
		uc := game.NewCreateGameUseCase(writer)
		err := uc.Execute(ctx, g)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		writer.AssertExpectations(t)
	})

}
