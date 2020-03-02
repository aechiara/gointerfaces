package animals

// Go interfaces encourages one to be lazy, and this is a good thing.
// Instead of writing types to fulfil interfaces, write interfaces to fulfil usage requirements.

type Dog struct{}

func (a Dog) Speaks() string {
	return "woof"
}
