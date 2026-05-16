package voice

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/shivamyadav0/vani-ledger-platform/pkg/messaging"
	"go.uber.org/zap"
)

type VoiceClient struct {
	producer *messaging.KafkaProducer
	logger   *zap.Logger
}

type VoiceRequest struct {
	RequestID   string `json:"request_id"`
	UserID      string `json:"user_id"`
	Text        string `json:"text,omitempty"`
	AudioBase64 string `json:"audio_base64,omitempty"`
	CreatedAt   string `json:"created_at"`
}

type VoiceResultData struct {
	UserID      string  `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type VoiceResultResponse struct {
	RequestID  string          `json:"request_id"`
	Status     string          `json:"status"`
	Intent     string          `json:"intent"`
	Transcript string          `json:"transcript,omitempty"`
	UserID     string 		   `json:"user_id"`
	Data       VoiceResultData `json:"data"`
	Error      string          `json:"error,omitempty"`
}

func NewClient(
	producer *messaging.KafkaProducer,
	logger *zap.Logger,
) *VoiceClient {

	return &VoiceClient{
		producer: producer,
		logger:   logger,
	}
}

func (c *VoiceClient) Process(
	ctx context.Context,
	userID,
	text,
	audioBase64 string,
) (string, error) {

	if text == "" && audioBase64 == "" {
		return "", errors.New("missing voice payload")
	}

	requestID, err := generateRequestID()
	if err != nil {
		return "", err
	}

	request := VoiceRequest{
		RequestID:   requestID,
		UserID:      userID,
		Text:        text,
		AudioBase64: audioBase64,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339Nano),
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	if err := c.producer.Produce(
		ctx,
		requestID,
		payload,
	); err != nil {
		return "", err
	}

	c.logger.Info(
		"voice request published",
		zap.String("request_id", requestID),
		zap.String("user_id", userID),
	)

	return requestID, nil
}

func generateRequestID() (string, error) {
	buffer := make([]byte, 16)

	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	return hex.EncodeToString(buffer), nil
}