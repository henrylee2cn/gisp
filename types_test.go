package gisp

import (
	"reflect"
	"testing"
)

func TestTypeBool(t *testing.T) {
	var b = true
	if !reflect.DeepEqual(reflect.TypeOf(b), BOOL) {
		t.Fatalf("except %v equal string reflect type.", BOOL)
	}
}

func TestTypeString(t *testing.T) {
	var str = ""
	if !reflect.DeepEqual(reflect.TypeOf(str), STRING) {
		t.Fatalf("except %v equal string reflect type.", STRING)
	}
}

func TestTypeInt(t *testing.T) {
	var i Int
	if !reflect.DeepEqual(reflect.TypeOf(i), INT) {
		t.Fatalf("except %v equal Int reflect type.", INT)
	}
}

func TestTypeFloat(t *testing.T) {
	var f Float
	if !reflect.DeepEqual(reflect.TypeOf(f), FLOAT) {
		t.Fatalf("except %v equal Float reflect type.", FLOAT)
	}
}

func TestTypeAny(t *testing.T) {
	var it interface{}
	it = ""
	typ := reflect.TypeOf(&it).Elem()
	if !reflect.DeepEqual(typ, ANY) {
		t.Fatalf("except %v equal interface{} reflect type %v.", ANY, typ)
	}
}

func TestTypeAtom(t *testing.T) {
	var atom = Atom{"any", Type{reflect.TypeOf(0), true}}
	if !reflect.DeepEqual(reflect.TypeOf(atom), ATOM) {
		t.Fatalf("except %v equal Atom reflect type.", ATOM)
	}
}
