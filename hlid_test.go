package hlid

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

var (
	testTime         = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	testTimeId       = NewWithTime(testTime)
	testTimeIdPrefix = "0fc9379ed800"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func TestNew(t *testing.T) {
	last := time.Now()

	for i := 0; i < 100*365; i++ {
		current := last.Add(24 * time.Hour)
		id := NewWithTime(current)

		if bytes.Compare(testTimeId[:], id[:]) >= 0 {
			t.Fatalf("expected %s < %s", testTimeId, id)
		}

		last = current
	}
}

func TestString(t *testing.T) {
	if !strings.HasPrefix(testTimeId.String(), testTimeIdPrefix) {
		t.Fatalf("expected prefix %s, got %s", testTimeIdPrefix, testTimeId)
	}
}

func TestMarshalUnmarshalBinary(t *testing.T) {
	b, err := testTimeId.MarshalBinary()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var id ID
	err = id.UnmarshalBinary(b)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != testTimeId {
		t.Fatalf("expected %s, got %s", testTimeId, id)
	}
}

func TestMarshalUnmarshalText(t *testing.T) {
	b, err := testTimeId.MarshalText()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var id ID
	err = id.UnmarshalText(b)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != testTimeId {
		t.Fatalf("expected %s, got %s", testTimeId, id)
	}
}

func TestValueScan(t *testing.T) {
	v, err := testTimeId.Value()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var id ID
	err = id.Scan(v)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != testTimeId {
		t.Fatalf("expected %s, got %s", testTimeId, id)
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	b, err := testTimeId.MarshalJSON()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var id ID
	err = id.UnmarshalJSON(b)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != testTimeId {
		t.Fatalf("expected %s, got %s", testTimeId, id)
	}
}
