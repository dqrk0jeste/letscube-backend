package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("dadsdfdsfsdfdsfsdfsdfsdffrkjmbdx")
	require.NoError(t, err)

	id := uuid.New()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(id, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	var payload *Payload

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, id, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker("dadsdfdsfsdfdsfsdfsdfsdffrkjmbdx")
	require.NoError(t, err)

	token, err := maker.CreateToken(uuid.New(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	var payload *Payload
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
