package handler

import (
	"ecommerce/common"
	model "ecommerce/model/user"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	type Payload struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	hash, _ := common.HashPassword(payload.Password)

	userOrm := new(model.User)

	var checkUsername string
	sql := `SELECT username FROM users WHERE username = ?`
	common.Database.Raw(sql, payload.Username).Scan(&checkUsername)

	if checkUsername == "" {
		userOrm.Username = payload.Username
		userOrm.Password = string(hash)
		userOrm.Firstname = payload.Firstname
		userOrm.Lastname = payload.Lastname
		errSave := common.Database.Create(userOrm)

		if errSave.Error != nil {
			common.PrintError(" save error error ", errSave.Error.Error())
			common.FiberError(c, "400", "can't save")
		}

		return common.FiberSuccess(c)
	} else {
		return common.FiberError(c, "400", "Username already exists")
	}
}
