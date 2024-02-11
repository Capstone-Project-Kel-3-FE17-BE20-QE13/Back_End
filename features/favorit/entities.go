package favorit

type Core struct {
	JobseekerID uint
	JobId       uint
	Name        string
}

type FavDataInterface interface {
	CreateFavorit(input Core) (uint, error)
	GetAllFavorit() ([]Core, error)
	// DeleteFavById(JobId uint) error
	DeleteFavById(input []Core, id int) error
}

type FavServiceInterface interface {
	CreateFavorit(input Core) (uint, error)
	GetAllFavorit() ([]Core, error)
	// DeleteFavById(JobId uint) error
	DeleteFavById(input []Core, id int) error
}
