package core

type HandlerFunc func(detail *Detail) error

func (f HandlerFunc) Notify(detail *Detail) error {
	// send sns message to details service
	return nil
}
