package entiry

type People struct {
	Id   int
	Name string
	Sex  string
}

func NewPeople(id int, name, sex string) *People {
	return &People{id, name, sex}
}

func (p *People) SetId(id int) {
	p.Id = id
}

func (p People) GetId() (id int, people People) {
	return p.Id, p
}

type Work interface {
	Remark()
}
