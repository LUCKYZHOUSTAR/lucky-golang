package goext

//三元表达式函数

func IfInt8(condition bool, trueValue, falseValue int8) int8 {

	if condition {
		return trueValue
	} else {
		return falseValue
	}

}

func IfInt16(condition bool, trueValue, falseValue int16) int16 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfInt32(condition bool, trueValue, falseValue int32) int32 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfInt64(condition bool, trueValue, falseValue int64) int64 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfInt(condition bool, trueValue, falseValue int) int {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfByte(condition bool, trueValue, falseValue byte) byte {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfUint16(condition bool, trueValue, falseValue uint16) uint16 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfUint32(condition bool, trueValue, falseValue uint32) uint32 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfUint64(condition bool, trueValue, falseValue uint64) uint64 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfUint(condition bool, trueValue, falseValue uint) uint {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfFloat32(condition bool, trueValue, falseValue float32) float32 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfFloat64(condition bool, trueValue, falseValue float64) float64 {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfRune(condition bool, trueValue, falseValue rune) rune {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfString(condition bool, trueValue, falseValue string) string {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func IfError(condition bool, trueValue, falseValue error) error {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}
