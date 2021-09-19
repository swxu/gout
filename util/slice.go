package util

import "reflect"

func SliceElemEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if va.Kind() != reflect.Slice || vb.Kind() != reflect.Slice ||
		va.Type().Elem().Kind() != vb.Type().Elem().Kind() ||
		(!va.IsNil()) != (!vb.IsNil()) ||
		va.Len() != vb.Len() {
		return false
	}

	m := make(map[interface{}]int)
	n := va.Len()
	for i := 0; i < n; i++ {
		m[va.Index(i).Interface()]++
	}

	for i := 0; i < n; i++ {
		k := vb.Index(i).Interface()
		if _, found := m[k]; !found {
			return false
		}
		m[k]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func SliceContains(s interface{}, target interface{}) bool {
	if s == nil || target == nil {
		return false
	}

	vs := reflect.ValueOf(s)
	vt := reflect.ValueOf(target)
	if vs.Kind() != reflect.Slice || vs.Type().Elem().Kind() != vt.Kind() {
		return false
	}

	n := vs.Len()
	for i := 0; i < n; i++ {
		v := vs.Index(i).Interface()
		if reflect.DeepEqual(v, target) {
			return true
		}
	}

	return false
}
