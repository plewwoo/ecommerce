package handler

import (
	"ecommerce/common"
	model "ecommerce/model/user"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return err
	}

	userOrm := new(model.User)

	var checkUsername string
	sql := `SELECT username FROM users WHERE username = ?`
	common.Database.Raw(sql, payload.Username).Scan(&checkUsername)

	if checkUsername == "" {
		userOrm.Username = payload.Username
		userOrm.Password = string(bytes)
		userOrm.Firstname = payload.Firstname
		userOrm.Lastname = payload.Lastname
		errSave := common.Database.Create(userOrm)

		if errSave.Error != nil {
			common.PrintError(" save error error ", errSave.Error.Error())
			common.FiberError(c, fiber.StatusBadRequest, "can't save")
		}

		return common.FiberSuccess(c)
	} else {
		return common.FiberError(c, fiber.StatusBadRequest, "Username already exists")
	}
}
