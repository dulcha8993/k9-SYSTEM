package services

import (
	"fmt"
	"gorm.io/gorm" 
	"errors"

	// "github.com/gin-gonic/gin"

	dto "k9-system/dto/k9_profile"
)

import (
	config "k9-system/config"
	k9Model "k9-system/models/k9_profile"
	utils "k9-system/utils"
	"k9-system/constants"
)

func CreateK9Profile(
	req dto.CreateK9ProfileRequest,
	) (*k9Model.K9Profile, error) {

	// Start transaction
	tx := config.DB.Begin()

	// Get next sequence value
	var nextVal int64

	if err := tx.Raw(
		"SELECT nextval('k9_code_seq')",
	).Scan(&nextVal).Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToCreate("K9 code"),
		)
	}

	// Generate formatted code
	k9Code := fmt.Sprintf(
		"K9-%06d",
		nextVal,
	)

	//DOB parsing

	dateOfBirth, err := utils.ParseDate(
		req.DateOfBirth,
	)

	if err != nil {
		tx.Rollback()

		return nil, errors.New(
			constants.InvalidDate("DateOfBirth"),
		)
	}

	// ServiceSince parsing

	serviceSince, err := utils.ParseDate(
		req.ServiceSince,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate("ServiceSince"),
		)
	}

	// Optional field

	parsedServiceEnd, err := utils.ParseDate(
		req.ServiceEnd,
	)

	if err != nil {

		tx.Rollback()
		return nil, errors.New(
			constants.InvalidDate("ServiceEnd"),
		)
	}


	k9Profile := k9Model.K9Profile{

		K9Code: k9Code,

		Name: req.Name,

		Breed: req.Breed,

		Gender: req.Gender,

		DateOfBirth: dateOfBirth,

		Microchip: req.Microchip,

		Color: req.Color,

		ServiceSince: serviceSince,

		ServiceEnd: parsedServiceEnd,

		Status: req.Status,
	}

	// Create profile
	if err := tx.Create(&k9Profile).Error; err != nil {

		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return &k9Profile, nil

}


func UpdateK9Profile(
	id string,
	req dto.UpdateK9ProfileRequest,
) (*k9Model.K9Profile, error) {

	var k9Profile k9Model.K9Profile

	if err := config.DB.First(&k9Profile, "id = ?", id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return nil, errors.New(
				constants.NotFound("K9 profile"),
			)

		} else {
			return nil, errors.New(
				constants.FailedToRetrieve("K9 profile"),
			)
		}
	}

	if req.Name != nil {
		k9Profile.Name = *req.Name
	}

	if req.Breed != nil {
		k9Profile.Breed = *req.Breed
	}

	if req.Gender != nil {
		k9Profile.Gender = *req.Gender
	}

	if req.DateOfBirth != nil {
		
		dateOfBirth, err := utils.ParseDate(
			*req.DateOfBirth,
		)

		if err != nil {
			return nil, errors.New(
				constants.InvalidDate("DateOfBirth"),
			)
		}
		k9Profile.DateOfBirth = dateOfBirth
	}

	if req.Microchip != nil {
		k9Profile.Microchip = *req.Microchip
	}

	if req.Color != nil {
		k9Profile.Color = *req.Color
	}

	if req.ServiceSince != nil {
		serviceSince, err := utils.ParseDate(
			*req.ServiceSince,
		)

		if err != nil {
			return nil, errors.New(
				constants.InvalidDate("ServiceSince"),
			)
		}
		k9Profile.ServiceSince = serviceSince
	}

	if req.ServiceEnd != nil {
		serviceEnd, err := utils.ParseDate(
			*req.ServiceEnd,
		)

		if err != nil {
			return nil, errors.New(
				constants.InvalidDate("ServiceEnd"),
			)
		}
		k9Profile.ServiceEnd = serviceEnd
	}

	if req.Status != nil {
		k9Profile.Status = *req.Status
	}

	if err := config.DB.Save(&k9Profile).Error; err != nil {

		return nil, err
	}

	return &k9Profile, nil
}

func GetK9Profiles(
	search string,
	status string,
	microchip string,
	pagination utils.Pagination,
	) ([]k9Model.K9Profile, utils.Pagination, error) {

	var k9Profiles []k9Model.K9Profile

	var total int64

	query := config.DB.Model(&k9Model.K9Profile{})

	if search != "" {
		query = query.Where(
			"name ILIKE ? OR breed ILIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if microchip != "" {
		query = query.Where("microchip = ?", microchip)
	}

	if err := query.Count(&total).Error; err != nil {

		return nil, pagination, errors.New(
			"Failed to count K9 profiles",
		)
	}

	if err := query.Order("created_at DESC").Limit(pagination.Limit).Offset(pagination.Offset).Find(&k9Profiles).Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve("K9 profiles"),
		)
	}

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	return k9Profiles, pagination, nil

}

func GetK9ProfileByID(
	id string,
) (*k9Model.K9Profile, error) {
	var k9Profile k9Model.K9Profile

	if err := config.DB.First(&k9Profile, "id = ?", id).Error; err != nil {
		return nil, err	
	}

	return &k9Profile, nil
}
