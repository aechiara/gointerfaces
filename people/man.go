package people

// Go interfaces encourages one to be lazy, and this is a good thing.
// Instead of writing types to fulfil interfaces, write interfaces to fulfil usage requirements.

type Man struct{}

func (a Man) Speaks() string {
	return "Hello World!"
}
