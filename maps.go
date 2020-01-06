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
