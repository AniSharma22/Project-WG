package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"project2/internal/domain/entities"
	"project2/pkg/globals"
	"testing"
	"time"
)

func TestSlotService_GetGameSlots(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	gameID := primitive.NewObjectID()
	slots := []entities.Slot{
		{ID: primitive.NewObjectID(), GameID: gameID, StartTime: time.Now().Add(1 * time.Hour)},
		{ID: primitive.NewObjectID(), GameID: gameID, StartTime: time.Now().Add(2 * time.Hour)},
	}

	mockSlotRepo.EXPECT().
		GetSlotsByDate(gomock.Any(), gameID).
		Return(slots, nil)

	result, err := slotService.GetGameSlots(&entities.Game{ID: gameID})
	assert.NoError(t, err)
	assert.ElementsMatch(t, slots, result)
}

func TestSlotService_BookSlot(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	userID := primitive.NewObjectID()
	gameID := primitive.NewObjectID()
	slotID := primitive.NewObjectID()

	globals.ActiveUser = "test@example.com"
	slot := &entities.Slot{
		ID:          slotID,
		GameID:      gameID,
		StartTime:   time.Now().Add(1 * time.Hour),
		BookedUsers: []primitive.ObjectID{},
	}

	mockUserRepo.EXPECT().
		GetUserByEmail(globals.ActiveUser).
		Return(&entities.User{ID: userID}, nil)

	mockSlotRepo.EXPECT().
		BookSlot(userID, slotID).
		Return(nil)

	mockGameHistoryRepo.EXPECT().
		AddGameHistory(gomock.Any()).
		Return(nil)

	err := slotService.BookSlot(&entities.Game{ID: gameID}, slot)
	assert.NoError(t, err)
}

func TestSlotService_InviteToSlot(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	userID := primitive.NewObjectID()
	slotID := primitive.NewObjectID()

	slot := &entities.Slot{
		ID:          slotID,
		StartTime:   time.Now().Add(1 * time.Hour),
		BookedUsers: []primitive.ObjectID{},
	}

	mockUserRepo.EXPECT().
		AddToInvites(userID, slotID).
		Return(nil)

	err := slotService.InviteToSlot(userID, &entities.Game{}, slot)
	assert.NoError(t, err)
}

func TestSlotService_GetSlotById(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	slotID := primitive.NewObjectID()
	slot := &entities.Slot{ID: slotID}

	mockSlotRepo.EXPECT().
		GetSlotById(slotID).
		Return(slot, nil)

	result, err := slotService.GetSlotById(slotID)
	assert.NoError(t, err)
	assert.Equal(t, slotID, result.ID)
}

func TestSlotService_GetUpcomingBookedSlots(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	userID := primitive.NewObjectID()
	slots := []entities.Slot{
		{ID: primitive.NewObjectID(), StartTime: time.Now().Add(1 * time.Hour)},
	}

	globals.ActiveUser = "test@example.com"

	mockUserRepo.EXPECT().
		GetUserByEmail(globals.ActiveUser).
		Return(&entities.User{ID: userID}, nil)

	mockSlotRepo.EXPECT().
		GetUpcomingBookedSlots(userID).
		Return(slots, nil)

	result, err := slotService.GetUpcomingBookedSlots()
	assert.NoError(t, err)
	assert.ElementsMatch(t, slots, result)
}

func TestSlotService_AddResultToSlot(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	userID := primitive.NewObjectID()
	slotID := primitive.NewObjectID()
	result := "Win"

	mockSlotRepo.EXPECT().
		AddResultToSlot(userID, slotID, result).
		Return(nil)

	err := slotService.AddResultToSlot(userID, slotID, result)
	assert.NoError(t, err)
}
