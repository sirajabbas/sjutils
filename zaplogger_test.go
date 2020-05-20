package sjutils

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger("info.log", "err.log")
	t.Log("Test completed")
}
