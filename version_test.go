package rpmcompare

import (
	"testing"
)

func b(s string) []byte {
	return []byte(s)
}

var (
	fieldTests = []struct {
		s1       string
		s2       string
		expected int
	}{
		{
			s1:       "1",
			s2:       "2",
			expected: -1,
		},
		{
			s1:       "a",
			s2:       "a",
			expected: 0,
		},
		{
			s1:       "a1a",
			s2:       "a2a",
			expected: -1,
		},
		{
			s1:       "A",
			s2:       "a",
			expected: -1,
		},
	}
)

func TestRPMFieldCompare(t *testing.T) {
	for testNum, test := range fieldTests {
		v1 := RPMField(b(test.s1))
		v2 := RPMField(b(test.s2))
		comp := v1.Compare(v2)
		if comp != test.expected {
			t.Errorf("%d: comp=%v expected=%d", testNum, comp, test.expected)
		}
		res := v1.EQ(v2)
		if res != (test.expected == 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.LT(v2)
		if res != (test.expected < 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.LTE(v2)
		if res != (test.expected <= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.GT(v2)
		if res != (test.expected > 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.GTE(v2)
		if res != (test.expected >= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
	}
}

var (
	fieldsTests = []struct {
		s1       []string
		s2       []string
		expected int
	}{
		{
			s1:       []string{},
			s2:       []string{},
			expected: 0,
		},
		{
			s1:       []string{"3", "10"},
			s2:       []string{"3", "10"},
			expected: 0,
		},
		{
			s1:       []string{"3", "10", "0"},
			s2:       []string{"3", "10", "1"},
			expected: -1,
		},
		{
			s1:       []string{"3", "10", "0"},
			s2:       []string{"3", "10"},
			expected: 1,
		},
		{
			s1:       []string{"3", "10"},
			s2:       []string{"3", "10", "0"},
			expected: -1,
		},
	}
)

func fields(in []string) RPMFields {
	out := RPMFields{}
	for _, v := range in {
		out = append(out, b(v))
	}
	return out
}

func TestRPMFields(t *testing.T) {
	for testNum, test := range fieldsTests {
		f1 := fields(test.s1)
		f2 := fields(test.s2)
		comp := f1.Compare(f2)
		if comp != test.expected {
			t.Errorf("%d: res=%v expected=%d", testNum, comp, test.expected)
		}
		res := f1.EQ(f2)
		if res != (test.expected == 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = f1.LT(f2)
		if res != (test.expected < 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = f1.LTE(f2)
		if res != (test.expected <= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = f1.GT(f2)
		if res != (test.expected > 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = f1.GTE(f2)
		if res != (test.expected >= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
	}
}

var (
	versionTests = []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		{
			name:     "exact same",
			s1:       "3.10.0-33.22.1.el7",
			s2:       "3.10.0-33.22.1.el7",
			expected: 0,
		},
		{
			name:     "exact same, no -",
			s1:       "3.10.0",
			s2:       "3.10.0",
			expected: 0,
		},
		{
			name:     "s1 less all fields",
			s1:       "3.10.0-33.22.1.el7",
			s2:       "3.10.0-33.22.2.el7",
			expected: -1,
		},
		{
			name:     "s1 greater all fields",
			s1:       "3.10.0-33.22.3.el7",
			s2:       "3.10.0-33.22.2.el7",
			expected: 1,
		},
		{
			name:     "s1 greater, s1 nothing after -",
			s1:       "3.10.1",
			s2:       "3.10.0-33.22.2.el7",
			expected: 1,
		},
		{
			name:     "s2 greater, s1 nothing after -",
			s1:       "3.10.0",
			s2:       "3.10.1-33.22.2.el7",
			expected: -1,
		},
		{
			name:     "s1 greater, s2 nothing after -",
			s1:       "3.10.0-33.22.2.el7",
			s2:       "3.10.0",
			expected: 1,
		},
		{
			name:     "s2 greater, s2 nothing after -",
			s1:       "3.10.0-33.22.2.el7",
			s2:       "3.10.1",
			expected: -1,
		},
		{
			name:     "s1 greater, s1 fewer before -",
			s1:       "3.11-33.22.2.el7",
			s2:       "3.10.1-33.22.2.el7",
			expected: 1,
		},
		{
			name:     "capital letters",
			s1:       "1.2.3.a1a",
			s2:       "1.2.3.a1A",
			expected: 1,
		},
		{
			name:     "capital letters after -",
			s1:       "1.2.3-a1a",
			s2:       "1.2.3-a1A",
			expected: 1,
		},
	}
)

func TestRPMVersionCompare(t *testing.T) {
	for testNum, test := range versionTests {
		v1 := New(test.s1)
		v2 := New(test.s2)
		comp := v1.Compare(v2)
		if comp != test.expected {
			t.Errorf("%d:%q comp=%d: expected=%d", testNum, test.name, comp, test.expected)
		}
		res := v1.EQ(v2)
		if res != (test.expected == 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.LT(v2)
		if res != (test.expected < 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.LTE(v2)
		if res != (test.expected <= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.GT(v2)
		if res != (test.expected > 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
		res = v1.GTE(v2)
		if res != (test.expected >= 0) {
			t.Errorf("%d: res=%v expected=%d", testNum, res, test.expected)
		}
	}
}
