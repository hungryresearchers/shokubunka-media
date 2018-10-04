package infrastructure

import (
	"api/domain"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
)

func AddResourceValidator(meta *admin.Resource) {
	meta.AddValidator(&resource.Validator{
		Name:    "ProcessStoreData",
		Handler: ValidateData,
	})
}

func ValidateData(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
	switch record.(type) {
	case *domain.User:
		return ValidateUser(record, metaValues)
	case *domain.Article:
		return nil
	case *domain.Tag:
		return nil
	case *domain.Shop:
		return nil
	default:
		return fmt.Errorf("Invalid Value")
	}
}

func ValidateUser(record interface{}, metaValues *resource.MetaValues) error {
	if meta := metaValues.Get("FirstName"); meta != nil {
		if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
			return validations.NewError(record, "FirstName", "FirstName can't be blank")
		}
	}
	if meta := metaValues.Get("LastName"); meta != nil {
		if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
			return validations.NewError(record, "LastName", "LastName can't be blank")
		}
	}
	if meta := metaValues.Get("Email"); meta != nil {
		if email := utils.ToString(meta.Value); strings.TrimSpace(email) == "" || !isValidEmail(email) {
			return validations.NewError(record, "Email", "Invalid email")
		}
	}
	return nil
}

func isValidEmail(email string) bool {
	if err := checkmail.ValidateFormat(email); err != nil {
		return false
	}
	return true
}
