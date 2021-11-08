package main

import "testing"

func TestUtf8_OneCharacter_OneByte(t *testing.T) {
	if !CreateUtf8Checker()(0x24) {
		t.Fatalf("Incorrect result for 0x24")
	}
	if CreateUtf8Checker()(0x80) {
		t.Fatalf("Incorrect result for 0x80")
	}
}

func TestUtf8_FourCharacters_OneByte(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0x4F) || !checker(0x6C) || !checker(0x6C) || !checker(0x65) {
		t.Fatalf("Incorrect result for 0x4F 0x6C 0x6C 0x65")
	}
}

func TestUtf8_OneCharacter_TwoBytes(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0xC2) || !checker(0xA2) {
		t.Fatalf("Incorrect result for 0xC2 0xA2")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_1(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0xE0) || !checker(0xA4) || !checker(0xB9) {
		t.Fatalf("Incorrect result for 0xE0 0xA4 0xB9")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_2(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0xE2) || !checker(0x82) || !checker(0xAC) {
		t.Fatalf("Incorrect result for 0xE2 0x82 0xAC")
	}
}

func TestUtf8_OneCharacter_ThreeBytes_3(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0xED) || !checker(0x95) || !checker(0x9C) {
		t.Fatalf("Incorrect result for 0xED 0x95 0x9C")
	}
}

func TestUtf8_OneCharacter_FourBytes(t *testing.T) {
	checker := CreateUtf8Checker()
	if !checker(0xF0) || !checker(0x90) || !checker(0x8D) || !checker(0x88) {
		t.Fatalf("Incorrect result for 0xF0 0x90 0x8D 0x88")
	}
}
