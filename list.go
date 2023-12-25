/**
 * @Author: ray
 * @Description: this is a tool kit of list.
 * @File: list
 * @Version: 1.0.0
 * @Date: 2023/12/21 14:17
 */
package list-kit
import (
	"container/list"
)

type List = list.List

// GetAllElementPositive2Slice get all elements from list to slice order by positive
//
func GetAllElementPositive2Slice(l *List) []interface{} {
	if l.Len() == 0 {
		return nil
	}
	var s = make([]interface{}, l.Len())
	for n, e := 0, l.Front(); n < l.Len(); n, e = n+1, e.Next() {
		s[n] = e.Value
	}
	return s
}

// GetAllElementReverse2Slice get all elements from list to slice order by reverse
//
func GetAllElementReverse2Slice(l *List) interface{} {
	if l.Len() == 0 {
		return nil
	}
	var s = make([]interface{}, l.Len())
	for n, e := 0, l.Back(); n < l.Len(); n, e = n+1, e.Prev() {
		s[n] = e.Value
	}
	return s
}

// IndexOfList get value as a any type ([]Type) by index.
func IndexOfList(l *List, start, end int) (v any) {
	if start < 0 || end < 0 || end-start >= l.Len() || start-end > 0 {
		return nil
	}
	var e *list.Element
	var vs []any
	e = l.Front()
	for i := 0; i < start; i++ {
		if i == 0 && start == 0 {
			vs = append(vs, e.Value)
			break
		} else {
			e = e.Next()
		}
	}
	if start == end {
		return append(vs, e.Value)
	}
	for i := start; i < end; i++ {
		vs = append(vs, e.Value)
		e = e.Next()
	}
	return vs
}
