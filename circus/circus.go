package circus

// Go interfaces encourages one to be lazy, and this is a good thing.
// Instead of writing types to fulfil interfaces, write interfaces to fulfil usage requirements.

type Speaker interface {
	Speaks() string
}

func Perform(a Speaker) string {
	return a.Speaks()
}
