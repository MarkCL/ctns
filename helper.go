package ctns

import "reflect"

func getDestinationInfo[T2 any](tagName string) (d_exist_map map[string]bool, d_idx_map map[string]int, d T2) {
	// var d T2
	dType := reflect.TypeOf(d)
	dKind := dType.Kind()
	if dKind == reflect.Ptr {
		dType = dType.Elem()
		d = reflect.New(dType).Interface().(T2)
	} else {
		d = reflect.New(reflect.TypeOf([1]T2{})).Elem().Index(0).Interface().(T2)
	}
	d_exist_map = make(map[string]bool)
	d_idx_map = make(map[string]int)
	for i := 0; i < dType.NumField(); i++ {
		d_tag := dType.Field(i).Tag.Get(tagName)
		if d_tag != "" {
			d_idx_map[d_tag] = i
			d_exist_map[d_tag] = true
		}
	}
	return d_exist_map, d_idx_map, d
}

func setDestinationStructValue(sType reflect.Type, sValue reflect.Value, dType reflect.Type, dValue reflect.Value, d_exist_map map[string]bool, d_idx_map map[string]int, tagName string) {
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
}
