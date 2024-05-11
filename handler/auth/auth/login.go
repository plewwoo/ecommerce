package handler

import (
	"encoding/json"
	"log"
	"time"

	"ecommerce/common"
	model "ecommerce/model/user"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginReg struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	payload := new(LoginReg)
	if err := c.BodyParser(payload); err != nil {
		log.Println("error = ", err)
		return common.FiberReviewPayload(c)
	}

	resultUser := new(model.User)
	sql := `SELECT * FROM users e WHERE e.username = ?`
	common.Database.Raw(sql, payload.Username).Scan(&resultUser)
	if len(resultUser.Username) == 0 {
		return common.FiberError(c, "500", "ไม่พบผู้ใช้งาน")
	}

	hashedPassword := common.CheckPasswordHash(payload.Password, resultUser.Password)

	if hashedPassword {
		jsonUser := common.StructToString(resultUser)
		var data common.Remark
		if err := json.Unmarshal([]byte(jsonUser), &data); err != nil {
			panic(err)
		}

		tokenString, err := common.GenerateJWTToken(resultUser.ID, data, time.Hour*24)
		if err != nil {
			return common.FiberError(c, "5000", "Failed to generate JWT token", err)
		}

		return c.JSON(fiber.Map{"status": 1, "token": tokenString})
	}

	return common.FiberError(c, "500", "พาสเวิร์ดไม่ถูกต้อง")
}
