package business

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type storer interface {
	SelectAnswerICE(ctx context.Context, roomID uuid.UUID) (*string, error)
	SelectAnswer(ctx context.Context, roomID uuid.UUID) (*string, error)
	SelectOfferICE(ctx context.Context, roomID uuid.UUID) (*string, error)
	SelectOffer(ctx context.Context, roomID uuid.UUID) (*string, error)
	InsertOffer(ctx context.Context, roomID uuid.UUID, offer []byte) error
	UpdateAnswerICE(ctx context.Context, roomID uuid.UUID, ice []byte) error
	UpdateAnswer(ctx context.Context, roomID uuid.UUID, answer []byte) error
	UpdateOfferICE(ctx context.Context, roomID uuid.UUID, ice []byte) error
}
type StunBusiness struct {
	store storer
}

func NewStunBusiness(store storer) *StunBusiness {
	return &StunBusiness{store: store}
}

func (sb *StunBusiness) StoreOffer(ctx context.Context, roomID string, offer []byte) error {
	if !json.Valid(offer) {
		return fmt.Errorf("offer is not valid")
	}

	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return err
	}

	return sb.store.InsertOffer(ctx, roomUUID, offer)
}

func (sb *StunBusiness) StoreOfferICE(ctx context.Context, roomID string, ice []byte) error {
	if !json.Valid(ice) {
		return fmt.Errorf("offer ice is not valid")
	}

	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return err
	}

	return sb.store.UpdateOfferICE(ctx, roomUUID, ice)
}

func (sb *StunBusiness) StoreAnswer(ctx context.Context, roomID string, answer []byte) error {
	if !json.Valid(answer) {
		return fmt.Errorf("answer is not valid")
	}

	fmt.Println("roomID", roomID)
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return err
	}

	return sb.store.UpdateAnswer(ctx, roomUUID, answer)
}

func (sb *StunBusiness) StoreAnswerICE(ctx context.Context, roomID string, ice []byte) error {
	if !json.Valid(ice) {
		return fmt.Errorf("answer ice is not valid")
	}

	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return err
	}

	return sb.store.UpdateAnswerICE(ctx, roomUUID, ice)
}

func (sb *StunBusiness) GetOffer(ctx context.Context, roomID string) (*string, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	return sb.store.SelectOffer(ctx, roomUUID)
}

func (sb *StunBusiness) GetOfferICE(ctx context.Context, roomID string) (*string, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	return sb.store.SelectOfferICE(ctx, roomUUID)
}

func (sb *StunBusiness) GetAnswer(ctx context.Context, roomID string) (*string, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	return sb.store.SelectAnswer(ctx, roomUUID)
}

func (sb *StunBusiness) GetAnswerICE(ctx context.Context, roomID string) (*string, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return nil, err
	}
	return sb.store.SelectAnswerICE(ctx, roomUUID)
}
