package example

type UserRepository struct{}

func (r *UserRepository) FindAll() []string {
	return []string{"alice", "bob", "charlie"}
}
