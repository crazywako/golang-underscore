package underscore

import (
	"reflect"
)

// Find is 根据断言获取元素
func Find(source, predicate, match interface{}) {
	matchRV := reflect.ValueOf(match)
	if matchRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var ok bool
	each(source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		ok = resRV.Bool()
		if ok {
			matchRV.Elem().Set(valueRV)
		}
		return ok
	})
}

// FindBy is 根据属性值获取元素
func FindBy(source interface{}, properties map[string]interface{}, match interface{}) {
	Find(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	}, match)
}

func (m *query) Find(predicate interface{}) IQuery {
	Find(m.Source, predicate, &m.Source)
	return m
}

func (m *query) FindBy(properties map[string]interface{}) IQuery {
	FindBy(m.Source, properties, &m.Source)
	return m
}
