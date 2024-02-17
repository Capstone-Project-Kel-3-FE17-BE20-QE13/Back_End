package service

import (
	"JobHuntz/features/favorit"
	"database/sql"
)

type FavService struct {
	Repo favorit.FavDataInterface
}

func New(Repo favorit.FavDataInterface) favorit.FavServiceInterface {
	return &FavService{
		Repo: Repo,
	}
}

func (uc *FavService) GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error) {

	result, err := uc.Repo.GetDataCompany(dbRaw, vacancyID)
	if err != nil {
		return favorit.DataCompanyCore{}, err
	}

	return result, nil
}

func (uc *FavService) CreateFavorit(input favorit.Core) (uint, error) {

	application, err := uc.Repo.CreateFavorit(input)

	if err != nil {
		return 0, err
	}

	return application, nil
}

func (uc *FavService) GetAllFavorit(userID uint) ([]favorit.Core, error) {
	results, err := uc.Repo.GetAllFavorit(userID)
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
