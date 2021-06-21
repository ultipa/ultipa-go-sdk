package structs


type Table struct {
	Name string
	Headers []*Property
	Rows []*Row
}

type Row []interface{}

func NewTable() *Table {
	return &Table{
		Headers: []*Property{},
		Rows: []*Row{},
	}
}

func (t *Table) GetHeaders() []*Property{
	return t.Headers
}

func (t *Table) GetRows() []*Row {
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
