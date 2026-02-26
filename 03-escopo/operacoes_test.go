package main

import "testing"

func TestSoma(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positivos", 2, 3, 5},
		{"negativos", -2, -3, -5},
		{"zero", 0, 5, 5},
		{"misturado", -2, 5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soma(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("soma(%d, %d) = %d; esperado %d", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}

func TestSubtrair(t *testing.T) {
	result := subtrair(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("subtrair(10, 5) = %d; esperado %d", result, expected)
	}
}

func TestMultiplicacao(t *testing.T) {
	result := multiplicacao(3, 4)
	expected := 12
	if result != expected {
		t.Errorf("multiplicacao(3, 4) = %d; esperado %d", result, expected)
	}
}

func TestDivisao(t *testing.T) {
	result := divisao(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("divisao(10, 2) = %d; esperado %d", result, expected)
	}
}
