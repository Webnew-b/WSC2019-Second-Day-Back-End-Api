package organizerDao

import (
	"errors"
	"gorm.io/gorm"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/organizerError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools"
)

func GetOrganizerInfoById(id int64) (*api.ApiOrganizer, error) {
	var organizer api.ApiOrganizer
	data := database.GetDatabase().Model(&model.Organizers{}).Find(&organizer, id)
	if data.Error != nil {
		return nil, data.Error
	}
	return &organizer, nil
}

func GetOrganizerIdBySlug(slug string) (int64, error) {
	var organizer api.ApiOrganizer
	data := database.GetDatabase().Model(&model.Organizers{}).Where(api.ApiOrganizer{Slug: slug}, "slug").First(&organizer)
	err := checkedError(data.Error, slug)
	if err != nil {
		return 0, err
	}
	return organizer.ID, nil
}

func checkedError(err error, msg string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		tools.Log.Println(msg, "organizer is not found")
		return &organizerError.OrganizerSlugNotFoundError{}
	case err != nil:
		return err
	}
	return nil
}
