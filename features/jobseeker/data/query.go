package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/jobseeker"
	"JobHuntz/utils/responses"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type JobseekerQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) jobseeker.JobseekerDataInterface {
	return &JobseekerQuery{
		db: db,
	}
}

// GetjobseekerByCompany implements jobseeker.JobseekerDataInterface.
func (repo *JobseekerQuery) GetjobseekerByCompany(input uint) (*jobseeker.JobseekerCore, error) {
	var jobData database.Jobseeker
	tx := repo.db.Preload("Cvs").Preload("Careers").Preload("Educations").Preload("Licenses").Preload("Skills").First(&jobData, input)
	if tx.Error != nil {
		return nil, tx.Error
	}

	data := &jobseeker.JobseekerCore{
		ID:                  jobData.ID,
		Full_name:           jobData.Full_name,
		Username:            jobData.Username,
		Email:               jobData.Email,
		Address:             jobData.Address,
		Phone:               jobData.Phone,
		Birth_date:          jobData.Birth_date,
		Gender:              jobData.Gender,
		Resume:              jobData.Resume,
		Status_Verification: jobData.Status_Verification,
		Banners:             jobData.Banners,
		Careers:             make([]jobseeker.CareerCore, len(jobData.Careers)),
		Educations:          make([]jobseeker.EducationCore, len(jobData.Educations)),
		Cvs:                 make([]jobseeker.CVCore, len(jobData.Cvs)),
		Licenses:            make([]jobseeker.LicenseCore, len(jobData.Licenses)),
		Skills:              make([]jobseeker.SkillCore, len(jobData.Skills)),
	}

	for i, license := range jobData.Licenses {
		data.Licenses[i] = jobseeker.LicenseCore{
			ID:             license.ID,
			JobseekerID:    license.JobseekerID,
			License_name:   license.License_name,
			Published_date: license.Published_date,
			Expiry_date:    license.Expiry_date,
			License_file:   license.License_file,
		}
	}

	for i, skil := range jobData.Skills {
		data.Skills[i] = jobseeker.SkillCore{
			ID:          skil.ID,
			JobseekerID: skil.JobseekerID,
			Skill:       skil.Skill,
			Description: skil.Description,
		}
	}

	for i, cvs := range jobData.Cvs {
		data.Cvs[i] = jobseeker.CVCore{
			JobseekerID: cvs.ID,
			CV_file:     cvs.CV_file,
		}
	}

	for i, education := range jobData.Educations {
		data.Educations[i] = jobseeker.EducationCore{
			ID:              education.ID,
			JobseekerID:     education.JobseekerID,
			Education_level: education.Education_level,
			Major:           education.Major,
			Graduation_date: education.Graduation_date,
			Jobseeker:       data,
		}
	}

	for i, career := range jobData.Careers {
		data.Careers[i] = jobseeker.CareerCore{
			ID:           career.ID,
			JobseekerID:  career.JobseekerID,
			Position:     career.Position,
			Company_name: career.Company_name,
			Date_start:   career.Date_start,
			Date_end:     career.Date_end,
			Jobseeker:    data,
		}
	}
	return data, nil
}
func (repo *JobseekerQuery) GetByIdJobSeeker(id uint) (*jobseeker.JobseekerCore, error) {
	var jobData database.Jobseeker
	tx := repo.db.Preload("Cvs").Preload("Careers").Preload("Educations").Preload("Licenses").Preload("Skills").First(&jobData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	data := &jobseeker.JobseekerCore{
		ID:                  jobData.ID,
		Full_name:           jobData.Full_name,
		Username:            jobData.Username,
		Email:               jobData.Email,
		Address:             jobData.Address,
		Phone:               jobData.Phone,
		Birth_date:          jobData.Birth_date,
		Gender:              jobData.Gender,
		Resume:              jobData.Resume,
		Status_Verification: jobData.Status_Verification,
		Banners:             jobData.Banners,
		Careers:             make([]jobseeker.CareerCore, len(jobData.Careers)),
		Educations:          make([]jobseeker.EducationCore, len(jobData.Educations)),
		Cvs:                 make([]jobseeker.CVCore, len(jobData.Cvs)),
		Licenses:            make([]jobseeker.LicenseCore, len(jobData.Licenses)),
		Skills:              make([]jobseeker.SkillCore, len(jobData.Skills)),
	}

	for i, license := range jobData.Licenses {
		data.Licenses[i] = jobseeker.LicenseCore{
			ID:             license.ID,
			JobseekerID:    license.JobseekerID,
			License_name:   license.License_name,
			Published_date: license.Published_date,
			Expiry_date:    license.Expiry_date,
			License_file:   license.License_file,
		}
	}

	for i, skil := range jobData.Skills {
		data.Skills[i] = jobseeker.SkillCore{
			ID:          skil.ID,
			JobseekerID: skil.JobseekerID,
			Skill:       skil.Skill,
			Description: skil.Description,
		}
	}

	for i, cvs := range jobData.Cvs {
		data.Cvs[i] = jobseeker.CVCore{
			JobseekerID: cvs.ID,
			CV_file:     cvs.CV_file,
		}
	}

	for i, education := range jobData.Educations {
		data.Educations[i] = jobseeker.EducationCore{
			ID:              education.ID,
			JobseekerID:     education.JobseekerID,
			Education_level: education.Education_level,
			Major:           education.Major,
			Graduation_date: education.Graduation_date,
			Jobseeker:       data,
		}
	}

	for i, career := range jobData.Careers {
		data.Careers[i] = jobseeker.CareerCore{
			ID:           career.ID,
			JobseekerID:  career.JobseekerID,
			Position:     career.Position,
			Company_name: career.Company_name,
			Date_start:   career.Date_start,
			Date_end:     career.Date_end,
			Jobseeker:    data,
		}
	}
	return data, nil
}

func (repo *JobseekerQuery) Register(input jobseeker.JobseekerRegistCore) error {
	newSeekerGorm := CoreJobseekerRegistToModel(input)
	newSeekerGorm.Password = responses.HashPassword(input.Password)
	newSeekerGorm.Status_Verification = "Unverified"

	tx := repo.db.Create(&newSeekerGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Login implements user.UserDataInterface.
func (repo *JobseekerQuery) AllEmails(email string) (jobseeker.JobseekerCore, error) {
	var dataSeeker database.Jobseeker
	tx := repo.db.Where("email = ?", email).First(&dataSeeker)
	if tx.Error != nil {
		return jobseeker.JobseekerCore{}, errors.New(tx.Error.Error() + ", wrong email")
	}

	userCore := ModelJobseekerToCore(dataSeeker)
	return userCore, nil
}

func (repo *JobseekerQuery) AllUsernames(username string) (jobseeker.JobseekerCore, error) {
	var dataSeeker database.Jobseeker
	tx := repo.db.Where("username = ?", username).First(&dataSeeker)
	if tx.Error == nil {
		return jobseeker.JobseekerCore{}, errors.New("username is already used")
	}

	userCore := ModelJobseekerToCore(dataSeeker)
	return userCore, nil
}

func (repo *JobseekerQuery) UpdateProfile(seekerID uint, data jobseeker.JobseekerUpdateCore) error {
	newUpdateGorm := CoreJobseekerToModelUpdate(data)
	txUpdates := repo.db.Model(&database.Jobseeker{}).Where("id = ?", seekerID).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) PDF(fileHeader *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	src, errOpen := fileHeader.Open()
	if errOpen != nil {
		return nil, errors.New(errOpen.Error() + "cannot open file")

	}
	defer src.Close()

	// Baca isi file ke dalam byte slice
	fileContent, errRead := io.ReadAll(src)
	if errRead != nil {
		return nil, errors.New(errRead.Error() + "cannot read file")
	}

	s3Config := &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIAZI2LCMCTD7PSZZ4Q", "ZdZdTOFg/bT3RkkPFJDzA83Piqtv9MNxOgxll2Zo", ""),
	}

	s3Session, _ := session.NewSession(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("jobhuntzbucket"),                // bucket's name
		Key:         aws.String("static/" + fileHeader.Filename), // files destination location
		Body:        bytes.NewReader(fileContent),                // content of the file
		ContentType: aws.String("application/pdf"),               // content type
	}

	output, errUpload := uploader.UploadWithContext(context.Background(), input)
	if errUpload != nil {
		return nil, errors.New(errUpload.Error() + "cannot upload file")
	}

	return output, nil
}

func (repo *JobseekerQuery) Photo(fileHeader *multipart.FileHeader) (*uploader.UploadResult, error) {
	urlCloudinary := "cloudinary://377166738273893:ga3Zq7Ts84gJ-Ltn-gyMkTgHd40@dltcy9ghn"

	file, errHeader := fileHeader.Open()
	if errHeader != nil {
		return nil, errors.New(errHeader.Error() + "cannot open fileHeader")
	}

	ctx := context.Background()
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	resp, errUpload := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if errUpload != nil {
		return nil, errors.New(errUpload.Error() + "cannot upload file")
	}

	return resp, nil
}

func (repo *JobseekerQuery) CountCV(dbRaw *sql.DB, seekerID uint) (uint, error) {
	var count uint

	query := `SELECT COUNT(cv_file) FROM cvs WHERE jobseeker_id = ? AND deleted_at IS NULL;`

	rowID := dbRaw.QueryRow(query, seekerID)

	rowID.Scan(&count)

	return count, nil
}

func (repo *JobseekerQuery) AddCV(input jobseeker.CVCore) error {
	newCV := CoreCVToModel(input)

	tx := repo.db.Create(&newCV) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) ReadCV(seekerID uint) (jobseeker.CVCore, error) {
	var singleCVGorm database.CV
	tx := repo.db.Where("jobseeker_id = ?", seekerID).First(&singleCVGorm)
	if tx.Error != nil {
		return jobseeker.CVCore{}, errors.New(tx.Error.Error() + "cannot get data of cv")
	}

	singleCVCore := ModelCVToCore(singleCVGorm)

	return singleCVCore, nil
}

func (repo *JobseekerQuery) UpdateCV(input jobseeker.CVCore) error {
	newCVGorm := CoreCVToModel(input)

	txUpdates := repo.db.Model(&database.CV{}).Where("jobseeker_id = ?", newCVGorm.JobseekerID).Updates(newCVGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveCV(input uint) error {
	result := repo.db.Where("jobseeker_id = ?", input).Delete(&database.CV{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete cv")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}

func (repo *JobseekerQuery) AddCareer(input jobseeker.CareerCore) error {
	// simpan ke DB
	newCareerGorm := CoreCareerToModel(input)

	tx := repo.db.Create(&newCareerGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) GetCareerByID(input uint) (jobseeker.CareerCore, error) {
	var singleCareerGorm database.Career
	tx := repo.db.First(&singleCareerGorm, input)
	if tx.Error != nil {
		return jobseeker.CareerCore{}, tx.Error
	}

	singleCareerCore := ModelCareerToCore(singleCareerGorm)

	return singleCareerCore, nil
}

func (repo *JobseekerQuery) GetCareerList(input uint) ([]jobseeker.CareerCore, error) {
	var careersDataGorm []database.Career
	tx := repo.db.Where("jobseeker_id = ?", input).Find(&careersDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allCareersCore := ModelCareersToCore(careersDataGorm)

	return allCareersCore, nil
}

func (repo *JobseekerQuery) UpdateCareer(career_id uint, input jobseeker.CareerCore) error {
	newCareerGorm := CoreCareerToModel(input)

	txUpdates := repo.db.Model(&database.Career{}).Where("id = ?", career_id).Updates(newCareerGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveCareer(input uint) error {
	result := repo.db.Where("id = ?", input).Delete(&database.Career{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete career")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}

func (repo *JobseekerQuery) AddEducation(input jobseeker.EducationCore) error {
	// simpan ke DB
	newEduGorm := CoreEducationToModel(input)

	tx := repo.db.Create(&newEduGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) GetEduByID(input uint) (jobseeker.EducationCore, error) {
	var singleEduGorm database.Education
	tx := repo.db.First(&singleEduGorm, input)
	if tx.Error != nil {
		return jobseeker.EducationCore{}, tx.Error
	}

	singleEduCore := ModelEduToCore(singleEduGorm)

	return singleEduCore, nil
}

func (repo *JobseekerQuery) GetEduList(input uint) ([]jobseeker.EducationCore, error) {
	var edusDataGorm []database.Education
	tx := repo.db.Where("jobseeker_id = ?", input).Find(&edusDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allEdusCore := ModelEdusToCore(edusDataGorm)

	return allEdusCore, nil
}

func (repo *JobseekerQuery) UpdateEducation(eduID uint, input jobseeker.EducationCore) error {
	newEduGorm := CoreEducationToModel(input)

	txUpdates := repo.db.Model(&database.Education{}).Where("id = ?", eduID).Updates(newEduGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveEducation(input uint) error {
	result := repo.db.Where("id = ?", input).Delete(&database.Education{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete education")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}

func (repo *JobseekerQuery) AddLicense(input jobseeker.LicenseCore) error {
	// simpan ke DB
	newLicenseGorm := CoreLicenseToModel(input)

	tx := repo.db.Create(&newLicenseGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) GetLicenseByID(input uint) (jobseeker.LicenseCore, error) {
	var singleLicenseGorm database.License
	tx := repo.db.First(&singleLicenseGorm, input)
	if tx.Error != nil {
		return jobseeker.LicenseCore{}, tx.Error
	}

	singleLicenseCore := ModelLicenseToCore(singleLicenseGorm)

	return singleLicenseCore, nil
}

func (repo *JobseekerQuery) GetLicenseList(input uint) ([]jobseeker.LicenseCore, error) {
	var licensesDataGorm []database.License
	tx := repo.db.Where("jobseeker_id = ?", input).Find(&licensesDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allLicensesCore := ModelLicensesToCore(licensesDataGorm)

	return allLicensesCore, nil
}

func (repo *JobseekerQuery) UpdateLicense(licenseID uint, input jobseeker.LicenseCore) error {
	newLicenseGorm := CoreLicenseToModel(input)

	txUpdates := repo.db.Model(&database.License{}).Where("id = ?", licenseID).Updates(newLicenseGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveLicense(input uint) error {
	result := repo.db.Where("id = ?", input).Delete(&database.License{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete license")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}

func (repo *JobseekerQuery) AddSkill(input jobseeker.SkillCore) error {
	// simpan ke DB
	newSkillGorm := CoreSkillToModel(input)

	tx := repo.db.Create(&newSkillGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *JobseekerQuery) GetSkillByID(input uint) (jobseeker.SkillCore, error) {
	var singleSkillGorm database.Skill
	tx := repo.db.First(&singleSkillGorm, input)
	if tx.Error != nil {
		return jobseeker.SkillCore{}, tx.Error
	}

	singleSkillCore := ModelSkillToCore(singleSkillGorm)

	return singleSkillCore, nil
}

func (repo *JobseekerQuery) GetSkillList(input uint) ([]jobseeker.SkillCore, error) {
	var skillsDataGorm []database.Skill
	tx := repo.db.Where("jobseeker_id = ?", input).Find(&skillsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allSkillsCore := ModelSkillsToCore(skillsDataGorm)

	return allSkillsCore, nil
}

func (repo *JobseekerQuery) UpdateSkill(skillID uint, input jobseeker.SkillCore) error {
	newSkillGorm := CoreSkillToModel(input)

	txUpdates := repo.db.Model(&database.Skill{}).Where("id = ?", skillID).Updates(newSkillGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *JobseekerQuery) RemoveSkill(input uint) error {
	result := repo.db.Where("id = ?", input).Delete(&database.Skill{})

	if result.Error != nil {
		return errors.New(result.Error.Error() + "cannot delete skill")
	}

	fmt.Println("row affected: ", result.RowsAffected)

	return nil
}
