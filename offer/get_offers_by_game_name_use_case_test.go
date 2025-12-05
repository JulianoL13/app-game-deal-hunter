package offer_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/julianoL13/app-game-deal-hunter/offer"
	"github.com/julianoL13/app-game-deal-hunter/offer/mocks"
)

func TestGetOffersByGameName(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		gameName := "Elden Ring"
		gameID := uuid.New()
		expectedOffers := []*offer.Offer{
			{
				ID:     uuid.New(),
				GameID: gameID,
				Store:  offer.StoreSteam,
				URL:    "http://store.steampowered.com/app/1245620",
			},
			{
				ID:     uuid.New(),
				GameID: gameID,
				Store:  offer.StoreNuuvem,
				URL:    "http://nuuvem.com/item/elden-ring",
			},
		}

		reader := new(mocks.MockReader)
		reader.On("FindByGameName", ctx, gameName).Return(expectedOffers, nil)

		uc := offer.NewGetOffersByGameNameUseCase(reader)
		offers, err := uc.Execute(ctx, gameName)

		assert.NoError(t, err)
		assert.Equal(t, expectedOffers, offers)
		reader.AssertExpectations(t)
	})

	t.Run("fail on error", func(t *testing.T) {
		gameName := "Elden Ring"
		expectedErr := assert.AnError

		reader := new(mocks.MockReader)
		reader.On("FindByGameName", ctx, gameName).Return(nil, expectedErr)

		uc := offer.NewGetOffersByGameNameUseCase(reader)
		offers, err := uc.Execute(ctx, gameName)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, offers)
		reader.AssertExpectations(t)
	})

	t.Run("empty result", func(t *testing.T) {
		gameName := "Game That Does Not Exist"
		expectedOffers := []*offer.Offer{}

		reader := new(mocks.MockReader)
		reader.On("FindByGameName", ctx, gameName).Return(expectedOffers, nil)

		uc := offer.NewGetOffersByGameNameUseCase(reader)
		offers, err := uc.Execute(ctx, gameName)

		assert.NoError(t, err)
		assert.Empty(t, offers)
		reader.AssertExpectations(t)
	})

	t.Run("case sensitivity", func(t *testing.T) {
		gameName := "elden ring"
		gameID := uuid.New()
		expectedOffers := []*offer.Offer{
			{
				ID:     uuid.New(),
				GameID: gameID,
				Store:  offer.StoreSteam,
				URL:    "http://store.steampowered.com/app/1245620",
			},
		}

		reader := new(mocks.MockReader)
		reader.On("FindByGameName", ctx, gameName).Return(expectedOffers, nil)

		uc := offer.NewGetOffersByGameNameUseCase(reader)
		offers, err := uc.Execute(ctx, gameName)

		assert.NoError(t, err)
		assert.Equal(t, expectedOffers, offers)
		reader.AssertExpectations(t)
	})
}
