package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswer   uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type Answer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	} else {
		return false
	}
}

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

func (q QuestionDifficulty) IsValid() bool {
	if q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard {
		return true
	}
	return false
}

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)
