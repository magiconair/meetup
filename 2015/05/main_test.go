package main

import (
	"encoding/json"
	"reflect"
	"regexp"
	"testing"
	"time"
)

// a OMIT
func compare(a, b string) bool {
	return a == b
}

// A OMIT

// b OMIT
func TestCompareEmpty(t *testing.T) {
	if !compare("", "") {
		t.Fatal()
	}
}

func TestCompareNonEmpty(t *testing.T) {
	if !compare("a", "a") {
		t.Fatal()
	}
}

func TestCompareNotEqual(t *testing.T) {
	if compare("a", "b") {
		t.Fatal()
	}
}

// B OMIT

// c OMIT
type compareTest struct {
	a, b string
	eq   bool
}

var compareTests = []compareTest{
	{"", "", true},
	{"a", "a", true},
	{"a", "b", false},
}

func TestCompare1(t *testing.T) {
	for _, tt := range compareTests {
		if compare(tt.a, tt.b) != tt.eq {
			t.Errorf("compare(%q, %q) did not return %v", tt.a, tt.b, tt.eq)
		}
	}
}

// C OMIT

// d OMIT
func TestCompare2(t *testing.T) {
	var tests = []struct {
		a, b string
		eq   bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for i, tt := range tests {
		if compare(tt.a, tt.b) != tt.eq {
			t.Errorf("%d: compare(%q, %q) did not return %v", i, tt.a, tt.b, tt.eq)
		}
	}
}

// D OMIT

// e OMIT
func TestCompare3(t *testing.T) {
	var tests = []struct {
		a, b string
		eq   bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for i, tt := range tests {
		if got, want := compare(tt.a, tt.b), tt.eq; got != want {
			t.Errorf("%d: got %v want %v", i, got, want)
		}
	}
}

// E OMIT

// f OMIT
func compareJSON(a, b string) (bool, error) {
	var aa, bb map[string]interface{}
	if err := json.Unmarshal([]byte(a), &aa); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(b), &bb); err != nil {
		return false, err
	}
	return reflect.DeepEqual(aa, bb), nil
}

// F OMIT

// g1 OMIT
func TestCompareJSON1(t *testing.T) {
	var tests = []struct {
		a, b string
		eq   bool
		err  string
	}{
		{`{}`, `{}`, true, ""},
		{`{"a":1}`, `{"a":1}`, true, ""},
		{`{"a":1}`, `{"b":2}`, false, ""},
		{"a", "b", false, "invalid character"},
	}
	// G1 OMIT
	for i, tt := range tests {
		cmp, err := compareJSON(tt.a, tt.b)
		if got, want := err, tt.err; got != nil &&
			!regexp.MustCompile(want).MatchString(got.Error()) {
			t.Errorf("%d: got %q want %q", i, got, want)
			continue
		}
		if got, want := err, tt.err; got == nil && want != "" {
			t.Errorf("%d: got %q want %q", i, got, want)
		}
		if got, want := cmp, tt.eq; got != want {
			t.Errorf("%d: got %v want %v", i, got, want)
		}
	}
}

// G2 OMIT

// h OMIT
type Address struct {
	street, city, zip string
}

func NewAddress(street, city, zip string) Address {
	return Address{street, city, zip}
}

// H OMIT

// i1 OMIT
func TestNewAddress(t *testing.T) {
	var tests = []struct {
		street, city, zip string
		addr              Address
	}{
		{"s", "c", "z", Address{"s", "c", "z"}},
	}

	for i, tt := range tests {
		addr := NewAddress(tt.street, tt.city, tt.zip)

		if got, want := addr.street, tt.street; got != want {
			t.Errorf("%d: got %q want %q", i, got, want)
		}
	}
}

// I1 OMIT

// i2 OMIT
func TestNewAddress2(t *testing.T) {
	var tests = []struct {
		street, city, zip string
		addr              Address
	}{
		{"s", "c", "z", Address{"s", "c", "z"}},
	}

	for i, tt := range tests {
		addr := NewAddress(tt.street, tt.city, tt.zip)

		if got, want := addr.street, tt.street; got != want {
			t.Errorf("%d: got %q want %q", i, got, want)
		}
		if got, want := addr.city, tt.city; got != want {
			t.Errorf("%d: got %q want %q", i, got, want)
		}
		if got, want := addr.zip, tt.zip; got != want {
			t.Errorf("%d: got %q want %q", i, got, want)
		}
	}
}

// I2 OMIT

// i3 OMIT
func TestNewAddress3(t *testing.T) {
	var tests = []struct {
		street, city, zip string
		addr              Address
	}{
		{"s", "c", "z", Address{"s", "c", "z"}},
	}

	for i, tt := range tests {
		addr := NewAddress(tt.street, tt.city, tt.zip)
		ttAddr := Address{tt.street, tt.city, tt.zip}

		if got, want := addr, ttAddr; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %#v want %#v", i, got, want)
		}
	}
}

// I3 OMIT

// j1 OMIT
func mysqlNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TestMySQLNow(t *testing.T) {
	if got, want := mysqlNow(), time.Now().Format("2006-01-02 15:04:05"); got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

// J1 OMIT

// j2 OMIT
func mysqlNowNS() string {
	return time.Now().Format("2006-01-02 15:04:05.999999999")
}

func TestMySQLNowNS(t *testing.T) {
	// will this work?
	if got, want := mysqlNowNS(), time.Now().Format("2006-01-02 15:04:05.999999999"); got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

// J2 OMIT

// j3 OMIT

// stubbed out for testing
var timeNow = time.Now

func mysqlNowNS2() string {
	return timeNow().Format("2006-01-02 15:04:05.999999999")
}

func TestMySQLNowNS2(t *testing.T) {
	sometime := time.Date(2015, 05, 07, 00, 01, 02, 03, time.Local)
	timeNow = func() time.Time { return sometime }

	if got, want := mysqlNowNS2(), sometime.Format("2006-01-02 15:04:05.999999999"); got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

// J3 OMIT
