package offer_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/julianoL13/app-game-deal-hunter/offer"
	"github.com/julianoL13/app-game-deal-hunter/offer/mocks"
)

func TestGetOffersByGameId(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		gameID := uuid.New()
		expectedOffers := []*offer.Offer{
			{
				ID:     uuid.New(),
				GameID: gameID,
				Store:  offer.StoreSteam,
				URL:    "http://example.com/offer1",
			},
			{
				ID:     uuid.New(),
				GameID: gameID,
				Store:  offer.StoreEpic,
				URL:    "http://example.com/offer2",
			},
		}

		reader := new(mocks.MockReader)
		reader.On("FindByGameID", ctx, gameID).Return(expectedOffers, nil)

		uc := offer.NewGetOffersByGameIdUseCase(reader)
		offers, err := uc.Execute(ctx, gameID)

		assert.NoError(t, err)
		assert.Equal(t, expectedOffers, offers)
		reader.AssertExpectations(t)
	})

	t.Run("fail on nil", func(t *testing.T) {
		gameID := uuid.New()
		expectedErr := assert.AnError

		reader := new(mocks.MockReader)
		reader.On("FindByGameID", ctx, gameID).Return(nil, expectedErr)

		uc := offer.NewGetOffersByGameIdUseCase(reader)
		offers, err := uc.Execute(ctx, gameID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, offers)
		reader.AssertExpectations(t)
	})

	t.Run("empty result", func(t *testing.T) {
		gameID := uuid.New()
		expectedOffers := []*offer.Offer{}

		reader := new(mocks.MockReader)
		reader.On("FindByGameID", ctx, gameID).Return(expectedOffers, nil)

		uc := offer.NewGetOffersByGameIdUseCase(reader)
		offers, err := uc.Execute(ctx, gameID)

		assert.NoError(t, err)
		assert.Empty(t, offers)
		reader.AssertExpectations(t)
	})
}
