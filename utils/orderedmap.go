package utils

// OrderedMap implements ordered map data structure.
type OrderedMap struct {
	keys []interface{}
	m map[interface{}]interface{}
}

// NewOrderedMap function returns a pointer to the initialized OrderedMap.
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keys: make([]interface{}, 0),
		m: make(map[interface{}]interface{}),
	}
}

// Keys method returns internal ordered keys.
func (p OrderedMap) Keys() []interface{} {
	return p.keys
}

// RawMap method returns internal unordered map.
func (p OrderedMap) RawMap() map[interface{}]interface{} {
	return p.m
}

// Val method returns the value if key exists, and returns nil otherwise.
func (p OrderedMap) Val(k interface{}) interface{} {
	v, e := p.m[k]
	if e {
		return v
	} else {
		return nil
	}
}

// Add method adds (k,v) pair to the ordered map, and if key(=k) already exist, v will overwrite current value.
func (p OrderedMap) Add(k interface{}, v interface{}) {
	_, e := p.m[k]
	if e {
		p.m[k] = v
	} else {
		p.keys = append(p.keys, k)
		p.m[k] = v
	}
}

// Len method returns size of the ordered map.
func (p OrderedMap) Len() int {
	return len(p.keys)
}
