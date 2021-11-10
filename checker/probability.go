package checker

type Probability int

const (
	LOW       Probability = iota
	MEDIUM    Probability = iota
	HIGH      Probability = iota
	VERY_HIGH Probability = iota
)

func (p Probability) String() string {
	switch p {
	case LOW:
		return "Low"
	case MEDIUM:
		return "Medium"
	case HIGH:
		return "High"
	case VERY_HIGH:
		return "Very high"
	default:
		return "???"
	}
}

