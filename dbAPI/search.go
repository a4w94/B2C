package dbAPI

type Search struct {
	Query []Condition
}
type Condition struct {
	Colum   string
	Compare string
	Value   interface{}
}


