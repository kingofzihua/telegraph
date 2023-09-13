package log

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultHandler_Handle(t *testing.T) {
	log, w := newBufferLogger()

	log.Info("context", "name")

	assert.Equal(t, w.String(), fmt.Sprintln("INFO msg=context name=KEYVALS UNPAIRED"))
}
