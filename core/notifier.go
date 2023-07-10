package core

type HandlerFunc func(detail *Detail) error

func (f HandlerFunc) Notify(detail *Detail) error {
	return f(detail)
}

var notifier HandlerFunc = func(detail *Detail) error {
	// TODO: Missing implementation
	return nil
}
