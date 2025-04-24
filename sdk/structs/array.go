package structs

type Array struct {
	Name string
	Rows []*Value
}

func NewArray() *Array {
	return &Array{
		Rows: []*Value{},
	}
}
