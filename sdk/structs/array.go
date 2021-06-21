package structs


type Array struct {
	Name string
	Rows []*Row
}

func NewArray() *Array {
	return &Array{
		Rows: []*Row{},
	}
}