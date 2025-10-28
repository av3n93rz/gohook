package hook

import (
	"fmt"
	"testing"

	"github.com/vcaesar/tt"
)

func TestAdd(t *testing.T) {
	fmt.Println("hook test...")

	e := Start()
	tt.NotNil(t, e)
}

func TestKey(t *testing.T) {
	k := RawcodetoKeychar(0)
	tt.Equal(t, "error", k)

	r := KeychartoRawcode("error")
	tt.Equal(t, 0, r)
}

func TestExtendedFunctionKeys(t *testing.T) {
	fmt.Println("Testing extended function keys F13-F24...")

	// Test that all extended function keys are properly mapped
	expectedKeys := map[string]uint16{
		"f13": 124,
		"f14": 125,
		"f15": 126,
		"f16": 127,
		"f17": 128,
		"f18": 129,
		"f19": 130,
		"f20": 131,
		"f21": 132,
		"f22": 133,
		"f23": 134,
		"f24": 135,
	}

	// Verify each extended function key exists and has correct keycode
	for key, expectedCode := range expectedKeys {
		actualCode, exists := Keycode[key]
		tt.True(t, exists, fmt.Sprintf("Key '%s' should exist in Keycode map", key))
		tt.Equal(t, expectedCode, actualCode, fmt.Sprintf("Key '%s' should have keycode %d", key, expectedCode))

		// Test reverse mapping if available
		if actualCode != 0 {
			fmt.Printf("✓ %s -> %d\n", key, actualCode)
		}
	}

	fmt.Printf("Successfully tested %d extended function keys\n", len(expectedKeys))
}

func TestExtendedKeysRange(t *testing.T) {
	fmt.Println("Testing extended function keys are in correct range...")

	// Test that F13-F24 keys are in the expected range (124-135)
	for i := 13; i <= 24; i++ {
		key := fmt.Sprintf("f%d", i)
		expectedCode := uint16(124 + (i - 13))

		actualCode, exists := Keycode[key]
		tt.True(t, exists, fmt.Sprintf("F%d should exist", i))
		tt.Equal(t, expectedCode, actualCode, fmt.Sprintf("F%d should have keycode %d", i, expectedCode))
	}
}

func TestStandardFunctionKeysStillWork(t *testing.T) {
	fmt.Println("Testing that standard function keys F1-F12 still work...")

	// Ensure we didn't break existing F1-F12 keys
	standardKeys := map[string]uint16{
		"f1":  59,
		"f2":  60,
		"f3":  61,
		"f4":  62,
		"f5":  63,
		"f6":  64,
		"f7":  65,
		"f8":  66,
		"f9":  67,
		"f10": 68,
		"f11": 69,
		"f12": 70,
	}

	for key, expectedCode := range standardKeys {
		actualCode, exists := Keycode[key]
		tt.True(t, exists, fmt.Sprintf("Standard key '%s' should still exist", key))
		tt.Equal(t, expectedCode, actualCode, fmt.Sprintf("Standard key '%s' should have keycode %d", key, expectedCode))
	}

	fmt.Printf("✓ All %d standard function keys verified\n", len(standardKeys))
}

func TestKeycodeMapIntegrity(t *testing.T) {
	fmt.Println("Testing keycode map integrity...")

	// Test that the Keycode map is not nil and contains expected entries
	tt.NotNil(t, Keycode, "Keycode map should not be nil")

	// Check that we have a reasonable number of keys (should be > 100)
	keyCount := len(Keycode)
	tt.True(t, keyCount > 100, fmt.Sprintf("Keycode map should have many keys, got %d", keyCount))

	// Test some essential keys are still there
	essentialKeys := []string{"a", "b", "c", "1", "2", "3", "space", "enter", "esc"}
	for _, key := range essentialKeys {
		_, exists := Keycode[key]
		tt.True(t, exists, fmt.Sprintf("Essential key '%s' should exist", key))
	}

	fmt.Printf("✓ Keycode map integrity verified with %d total keys\n", keyCount)
}

func TestF13SpecificFunctionality(t *testing.T) {
	fmt.Println("Testing F13 specific functionality...")

	// Test F13 keycode specifically
	f13Code, exists := Keycode["f13"]
	tt.True(t, exists, "F13 should exist in keycode map")
	tt.Equal(t, uint16(124), f13Code, "F13 should have keycode 124")

	// Test that F13 is different from other function keys
	f12Code := Keycode["f12"]
	f14Code := Keycode["f14"]

	tt.True(t, f13Code != f12Code, "F13 should have different keycode than F12")
	tt.True(t, f13Code != f14Code, "F13 should have different keycode than F14")

	// Test keycode conversion functions with F13
	keychar := RawcodetoKeychar(124)
	fmt.Printf("F13 rawcode 124 converts to keychar: '%s'\n", keychar)

	rawcode := KeychartoRawcode("f13")
	tt.Equal(t, uint16(124), rawcode, "F13 keychar should convert to rawcode 124")

	fmt.Println("✓ F13 specific functionality verified")
}

func TestF13HotkeyRegistration(t *testing.T) {
	fmt.Println("Testing F13 hotkey registration...")

	// Test that F13 can be registered as a hotkey
	// This tests the registration mechanism without actually starting the hook
	// (since we can't easily simulate key presses in a unit test)

	// Verify that F13 exists in the keycode map before registration
	_, exists := Keycode["f13"]
	tt.True(t, exists, "F13 must exist in keycode map for registration")

	// Test that the registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("F13 registration should not panic: %v", r)
		}
	}()

	// Simulate what happens during registration
	testCallback := func(e Event) {
		fmt.Printf("F13 callback triggered with event: Keycode=%d, Rawcode=%d\n",
			e.Keycode, e.Rawcode)
	}

	// Test single key registration
	fmt.Println("  Testing single F13 registration...")
	Register(KeyDown, []string{"f13"}, testCallback)

	// Test F13 with modifiers
	fmt.Println("  Testing F13 with Ctrl modifier...")
	Register(KeyDown, []string{"f13", "ctrl"}, testCallback)

	fmt.Println("  Testing F13 with Shift modifier...")
	Register(KeyDown, []string{"f13", "shift"}, testCallback)

	fmt.Println("  Testing F13 with Alt modifier...")
	Register(KeyDown, []string{"f13", "alt"}, testCallback)

	fmt.Println("✓ F13 hotkey registration completed successfully")
}
