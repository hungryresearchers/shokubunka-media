package infrastructure

import (
	"api/domain"
	"api/service"
	"log"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
)

func defineUserMetaInfo(user *admin.Resource) {
	user.NewAttrs("FirstName", "LastName", "NickName", "Email", "Role", "Password")
	user.EditAttrs("FirstName", "LastName", "NickName", "Email", "InvitationToken", "Role")
	user.IndexAttrs("FirstName", "LastName", "NickName", "Email", "Role")
	user.Meta(&admin.Meta{
		Name:   "Password",
		Type:   "password",
		Valuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: encryptPassword,
	})
	user.Meta(&admin.Meta{
		Name:   "Role",
		Type:   "role",
		Config: &admin.SelectOneConfig{Collection: []string{"user", "writer", "admin"}},
		Valuer: changeRoleStr,
		Setter: ConvertRole,
	})
	user.Action(&admin.Action{
		Name:    "ResetPassword",
		Handler: generateResetPasswordToken,
		Visible: tokenExist,
		Modes:   []string{"edit", "show", "collection", "menu_item"},
	})
}

func changeRoleStr(record interface{}, context *qor.Context) interface{} {
	user := record.(*domain.User)
	switch user.Role {
	case 1:
		return "writer"
	case 2:
		return "admin"
	default:
		return "user"
	}
}

func ConvertRole(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
	var roleNumber int
	if role := utils.ToString(metaValue.Value); role != "" {
		switch role {
		case "writer":
			roleNumber = 1
		case "admin":
			roleNumber = 2
		default:
			roleNumber = 0
		}
		record.(*domain.User).Role = roleNumber
	} else {
		record.(*domain.User).Role = 0
	}
}

func tokenExist(record interface{}, context *admin.Context) bool {
	user := record.(*domain.User)
	if user.ResetPasswordToken == "" {
		return true
	}
	return false
}

func generateResetPasswordToken(argument *admin.ActionArgument) error {
	for _, record := range argument.FindSelectedRecords() {
		token, err := service.GenerateToken()
		if err != nil {
			log.Fatal(err)
			return err
		}
		argument.Context.DB.Model(record.(*domain.User)).Update("ResetPasswordToken", token)
	}
	return nil
}

func encryptPassword(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
	if password := utils.ToString(metaValue.Value); password != "" {
		passwordHash := service.ToHash(password)
		record.(*domain.User).EncryptedPassword = passwordHash
	}
}
