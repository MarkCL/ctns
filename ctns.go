package ctns

import (
	"errors"
	"reflect"
)

// ConvertToNewType converts a struct to another struct with same tag
// if tagName is not passed into this function, it will convert T1 to T2 with forced transformation
// tagName accepts json or other tag name which will be used to match fields
// it will only take first tagName for conversion
func ConvertToNewType[T1, T2 any](obj T1, tagName ...string) (T2, error) {
	sType := reflect.TypeOf(obj)
	sValue := reflect.ValueOf(obj)
	sKind := sType.Kind()
	if sKind == reflect.Ptr {
		if sValue.IsNil() {
			return []T2{}[0], errors.New("nil pointer")
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
		if tagName != nil {
			d = reflect.New(dType).Interface().(T2)
		}
	} else if tagName != nil {
		d = reflect.New(reflect.TypeOf([1]T2{})).Elem().Index(0).Interface().(T2)
	}

	if tagName == nil {
		if !sValue.CanConvert(dType) || (sKind == reflect.Ptr && dKind != reflect.Ptr) || (sKind != reflect.Ptr && dKind == reflect.Ptr) {
			return d, errors.New("can not convert to new type")
		}
		return sValue.Convert(dType).Interface().(T2), nil
	}

	var dValue reflect.Value
	if dKind == reflect.Ptr {
		dValue = reflect.ValueOf(d).Elem()
	} else {
		dValue = reflect.ValueOf(&d).Elem()
	}
	setDestinationStructValue(sType, sValue, dType, dValue, d_exist_map, d_idx_map, tagName[0])
	return d, nil
}
