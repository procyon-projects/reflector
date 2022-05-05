package reflector

import (
	"errors"
	"reflect"
	"strings"
)

type Interface interface {
	Type
	Elem() Type
	Methods() []Method
	NumMethod() int
}

type interfaceType struct {
	parent       Type
	reflectType  reflect.Type
	reflectValue *reflect.Value
}

func (i *interfaceType) Name() string {
	if i.reflectType.Name() == "" {
		return "any"
	}

	return i.reflectType.Name()
}

func (i *interfaceType) PackageName() string {
	name := i.reflectType.PkgPath()
	slashLastIndex := strings.LastIndex(name, "/")

	if slashLastIndex != -1 {
		name = name[slashLastIndex+1:]
	}

	return name
}

func (i *interfaceType) PackagePath() string {
	return i.reflectType.PkgPath()
}

func (i *interfaceType) CanSet() bool {
	if i.reflectValue == nil {
		return false
	}

	return i.reflectValue.CanSet()
}

func (i *interfaceType) HasValue() bool {
	return i.reflectValue != nil
}

func (i *interfaceType) Value() (any, error) {
	if i.reflectValue == nil {
		return "", errors.New("value reference is nil")
	}

	return i.reflectValue.Interface(), nil
}

func (i *interfaceType) SetValue(val any) error {
	if !i.CanSet() {
		return errors.New("value cannot be set")
	}

	i.reflectValue.Set(reflect.ValueOf(val))
	return nil
}

func (i *interfaceType) Parent() Type {
	return i.parent
}

func (i *interfaceType) ReflectType() reflect.Type {
	return i.reflectType
}

func (i *interfaceType) ReflectValue() *reflect.Value {
	return i.reflectValue
}

func (i *interfaceType) Compare(another Type) bool {
	if another == nil {
		return false
	}

	return i.reflectType == another.ReflectType()
}

func (i *interfaceType) IsInstantiable() bool {
	return false
}

func (i *interfaceType) Instantiate() (Value, error) {
	return nil, errors.New("interfaces are not instantiable")
}

func (i *interfaceType) Elem() Type {
	return nil
}

func (i *interfaceType) Methods() []Method {
	functions := make([]Method, 0)
	numMethod := i.reflectType.NumMethod()

	for index := 0; index < numMethod; index++ {
		function := i.reflectType.Method(index)
		functions = append(functions, &methodType{
			name:        function.Name,
			pkgPath:     function.PkgPath,
			isExported:  function.IsExported(),
			reflectType: function.Type,
		})
	}

	return functions
}

func (i *interfaceType) NumMethod() int {
	return i.reflectType.NumMethod()
}
