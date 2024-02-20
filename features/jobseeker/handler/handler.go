package handler

import (
	config "JobHuntz/app/configs"
	"JobHuntz/app/database"
	"JobHuntz/app/middlewares"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JobseekerHandler struct {
	jobseekerService jobseeker.JobseekerServiceInterface
}

func New(service jobseeker.JobseekerServiceInterface) *JobseekerHandler {
	return &JobseekerHandler{
		jobseekerService: service,
	}
}

func (handler *JobseekerHandler) GetByIdJobSeeker(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errGetByID := handler.jobseekerService.GetByIdJobSeeker(seekerID)
	if errGetByID != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errGetByID.Error(), nil))
	}

	res := CoreResponById(*result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get my profile", res))
}
func (handler *JobseekerHandler) GetjobseekerByCompany(c echo.Context) error {
	jobseekerID := c.QueryParam("jobseeker_id")

	jobseekerID_int, errConv := strconv.Atoi(jobseekerID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobseekerService.GetjobseekerByCompany(uint(jobseekerID_int))
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	jobseekerResponse := CoreResponById(*result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get detail license", jobseekerResponse))
}
func (handler *JobseekerHandler) RegisterJobseeker(c echo.Context) error {
	newSeeker := JobseekerRegistRequest{}

	errBind := c.Bind(&newSeeker)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	registCore := RequestJobseekerRegistToCore(newSeeker)

	errVal := handler.jobseekerService.RegisterValidation(registCore)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, errVal.Error(), nil))
	}

	errCreate := handler.jobseekerService.Register(registCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully registered", nil))
}

func (handler *JobseekerHandler) LoginJobseeker(c echo.Context) error {
	newLogin := JobseekerRequest{}
	errBind := c.Bind(&newLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	resLogin, token, err := handler.jobseekerService.Login(newLogin.Email, newLogin.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	loginResponse := CoreJobseekerToResponseLogin(resLogin, token)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully login", loginResponse))
}

func (handler *JobseekerHandler) UpdateJobseeker(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newUpdate := JobseekerUpdateRequest{}

	newBanners, _ := c.FormFile("banners")
	if newBanners != nil {
		newBannersURL, errURL := handler.jobseekerService.Photo(newBanners)
		if errURL != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot get the photo's url", nil))
		}
		newUpdate.Banners = newBannersURL.SecureURL
	}

	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("isi newUpdate: ", newUpdate)

	newUpdateCore := RequestJobseekerUpdateToCore(newUpdate)

	errVal := handler.jobseekerService.UpdateValidation(newUpdateCore)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, errVal.Error(), nil))
	}

	err := handler.jobseekerService.UpdateProfile(seekerID, newUpdateCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update profile", nil))
}

func (handler *JobseekerHandler) CreateCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	inputCV, errRead := c.FormFile("cv_file")
	if errRead != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot read file", nil))
	}

	responURL, errURL := handler.jobseekerService.PDF(inputCV)
	if errURL != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot get resp", nil))
	}

	cfg := config.InitConfig()
	dbRaw := database.InitRawSql(cfg)

	count, _ := handler.jobseekerService.CountCV(dbRaw, seekerID)

	newCV := CVRequest{}
	newCV.JobseekerID = seekerID
	newCV.CV_file = responURL.Location

	errBind := c.Bind(&newCV)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	cvCore := RequestCVToCore(newCV)

	errCreate := handler.jobseekerService.AddCV(cvCore, count)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully upload cv", nil))
}

