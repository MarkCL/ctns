package ctns

import "reflect"

// ConvertToNewType converts a struct to another struct with same tag
// if tagName is not passed into this function, it will convert T1 to T2 with forced transformation
// tagName accepts json or other tag name which will be used to match fields
// it will only take first tagName for conversion
func ConvertToNewType[T1, T2 any](obj T1, tagName ...string) T2 {
	sType := reflect.TypeOf(obj)
	sValue := reflect.ValueOf(obj)
	if sType.Kind() == reflect.Ptr {
		if sValue.IsNil() {
			return []T2{}[0]
		}
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	var d T2
	var d_exist_map map[string]bool
	var d_idx_map map[string]int
	if tagName != nil {
		d_exist_map, d_idx_map, d = getDestinationInfo[T2](tagName[0])
	}
	dType := reflect.TypeOf(d)
	dKind := dType.Kind()
	if dKind == reflect.Ptr {
		dType = dType.Elem()
		d = reflect.New(dType).Interface().(T2)
	} else {
		d = reflect.New(reflect.TypeOf([1]T2{})).Elem().Index(0).Interface().(T2)
	}

	if tagName == nil {
		return sValue.Convert(dType).Interface().(T2)
	}

	var dValue reflect.Value
	if dKind == reflect.Ptr {
		dValue = reflect.ValueOf(d).Elem()
	} else {
		dValue = reflect.ValueOf(&d).Elem()
	}
	setDestinationStructValue(sType, sValue, dType, dValue, d_exist_map, d_idx_map, tagName[0])
	return d
}
