package validations

import "github.com/burl/inquire/widget"

func Required(w *widget.Input) {
	w.Valid(func(value string) string {
		if len(value) < 1 {
			return "You must return a value"
		}
		return ""
	})
}
