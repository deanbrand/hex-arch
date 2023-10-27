package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 2)
	if err != nil {
		t.Fatalf("Addition failed: %s", err)
	}

	require.Equal(t, answer, int32(3))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(2, 1)
	if err != nil {
		t.Fatalf("Subtraction failed: %s", err)
	}

	require.Equal(t, answer, int32(1))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(2, 3)
	if err != nil {
		t.Fatalf("Multiplication failed: %s", err)
	}

	require.Equal(t, answer, int32(6))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(6, 3)
	if err != nil {
		t.Fatalf("Division failed: %s", err)
	}

	require.Equal(t, answer, int32(2))
}
