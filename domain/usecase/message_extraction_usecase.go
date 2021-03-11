package usecase

import (
	"errors"
	"fmt"
	"strings"
)

type MessageBuilderUsecase struct {
}

func (u *MessageBuilderUsecase) Process(messages [][]string) (string, error) {
	result := make([]string, u.countArrayMaximum(messages))
	for v := range messages {
		msg := messages[v]
		vals := len(result)
		if len(msg) < vals {
			return "", errors.New("INCOMPLETE MESSAGE!")
		}
		for i := 0; i < vals; i++ {
			if len(msg[i]) > 0 {
				if len(result[i]) > 0 && result[i] != msg[i] {
					fmt.Printf("ERROR Message missmatch: [%s - %s]", result[i], msg[i])
					return "", errors.New("Message missmatch")
				}
				result[i] = msg[i]
			}
		}
	}
	return strings.Join(result[:], " "), nil
}

func (u *MessageBuilderUsecase) countArrayMaximum(messages [][]string) int {
	max := 0
	for v := range messages {
		if len(messages[v]) >= max {
			max = len(messages[v])
		}
	}
	return max
}
