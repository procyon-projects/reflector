package reflector

import "reflect"

type Struct interface {
	Type
	Instantiable
	Fields() []Field
	NumField() int
	Methods() []Function
	NumMethod() int
	Implements(i Interface) bool
}

type structType struct {
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (s *structType) HasReference() bool {
	return s.reflectValue != nil
}

func (s *structType) ReflectType() reflect.Type {
	return s.reflectType
}

func (s *structType) ReflectValue() *reflect.Value {
	return s.reflectValue
}

func (s *structType) NumField() int {
	return s.reflectType.NumField()
}

func (s *structType) Fields() []Field {
	fields := make([]Field, 0)

	numField := s.reflectType.NumField()
	for i := 0; i < numField; i++ {
		f := s.reflectType.Field(i)
		if f.Type != nil {

		}
		x := f.Type.Name()
		if x != "" {

		}
		fields = append(fields, &field{})
	}
	return fields
}

func (s *structType) Methods() []Function {
	functions := make([]Function, 0)

	numMethod := s.reflectType.NumMethod()
	for i := 0; i < numMethod; i++ {
		f := s.reflectType.Field(i)
		if f.Type != nil {

		}
		functions = append(functions, &functionType{})
	}

	return functions
}

func (s *structType) NumMethod() int {
	return s.reflectType.NumMethod()
}

func (s *structType) Implements(i Interface) bool {
	return false
}

func (s *structType) Instantiate() any {
	return reflect.New(s.reflectType).Interface()
}
