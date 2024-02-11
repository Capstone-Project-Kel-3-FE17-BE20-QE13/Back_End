package service

import (
	"JobHuntz/features/favorit"
)

type FavService struct {
	Repo favorit.FavDataInterface
}

func New(Repo favorit.FavDataInterface) favorit.FavServiceInterface {
	return &FavService{
		Repo: Repo,
	}
}
func (uc *FavService) CreateFavorit(input favorit.Core) (uint, error) {

	application, err := uc.Repo.CreateFavorit(input)

	if err != nil {
		return 0, err
	}

	return application, nil
}

func (uc *FavService) GetAllFavorit() ([]favorit.Core, error) {
	results, err := uc.Repo.GetAllFavorit()
	return results, err
}

// func (uc *FavService) DeleteFavById(JobId uint) error {
// 	err := uc.Repo.DeleteFavById(JobId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (uc *FavService) DeleteFavById(input []favorit.Core, id int) error {
	err := uc.Repo.DeleteFavById(input, id)
	return err
}