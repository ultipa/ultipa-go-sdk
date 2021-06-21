package structs


type Attr struct {
	Name string
	Rows Row
}

func NewAttr() *Attr {
	return &Attr{
		Rows: Row{},
	}
}
