package main

import (
	"strings"
	"testing"
)

// Test básico - aprovechando tu experiencia QA
func TestCreateGreeting(t *testing.T) {
	// Arrange
	role := "Backend Developer"
	expected := "¡Hola futuro Backend Developer!"

	// Act
	result := createGreeting(role)

	// Assert
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

// Test con casos edge - como harías en automation
func TestCreateGreeting_EdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "¡Hola futuro !"},
		{"special chars", "Go-Developer!", "¡Hola futuro Go-Developer!!"},
		{"numbers", "Dev123", "¡Hola futuro Dev123!"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := createGreeting(tc.input)
			if result != tc.expected {
				t.Errorf("Test '%s' failed. Expected '%s', got '%s'",
					tc.name, tc.expected, result)
			}
		})
	}
}

// Test que verifica el comportamiento, no solo el output
func TestCreateGreeting_Behavior(t *testing.T) {
	role := "QA Engineer"
	result := createGreeting(role)

	// Verificaciones múltiples (como en tus tests de automation)
	if !strings.Contains(result, role) {
		t.Error("Result should contain the input role")
	}

	if !strings.HasPrefix(result, "¡Hola futuro") {
		t.Error("Result should start with greeting prefix")
	}

	if !strings.HasSuffix(result, "!") {
		t.Error("Result should end with exclamation mark")
	}
}
