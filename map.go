package reflector

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Entry interface {
	Key() any
	Value() any
}

type entry struct {
	key any
	val any
}

func (e entry) Key() any {
	return e.key
}

func (e entry) Value() any {
	return e.val
}

type Map interface {
	Type
	Key() Type
	Elem() Type
	Len() (int, error)
	Contains(key any) (bool, error)
	Get(key any) (any, error)
	Put(key any, val any) error
	Delete(key any) error
	Clear() error
	KeySet() ([]any, error)
	ValueSet() ([]any, error)
	EntrySet() ([]Entry, error)
}

type mapType struct {
	key  Type
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (m *mapType) Name() string {
	var builder strings.Builder
	builder.WriteString("map[")
	builder.WriteString(m.key.Name())
	builder.WriteString("]")
	builder.WriteString(m.elem.Name())
	return builder.String()
}

func (m *mapType) PackageName() string {
	return ""
}

func (m *mapType) PackagePath() string {
	return ""
}

func (m *mapType) CanSet() bool {
	if m.reflectValue == nil {
		return false
	}

	return m.reflectValue.CanSet()
}

func (m *mapType) HasValue() bool {
	return m.reflectValue != nil
}

func (m *mapType) Value() (any, error) {
	if m.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return m.reflectValue.Interface(), nil
}

func (m *mapType) SetValue(val any) error {
	if !m.CanSet() {
		return errors.New("value cannot be set")
	}

	m.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (m *mapType) Parent() Type {
	return m.parent
}

func (m *mapType) ReflectType() reflect.Type {
	return m.reflectType
}

func (m *mapType) ReflectValue() *reflect.Value {
	return m.reflectValue
}

func (m *mapType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return m.reflectType == another.ReflectType()
}

func (m *mapType) IsInstantiable() bool {
	return true
}

func (m *mapType) Instantiate() (Value, error) {
	ptr := reflect.New(m.reflectType).Interface()
	emptyMap := reflect.MakeMapWithSize(reflect.MapOf(m.key.ReflectType(), m.elem.ReflectType()), 0)
	reflect.ValueOf(ptr).Elem().Set(emptyMap)
	return &value{
		reflect.ValueOf(ptr),
	}, nil
}

func (m *mapType) CanConvert(typ Type) bool {
	if typ == nil {
		return false
	}

	if m.reflectValue == nil {
		return m.reflectType.ConvertibleTo(typ.ReflectType())
	}

	return m.reflectValue.CanConvert(typ.ReflectType())
}

func (m *mapType) Convert(typ Type) (Value, error) {
	if typ == nil {
		return nil, errors.New("typ should not be nil")
	}

	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	if !m.CanConvert(typ) {
		return nil, errors.New("type is not valid")
	}

	val := m.reflectValue.Convert(typ.ReflectType())

	return &value{
		val,
	}, nil
}

func (m *mapType) Key() Type {
	return m.key
}

func (m *mapType) Elem() Type {
	return m.elem
}

func (m *mapType) Len() (int, error) {
	if m.reflectValue == nil {
		return -1, errors.New("value reference is nil")
	}

	return m.reflectValue.Len(), nil
}

func (m *mapType) KeySet() ([]any, error) {
	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	keySet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		keySet = append(keySet, key.Interface())
	}

	return keySet, nil
}

func (m *mapType) ValueSet() ([]any, error) {
	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	valueSet := make([]any, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		value := m.reflectValue.MapIndex(key)
		valueSet = append(valueSet, value.Interface())
	}

	return valueSet, nil
}

func (m *mapType) EntrySet() ([]Entry, error) {
	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	valueSet := make([]Entry, 0)
	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		value := m.reflectValue.MapIndex(key)
		valueSet = append(valueSet, entry{key: key.Interface(), val: value.Interface()})
	}

	return valueSet, nil
}

func (m *mapType) Contains(key any) (bool, error) {
	_, err := m.Get(key)

	if err != nil {
		return false, err
	}

	return true, err
}

func (m *mapType) Get(key any) (any, error) {
	if m.reflectValue == nil {
		return nil, errors.New("value reference is nil")
	}

	val := m.reflectValue.MapIndex(reflect.ValueOf(key))

	if val.Kind() == reflect.Invalid {
		return nil, fmt.Errorf("element with key '%v' does not exist", key)
	}

	return val.Interface(), nil
}

func (m *mapType) Put(key any, val any) error {
	if m.reflectValue == nil {
		return errors.New("value reference is nil")
	}

	m.reflectValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	return nil
}

func (m *mapType) Delete(key any) error {
	if m.reflectValue == nil {
		return errors.New("value reference is nil")
	}

	m.reflectValue.SetMapIndex(reflect.ValueOf(key), reflect.Value{})
	return nil
}

func (m *mapType) Clear() error {
	if m.reflectValue == nil {
		return errors.New("value reference is nil")
	}

	keys := m.reflectValue.MapKeys()

	for _, key := range keys {
		m.reflectValue.SetMapIndex(key, reflect.Value{})
	}
	return nil
}
