package offer_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/julianoL13/app-game-deal-hunter/offer"
	"github.com/julianoL13/app-game-deal-hunter/offer/mocks"
)

func TestCreateOffer(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		o, _ := offer.NewOffer(
			uuid.New(),
			offer.StoreSteam,
			"http://example.com/offer",
		)

		writer := new(mocks.MockWriter)
		writer.On("Insert", ctx, o).Return(nil)
		uc := offer.NewCreateOfferUseCase(writer)
		err := uc.Execute(ctx, o)
		assert.NoError(t, err)
		writer.AssertExpectations(t)
	})
	t.Run("fail", func(t *testing.T) {
		o, _ := offer.NewOffer(
			uuid.New(),
			offer.StoreSteam,
			"http://example.com/offer",
		)

		expectedErr := assert.AnError
		writer := new(mocks.MockWriter)
		writer.On("Insert", ctx, o).Return(expectedErr)
		uc := offer.NewCreateOfferUseCase(writer)
		err := uc.Execute(ctx, o)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		writer.AssertExpectations(t)
	})
}
