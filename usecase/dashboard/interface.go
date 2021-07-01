package dashboard

//Reader interface
type Reader interface {
	Get(u, p, groupId string) (*map[string]interface{}, error)
}

//Writer user writer
type Writer interface {
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetDashboardData(u, p, groupId string) (*map[string]interface{}, error)
}
