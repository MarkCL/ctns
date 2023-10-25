package ctns

import "reflect"

// ConvertToNewStruct ToNewStruct converts a struct to another struct with same tag
// tagName accepts json or other tag name which will be used to match fields
func ConvertToNewStruct[T1, T2 any](obj T1, tagName string) T2 {
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
	dType := reflect.TypeOf(d)
	dKind := dType.Kind()
	if dKind == reflect.Ptr {
		dType = dType.Elem()
		d = reflect.New(dType).Interface().(T2)
	} else {
		d = reflect.New(reflect.TypeOf([1]T2{})).Elem().Index(0).Interface().(T2)
	}
	var dValue reflect.Value
	if dKind == reflect.Ptr {
		dValue = reflect.ValueOf(d).Elem()
	} else {
		dValue = reflect.ValueOf(&d).Elem()
	}
	d_exist_map := make(map[string]bool)
	d_idx_map := make(map[string]int)
	for i := 0; i < dType.NumField(); i++ {
		d_tag := dType.Field(i).Tag.Get(tagName)
		if d_tag != "" {
			d_idx_map[d_tag] = i
			d_exist_map[d_tag] = true
		}
	}
	for i := 0; i < sType.NumField(); i++ {
		s_tag := sType.Field(i).Tag.Get(tagName)
		if s_tag != "" && d_exist_map[s_tag] {
			if dType.Field(d_idx_map[s_tag]).Type != sType.Field(i).Type {
				continue
			}
			dField := dValue.Field(d_idx_map[s_tag])
			if dField.CanSet() {
				dField.Set(sValue.Field(i))
			}
		}
	}
	return d
}
