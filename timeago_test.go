package timeago

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	testTime := New(time.Now().Add(-time.Minute*60*24*365*4 - time.Minute*60*10 - time.Second*10))
	if testTime.WithLocale("ru").Format() != "4 года назад" {
		t.Errorf("%s doesn't match %s", testTime.Format(), "4 года назад")
	}
	if testTime.WithLocale("en").Format() != "4 years ago" {
		t.Errorf("%s doesn't match %s", testTime.Format(), "4 years ago")
	}
	RegisterLocale("custom", ruLocale)
	if testTime.WithLocale("custom").Format() != "4 года назад" {
		t.Errorf("%s doesn't match %s", testTime.Format(), "4 года назад")
	}
}
