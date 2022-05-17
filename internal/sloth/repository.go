package sloth

//
// DB
//

var DB = map[int]*Sloth{1: NewSloth(1, "first", "none")}

type Repository interface {
	Add(sloth *Sloth) *Sloth
	Update(sloth *Sloth) *Sloth
	Delete(sloth *Sloth)
	Get(sloth *Sloth) *Sloth
	List() []*Sloth
}

type repository struct{}

func NewRepository() Repository { return new(repository) }

func (r *repository) Add(sloth *Sloth) *Sloth {
	sloth.SetID(len(DB) + 1)
	DB[sloth.GetID()] = sloth
	return sloth
}

func (r *repository) Delete(sloth *Sloth) {
	DB[sloth.GetID()] = nil
}

func (r *repository) Update(sloth *Sloth) *Sloth {
	DB[sloth.GetID()] = sloth
	return sloth
}

func (r *repository) Get(sloth *Sloth) *Sloth {
	return DB[sloth.GetID()]
}

func (r *repository) List() []*Sloth {
	var sloths []*Sloth

	for _, sloth := range DB {
		sloths = append(sloths, sloth)
	}

	return sloths
}
