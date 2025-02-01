package in_memory

type Repo struct {
	data map[string]string
	//mu   sync.Mutex
}

func NewRepo() *Repo {
	return &Repo{
		data: make(map[string]string)}
}
