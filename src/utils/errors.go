package utils

import (
	"github.com/gofiber/fiber/v2/log"
)

// Executes the given function within a safe context, recovers from panics
// and logs the error to the console or invoke a given callback.
func Protect(f func(), onPanic ...func(err any)) {
	defer func() {
		if err := recover(); err != nil {
			if len(onPanic) > 0 {
				onPanic[0](err)
			} else {
				log.Errorw("Recovered from panic", "error", err)
			}
		}
	}()
	f()
}
