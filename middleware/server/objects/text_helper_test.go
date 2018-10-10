package objects

import (
	"testing"
)

func TestUpperText(t *testing.T) {
	textHelper := TextHelper{}
	expectedText := "ABCDEFG123"
	actualText := textHelper.UpperText("abCdEfg123")
	if expectedText != actualText {
		t.Errorf("UpperText was incorrect, got: %s, want: %s.", actualText, expectedText)
	}
}

func TestLowerText(t *testing.T) {
	textHelper := TextHelper{}
	expectedText := "abcdefg123"
	actualText := textHelper.LowerText("abCdEfg123")
	if expectedText != actualText {
		t.Errorf("LowerText was incorrect, got: %s, want: %s.", actualText, expectedText)
	}
}

func TestInverText(t *testing.T) {
	textHelper := TextHelper{}
	expectedText := "321gfedcba"
	actualText := textHelper.InvertText("abcdefg123")
	if expectedText != actualText {
		t.Errorf("InvertText was incorrect, got: %s, want: %s.", actualText, expectedText)
	}
}
