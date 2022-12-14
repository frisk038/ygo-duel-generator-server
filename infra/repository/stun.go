package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Client struct {
	conn *pgx.Conn
}

const (
	insertOffer     = "insertOffer"
	updateOfferICE  = "updateOfferICE"
	selectOffer     = "selectOffer"
	selectOfferICE  = "selectOfferICE"
	updateAnswer    = "updateAnswer"
	updateAnswerICE = "updateAnswerICE"
	selectAnswer    = "selectAnswer"
	selectAnswerICE = "selectAnswerICE"
)

var queries = map[string]string{
	insertOffer:     "INSERT INTO stun (id, offer, created_at) VALUES ($1, $2, now());",
	updateOfferICE:  "UPDATE stun SET offer_ice=$1 WHERE id=$2;",
	updateAnswer:    "UPDATE stun SET answer=$1 WHERE id=$2;",
	updateAnswerICE: "UPDATE stun SET answer_ice=$1 WHERE id=$2;",
	selectOffer:     "SELECT offer FROM stun WHERE id=$1;",
	selectOfferICE:  "SELECT offer_ice FROM stun WHERE id=$1;",
	selectAnswer:    "SELECT answer FROM stun WHERE id=$1;",
	selectAnswerICE: "SELECT answer_ice FROM stun WHERE id=$1;",
}

func NewClient(ctx context.Context) (*Client, error) {
	connStr := os.Getenv("DATABASE_URL")
	if len(connStr) == 0 {
		return nil, fmt.Errorf("can't get database url from env")
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	for tag, q := range queries {
		_, err := conn.Prepare(ctx, tag, q)
		if err != nil {
			return nil, err
		}
	}
	return &Client{conn: conn}, nil
}

func (c *Client) InsertOffer(ctx context.Context, roomID uuid.UUID, offer []byte) error {
	_, err := c.conn.Exec(ctx, insertOffer, roomID, offer)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateOfferICE(ctx context.Context, roomID uuid.UUID, ice []byte) error {
	_, err := c.conn.Exec(ctx, updateOfferICE, ice, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateAnswer(ctx context.Context, roomID uuid.UUID, answer []byte) error {
	_, err := c.conn.Exec(ctx, updateAnswer, answer, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateAnswerICE(ctx context.Context, roomID uuid.UUID, ice []byte) error {
	_, err := c.conn.Exec(ctx, updateAnswerICE, ice, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SelectOffer(ctx context.Context, roomID uuid.UUID) (*string, error) {
	var offer string
	err := c.conn.QueryRow(ctx, selectOffer, roomID).Scan(&offer)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	return &offer, nil
}

func (c *Client) SelectOfferICE(ctx context.Context, roomID uuid.UUID) (*string, error) {
	var ice string
	err := c.conn.QueryRow(ctx, selectOfferICE, roomID).Scan(&ice)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	return &ice, nil
}

func (c *Client) SelectAnswer(ctx context.Context, roomID uuid.UUID) (*string, error) {
	var answer string
	err := c.conn.QueryRow(ctx, selectAnswer, roomID).Scan(&answer)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	return &answer, nil
}

func (c *Client) SelectAnswerICE(ctx context.Context, roomID uuid.UUID) (*string, error) {
	var ice string
	err := c.conn.QueryRow(ctx, selectAnswerICE, roomID).Scan(&ice)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	return &ice, nil
}
