package structs

type Table struct {
	Name    string
	Headers []*Property
	Rows    []*Value
}

type Value []interface{}

func NewTable() *Table {
	return &Table{
		Headers: []*Property{},
		Rows:    []*Value{},
	}
}

func (t *Table) GetHeaders() []*Property {
	return t.Headers
}

func (t *Table) GetRows() []*Value {
	return t.Rows
}

func (t *Table) ToKV() []*Values {

	var values []*Values

	for _, row := range t.GetRows() {
		v := NewValues()
		for i, header := range t.GetHeaders() {
			f := (*row)[i]
			v.Set(header.Name, f)
		}

		values = append(values, v)
	}

	return values
}