func (handler *JobseekerHandler) GetCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errFirst := handler.jobseekerService.ReadCV(seekerID)
	if errFirst != nil {
		return c.JSON(http.StatusNotFound, responses.WebResponse(http.StatusNotFound, "cannot find data "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get cv", result))
}

func (handler *JobseekerHandler) UpdateCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newCV := CVRequest{}
	newCV.JobseekerID = seekerID
	oldCV, _ := c.FormFile("cv_file")
	if oldCV != nil {
		responURL, errResp := handler.jobseekerService.PDF(oldCV)
		if errResp != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error get resp "+errResp.Error(), nil))
		}
		newCV.CV_file = responURL.Location
	}

	errBind := c.Bind(&newCV)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	newCVCore := RequestCVToCore(newCV)

	fmt.Println("isi update: ", newCVCore)

	errUpdate := handler.jobseekerService.UpdateCV(newCVCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update cv", nil))
}

func (handler *JobseekerHandler) DeleteCV(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	errUpdate := handler.jobseekerService.RemoveCV(seekerID)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully delete cv", nil))
}

func (handler *JobseekerHandler) CreateCareer(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newCareer := CareerRequest{}
	newCareer.JobseekerID = uint(seekerID)

	errBind := c.Bind(&newCareer)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	careerCore := RequestCareerToCore(newCareer)

	errCreate := handler.jobseekerService.AddCareer(careerCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	careerResponse := CoreCareerToResponse(careerCore)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create career", careerResponse))
}

func (handler *JobseekerHandler) GetSingleCareer(c echo.Context) error {
	careerID := c.Param("career_id")

	careerID_int, errConv := strconv.Atoi(careerID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobseekerService.GetCareerByID(uint(careerID_int))
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	careerResponse := CoreCareerToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get detail career", careerResponse))
}

func (handler *JobseekerHandler) GetAllCareers(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errAll := handler.jobseekerService.GetCareerList(seekerID)
	if errAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errAll.Error(), nil))
	}

	careerResponse := CoreCareersToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get all careers", careerResponse))
}

func (handler *JobseekerHandler) UpdateCareer(c echo.Context) error {
	careerID := c.Param("career_id")

	careerID_int, errConv := strconv.Atoi(careerID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newUpdate := CareerRequest{}

	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("data update: ", newUpdate)

	newUpdateCore := RequestCareerToCore(newUpdate)

	errUpdate := handler.jobseekerService.UpdateCareer(uint(careerID_int), newUpdateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update career", nil))
}

func (handler *JobseekerHandler) DeleteCareer(c echo.Context) error {
	careerID := c.Param("career_id")

	careerID_int, errConv := strconv.Atoi(careerID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	errDel := handler.jobseekerService.RemoveCareer(uint(careerID_int))
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully delete career", nil))
}

func (handler *JobseekerHandler) CreateEducation(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newEdu := EducationRequest{}
	newEdu.JobseekerID = uint(seekerID)

	errBind := c.Bind(&newEdu)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	eduCore := RequestEduToCore(newEdu)

	errCreate := handler.jobseekerService.AddEducation(eduCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	eduResponse := CoreEducationToResponse(eduCore)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create education", eduResponse))
}

func (handler *JobseekerHandler) GetSingleEducation(c echo.Context) error {
	eduID := c.Param("education_id")

	eduID_int, errConv := strconv.Atoi(eduID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobseekerService.GetEduByID(uint(eduID_int))
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	eduResponse := CoreEducationToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get detail education", eduResponse))
}

func (handler *JobseekerHandler) GetAllEducations(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errAll := handler.jobseekerService.GetEduList(seekerID)
	if errAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errAll.Error(), nil))
	}

	edusResponse := CoreEdusToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get all educations", edusResponse))
}

func (handler *JobseekerHandler) UpdateEducation(c echo.Context) error {
	eduID := c.Param("education_id")

	eduID_int, errConv := strconv.Atoi(eduID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newUpdate := EducationRequest{}

	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("data update: ", newUpdate)

	newUpdateCore := RequestEduToCore(newUpdate)

	errUpdate := handler.jobseekerService.UpdateEducation(uint(eduID_int), newUpdateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update education", nil))
}

func (handler *JobseekerHandler) DeleteEducation(c echo.Context) error {
	eduID := c.Param("education_id")

	eduID_int, errConv := strconv.Atoi(eduID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	errDel := handler.jobseekerService.RemoveEducation(uint(eduID_int))
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully delete education", nil))
}

func (handler *JobseekerHandler) CreateLicense(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	inputLicense, errRead := c.FormFile("license")
	if errRead != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot read file", nil))
	}

	responURL, errURL := handler.jobseekerService.PDF(inputLicense)
	if errURL != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "cannot get resp", nil))
	}

	newLicense := LicenseRequest{}
	newLicense.JobseekerID = uint(seekerID)
	newLicense.License_file = responURL.Location

	errBind := c.Bind(&newLicense)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	licenseCore := RequestLicenseToCore(newLicense)

	errCreate := handler.jobseekerService.AddLicense(licenseCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create license", nil))
}

func (handler *JobseekerHandler) GetSingleLicense(c echo.Context) error {
	licenseID := c.Param("license_id")

	licenseID_int, errConv := strconv.Atoi(licenseID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobseekerService.GetLicenseByID(uint(licenseID_int))
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	licenseResponse := CoreLicenseToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get detail license", licenseResponse))
}

func (handler *JobseekerHandler) GetAllLicenses(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errAll := handler.jobseekerService.GetLicenseList(seekerID)
	if errAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errAll.Error(), nil))
	}

	licensesResonse := CoreLicensesToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get all licenses", licensesResonse))
}

