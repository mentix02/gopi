package gopi

type apply func(interface{}) interface{}

// Map is similar to Python's map() builtin.
func Map(fn apply, data []interface{}) []interface{} {
	d := make([]interface{}, len(data))
	copy(d, data)
	for index, item := range d {
		d[index] = fn(item)
	}
	return d
}


// SelfMap is similar to Map but instead of returning
// a new slice of the data provided, it directly changes
// the data slice provided.
func SelfMap(fn apply, data []interface{}) {
	for index, item := range data {
		data[index] = fn(item)
	}
}
