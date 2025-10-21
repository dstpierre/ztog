package logparser

import (
	"testing"
)

func TestLogParser(t *testing.T) {
	pairs := []string{
		"[INFO] Lesson learned", "lesson learned:info",
		"[WARNING] I'm using [] in msg", "i'm using [] in msg:warning",
		"[ERROR]   Something went wrong", "something went wrong:error",
		"", "",
		"invalid", "",
	}

	for i := 0; i < len(pairs); i += 2 {
		if v := LogParser(pairs[i]); v != pairs[i+1] {
			t.Errorf("expected %s, got %s", pairs[i+1], v)
		}
	}
}
