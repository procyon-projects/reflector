package reflector

import "reflect"

type Slice interface {
	CanSet() bool
	Elem() Type
	Value() any
	Index(i int) any
	Len() int
	Append(values ...any)
}

type sliceType struct {
	elem Type

	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *sliceType) Name() string {
	return s.reflectType.Name()
}

func (s *sliceType) PackageName() string {
	return ""
}

func (s *sliceType) CanSet() bool {
	if s.reflectValue == nil {
		return false
	}

	return s.reflectValue.CanSet()
}

func (s *sliceType) Elem() Type {
	return s.elem
}

func (s *sliceType) Value() any {
	return s.reflectValue.Interface()
}

func (s *sliceType) Len() int {
	return s.reflectValue.Len()
}

func (s *sliceType) Index(i int) any {
	return s.reflectValue.Index(i)
}

func (s *sliceType) Append(values ...any) {
	if !s.CanSet() {
		return
	}

	for _, value := range values {
		*s.reflectValue = reflect.Append(*s.reflectValue, reflect.ValueOf(value))
	}
}

func (s *sliceType) HasValue() bool {
	return s.reflectValue != nil
}

func (s *sliceType) Parent() Type {
	return s.parent
}

func (s *sliceType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *sliceType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *sliceType) Instantiate() Value {
	return &value{
		reflect.New(s.reflectType),
	}
}
