package os_test

import (
	"testing"

	"github.com/dawnsgo/dawn_cli/internal/os"
)

func TestIsEmptyDir(t *testing.T) {
	ok, err := os.IsEmptyDir("./gate")

	t.Log(ok, err)
}
