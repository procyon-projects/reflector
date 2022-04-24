package reflector

func IsPointer(typ Type) bool {
	_, ok := typ.(*pointer)
	return ok
}

func ToPointer(typ Type) (Pointer, bool) {
	if ptr, ok := typ.(*pointer); ok {
		return ptr, true
	}

	return nil, false
}

func IsStruct(typ Type) bool {
	_, ok := typ.(Struct)
	return ok
}

func ToStruct(typ Type) (Struct, bool) {
	if structType, ok := typ.(Struct); ok {
		return structType, true
	}

	return nil, false
}

func IsInterface(typ Type) bool {
	_, ok := typ.(*interfaceType)
	return ok
}

func ToInterface(typ Type) (Interface, bool) {
	if interfaceType, ok := typ.(*interfaceType); ok {
		return interfaceType, true
	}

	return nil, false
}

func IsFunction(typ Type) bool {
	_, ok := typ.(Function)
	return ok
}

func ToFunction(typ Type) (Function, bool) {
	if functionType, ok := typ.(Function); ok {
		return functionType, true
	}

	return nil, false
}

func IsArray(typ Type) bool {
	_, ok := typ.(Array)
	return ok
}

func ToArray(typ Type) (Array, bool) {
	if arrayType, ok := typ.(Array); ok {
		return arrayType, true
	}

	return nil, false
}

func IsSlice(typ Type) bool {
	_, ok := typ.(Slice)
	return ok
}

func ToSlice(typ Type) (Slice, bool) {
	if sliceType, ok := typ.(Slice); ok {
		return sliceType, true
	}

	return nil, false
}

func IsMap(typ Type) bool {
	_, ok := typ.(Map)
	return ok
}

func ToMap(typ Type) (Map, bool) {
	if mapType, ok := typ.(Map); ok {
		return mapType, true
	}

	return nil, false
}

func IsString(typ Type) bool {
	_, ok := typ.(String)
	return ok
}

func ToString(typ Type) (String, bool) {
	if stringType, ok := typ.(String); ok {
		return stringType, true
	}

	return nil, false
}

func IsBoolean(typ Type) bool {
	_, ok := typ.(Boolean)
	return ok
}

func ToBoolean(typ Type) (Boolean, bool) {
	if stringType, ok := typ.(Boolean); ok {
		return stringType, true
	}

	return nil, false
}

func IsInteger(typ Type) bool {
	return IsUnsignedInteger(typ) || IsSignedInteger(typ)
}

func IsSignedInteger(typ Type) bool {
	_, ok := typ.(SignedInteger)
	return ok
}

func ToSignedInteger(typ Type) (SignedInteger, bool) {
	if signedIntegerType, ok := typ.(SignedInteger); ok {
		return signedIntegerType, true
	}

	return nil, false
}

func IsUnsignedInteger(typ Type) bool {
	_, ok := typ.(UnsignedInteger)
	return ok
}

func ToUnsignedInteger(typ Type) (UnsignedInteger, bool) {
	if unsignedIntegerType, ok := typ.(UnsignedInteger); ok {
		return unsignedIntegerType, true
	}

	return nil, false
}

func IsFloat(typ Type) bool {
	_, ok := typ.(Float)
	return ok
}

func ToFloat(typ Type) (Float, bool) {
	if floatType, ok := typ.(Float); ok {
		return floatType, true
	}

	return nil, false
}

func IsComplex(typ Type) bool {
	_, ok := typ.(Complex)
	return ok
}

func ToComplex(typ Type) (Complex, bool) {
	if complexType, ok := typ.(Complex); ok {
		return complexType, true
	}

	return nil, false
}

func IsNumber(typ Type) bool {
	return IsInteger(typ) || IsFloat(typ) || IsComplex(typ)
}

func IsBasic(typ Type) bool {
	return IsBoolean(typ) || IsString(typ) || IsNumber(typ)
}

func IsInstantiable(typ Type) bool {
	_, ok := typ.(Instantiable)
	return ok
}

func ToInstantiable(typ Type) (Instantiable, bool) {
	if instantiableType, ok := typ.(Instantiable); ok {
		return instantiableType, true
	}

	return nil, false
}
