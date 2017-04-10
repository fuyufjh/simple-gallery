package obs

type Service interface {
	List() ([]*Category, error)
}
