package rpmcompare

import (
	"bytes"
	"strings"
)

type RPMField []byte

func (rf RPMField) Compare(o RPMField) int {
	//FIXME handle ~
	return bytes.Compare(rf, o)
}

func (rf RPMField) EQ(o RPMField) bool {
	return rf.Compare(o) == 0
}

func (rf RPMField) GT(o RPMField) bool {
	return rf.Compare(o) > 1
}

func (rf RPMField) GTE(o RPMField) bool {
	return rf.Compare(o) >= 0
}

func (rf RPMField) LT(o RPMField) bool {
	return rf.Compare(o) < 0
}

func (rf RPMField) LTE(o RPMField) bool {
	return rf.Compare(o) <= 0
}

type RPMFields []RPMField

func newField(s string) RPMField {
	b := []byte(s)
	return RPMField(b)
}

func (rfs RPMFields) Compare(os RPMFields) int {
	for i, v1 := range rfs {
		// if rfs was longer it is larger
		if i >= len(os) {
			return 1
		}
		v2 := os[i]
		c := v1.Compare(v2)
		if c != 0 {
			return c
		}
	}

	// Same length they are equal
	if len(rfs) == len(os) {
		return 0
	}
	// rfs is shorts but has the same (subset of) values as os so it is less
	return -1
}

func (rfs RPMFields) EQ(os RPMFields) bool {
	return rfs.Compare(os) == 0
}

func (rfs RPMFields) GT(os RPMFields) bool {
	return rfs.Compare(os) > 0
}

func (rfs RPMFields) GTE(os RPMFields) bool {
	return rfs.Compare(os) >= 0
}

func (rfs RPMFields) LT(os RPMFields) bool {
	return rfs.Compare(os) < 0
}

func (rfs RPMFields) LTE(os RPMFields) bool {
	return rfs.Compare(os) <= 0
}

type RPMVersion struct {
	Version RPMFields
	Release RPMFields
}

func New(s string) RPMVersion {
	out := RPMVersion{}

	vr := strings.SplitN(s, "-", 2)
	vals := strings.Split(vr[0], ".")
	for _, val := range vals {
		rv := newField(val)
		out.Version = append(out.Version, rv)
	}
	if len(vr) == 1 {
		// no -, or nothing after -
		return out
	}
	vals = strings.Split(vr[1], ".")
	for _, val := range vals {
		rv := newField(val)
		out.Release = append(out.Release, rv)
	}
	return out
}

func (rv RPMVersion) Compare(o RPMVersion) int {
	comp := rv.Version.Compare(o.Version)
	if comp != 0 {
		return comp
	}
	return rv.Release.Compare(o.Release)
}

func (rv RPMVersion) EQ(o RPMVersion) bool {
	return rv.Compare(o) == 0
}

func (rv RPMVersion) GT(o RPMVersion) bool {
	return rv.Compare(o) > 0
}

func (rv RPMVersion) GTE(o RPMVersion) bool {
	return rv.Compare(o) >= 0
}

func (rv RPMVersion) LT(o RPMVersion) bool {
	return rv.Compare(o) < 0
}

func (rv RPMVersion) LTE(o RPMVersion) bool {
	return rv.Compare(o) <= 0
}
