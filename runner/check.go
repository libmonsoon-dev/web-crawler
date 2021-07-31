package runner

import (
	"context"
	"fmt"
)

var _ Interface = new(errWrapper)

type errWrapper struct {
	message string
	r       Interface
}

func Errorf(message string, r Interface) Interface {
	return errWrapper{message, r}
}

func (c errWrapper) Run(ctx context.Context) error {
	err := c.r.Run(ctx)
	if err != nil {
		err = fmt.Errorf(c.message+": %w", err)
	}

	return err
}
