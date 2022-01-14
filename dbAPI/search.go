package dbAPI

type Search struct {
	Query []Condition
}
type Condition struct {
	Colum   string
	Compare string
	Value   interface{}
}

func (d DBinfo) SearchByQuery(q Search) {

}