func (handler *JobseekerHandler) UpdateLicense(c echo.Context) error {
	licenseID := c.Param("license_id")

	licenseID_int, errConv := strconv.Atoi(licenseID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newUpdate := LicenseRequest{}

	// oldPubDateString := c.FormValue("pub_date")
	// if oldPubDateString != "" {
	// 	oldPubDate, err := time.Parse("2006-01-02T15:04:05Z", oldPubDateString)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
	// 	}
	// 	newUpdate.Published_date = oldPubDate
	// }

	// OldExpDateString := c.FormValue("exp_date")
	// if OldExpDateString != "" {
	// 	oldExpDate, err := time.Parse("2006-01-02T15:04:05Z", OldExpDateString)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
	// 	}
	// 	newUpdate.Expiry_date = oldExpDate
	// }

	oldLicense, _ := c.FormFile("license")
	if oldLicense != nil {
		respURL, err := handler.jobseekerService.PDF(oldLicense)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		newUpdate.License_file = respURL.Location
	}

	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("data update: ", newUpdate)

	newUpdateCore := RequestLicenseToCore(newUpdate)

	errUpdate := handler.jobseekerService.UpdateLicense(uint(licenseID_int), newUpdateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update license", nil))
}

func (handler *JobseekerHandler) DeleteLicense(c echo.Context) error {
	licenseID := c.Param("license_id")

	licenseID_int, errConv := strconv.Atoi(licenseID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	errDel := handler.jobseekerService.RemoveLicense(uint(licenseID_int))
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully delete license", nil))
}

func (handler *JobseekerHandler) CreateSkill(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	newSkill := SkillRequest{}
	newSkill.JobseekerID = seekerID

	errBind := c.Bind(&newSkill)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	skillCore := RequestSkillToCore(newSkill)

	errCreate := handler.jobseekerService.AddSkill(skillCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully create skill", nil))
}

func (handler *JobseekerHandler) GetSingleSkill(c echo.Context) error {
	skillID := c.Param("skill_id")

	skillID_int, errConv := strconv.Atoi(skillID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	result, errFirst := handler.jobseekerService.GetSkillByID(uint(skillID_int))
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errFirst.Error(), nil))
	}

	skillResponse := CoreSkillToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get detail skill", skillResponse))
}

func (handler *JobseekerHandler) GetAllSkills(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	result, errAll := handler.jobseekerService.GetSkillList(seekerID)
	if errAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error read data. "+errAll.Error(), nil))
	}

	skillsResponse := CoreSkillsToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully get all skills", skillsResponse))
}

func (handler *JobseekerHandler) UpdateSkill(c echo.Context) error {
	skillID := c.Param("skill_id")

	skillID_int, errConv := strconv.Atoi(skillID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	newUpdate := SkillRequest{}
	errBind := c.Bind(&newUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	fmt.Println("data update: ", newUpdate)

	newUpdateCore := RequestSkillToCore(newUpdate)

	errUpdate := handler.jobseekerService.UpdateSkill(uint(skillID_int), newUpdateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error update data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully update skill", nil))
}

func (handler *JobseekerHandler) DeleteSkill(c echo.Context) error {
	skillID := c.Param("skill_id")

	skillID_int, errConv := strconv.Atoi(skillID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse(http.StatusBadRequest, "error convert id param", nil))
	}

	errDel := handler.jobseekerService.RemoveSkill(uint(skillID_int))
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(http.StatusInternalServerError, "error delete data "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse(http.StatusOK, "successfully delete skill", nil))
}
