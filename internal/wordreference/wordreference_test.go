package wordreference

import (
	"context"
	"os"
	"testing"
)

func TestPONSLookupGermanEnglish(t *testing.T) {
	if testing.Short() || os.Getenv("TUIFERENCE_LIVE_TESTS") == "" {
		t.Skip("set TUIFERENCE_LIVE_TESTS=1 to run live PONS lookup")
	}

	results, err := NewClient().Lookup(context.Background(), Language{Name: "German", Code: "de"}, Language{Name: "English", Code: "en"}, "Haus")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) == 0 {
		t.Fatal("expected PONS results")
	}
}
