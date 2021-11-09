package checker

import "testing"

func TestUtf8_OneCharacter_OneByte_1(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0x24)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0x24")
	}
}

func TestUtf8_OneCharacter_OneByte_2(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0x80)
	if checker.Validates() {
		t.Fatalf("Incorrect result for 0x80")
	}
}

func TestUtf8_FourCharacters_OneByte(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0x4F)
	checker.CheckNext(0x6C)
	checker.CheckNext(0x6C)
	checker.CheckNext(0x65)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0x4F 0x6C 0x6C 0x65")
	}
}

func TestUtf8_OneCharacter_TwoBytes(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0xC2)
	checker.CheckNext(0xA2)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0xC2 0xA2")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_1(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0xE0)
	checker.CheckNext(0xA4)
	checker.CheckNext(0xB9)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0xE0 0xA4 0xB9")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_2(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0xE2)
	checker.CheckNext(0x82)
	checker.CheckNext(0xAC)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0xE2 0x82 0xAC")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_3(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0xED)
	checker.CheckNext(0x95)
	checker.CheckNext(0x9C)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0xED 0x95 0x9C")
	}
}

func TestUtf8_OneCharacter_FourBytes(t *testing.T) {
	checker := NewUtf8Checker()
	checker.CheckNext(0xF0)
	checker.CheckNext(0x90)
	checker.CheckNext(0x8D)
	checker.CheckNext(0x88)
	if !checker.Validates() {
		t.Fatalf("Incorrect result for 0xF0 0x90 0x8D 0x88")
	}
}
