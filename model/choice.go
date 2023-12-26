package model

import "errors"

// Choice enum
type Choice int

const (
	ROCK Choice = iota
	PAPER
	SCISSOR
)

func (c Choice) String() string {
	switch c {
	case ROCK:
		return "rock"
	case PAPER:
		return "paper"
	case SCISSOR:
		return "scissor"
	default:
		return "unknown"
	}
}

func GetChoice(choice string) (Choice, error) {
	switch choice {
	case "rock":
		return ROCK, nil
	case "paper":
		return PAPER, nil
	case "scissor":
		return SCISSOR, nil
	default:
		return 0, errors.New("invalid option")
	}
}
