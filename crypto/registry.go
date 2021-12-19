package crypto

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var (
	keyTypes = make(map[string]reflect.Type)
)

// UnmarshalPubKeyJSON decodes a JSON text as a PubKey implementation.
func UnmarshalPubKey(data []byte) (PubKey, error) {
	pk, err := unmarshalKeyType(data)
	if err != nil {
		return nil, err
	} else if pub, ok := pk.(PubKey); ok {
		return pub, nil
	}
	return nil, fmt.Errorf("type %T is not a PubKey", pk)
}

// UnmarshalPrivKeyJSON decodes a JSON text as a PrivKey implementation.
func UnmarshalPrivKey(data []byte) (PrivKey, error) {
	pk, err := unmarshalKeyType(data)
	if err != nil {
		return nil, err
	} else if priv, ok := pk.(PrivKey); ok {
		return priv, nil
	}
	return nil, fmt.Errorf("type %T is not a PrivKey", pk)
}

func unmarshalKeyType(data []byte) (interface{}, error) {
	var shim struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &shim); err != nil {
		return nil, err
	}
	ktype := newKeyValue(shim.Type)
	if ktype == nil {
		return nil, fmt.Errorf("unknown key type %q", shim.Type)
	} else if err := json.Unmarshal(shim.Value, ktype); err != nil {
		return nil, err
	}
	return reflect.ValueOf(ktype).Elem().Interface(), nil
}

func registerKeyType(name string, v reflect.Type) {
	if _, ok := keyTypes[name]; ok {
		panic("duplicate registration for: " + name)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	keyTypes[name] = v
}

// newKeyValue constructs a pointer to a freshly-allocated instance of a
// registered type, or returns nil if no such type is known.
func newKeyValue(name string) interface{} {
	t, ok := keyTypes[name]
	if !ok {
		return nil
	}
	return reflect.New(t).Interface()
}

// RegisterPubKeyType registers a public key implementation type under the
// given name.  It will panic if name is already registered.
// This should only be used at program initialization.
func RegisterPubKeyType(name string, v PubKey) {
	registerKeyType(name, reflect.TypeOf(v))
}

// RegisterPubKeyType registers a public key implementation type under the
// given name.  It will panic if name is already registered.
// This should only be used at program initialization.
func RegisterPrivKeyType(name string, v PrivKey) {
	registerKeyType(name, reflect.TypeOf(v))
}
