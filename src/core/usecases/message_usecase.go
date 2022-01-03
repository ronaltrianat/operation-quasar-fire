package usecases

import (
	"strings"

	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
)

type MessageUseCase struct{}

func NewMessageUseCase() *MessageUseCase {
	return &MessageUseCase{}
}

func (useCase *MessageUseCase) CalculateMessage(data *domain.SatellitesData) (string, error) {
	msg := data.Satellites[0].Message
	for _, v := range data.Satellites[1:] {
		msg = useCase.mergeMessages(msg, v.Message)
	}

	response := useCase.cleanMessage(msg)

	return response, nil
}

func (useCase MessageUseCase) mergeMessages(messageA, messageB []string) []string {
	msg := []string{}
	maxLen := len(messageA)
	if len(messageB) > len(messageA) {
		maxLen = len(messageB)
	}
	for i := 0; i < maxLen; i++ {
		if i < len(messageA) && messageA[i] != "" {
			msg = append(msg, messageA[i])
		} else if i < len(messageB) && messageB[i] != "" {
			msg = append(msg, messageB[i])
		} else {
			msg = append(msg, "")
		}
	}
	return msg
}

func (useCase MessageUseCase) cleanMessage(message []string) string {
	msg := strings.Builder{}
	for idx, v := range message {
		if idx < len(message)-1 && message[idx] != message[idx+1] {
			msg.WriteString(v)
			msg.WriteString(" ")
		}
	}

	msg.WriteString(message[len(message)-1])

	return msg.String()
}
