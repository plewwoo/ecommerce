package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FiberDisableXFrame(c *fiber.Ctx) error {
	c.Set("X-Frame-Options", "DENY")
	return c.Next()
}

func FiberNoSniff(c *fiber.Ctx) error {
	c.Set("X-Content-Type-Options", "nosniff")
	return c.Next()
}

func FiberReviewPayload(c *fiber.Ctx) error {
	PrintError("FiberReviewPayload", "")
	return FiberError(c, "1002", "review your payload")
}

func FiberSuccess(c *fiber.Ctx) error {
	Print("FiberSuccess", getFiberInfo(c))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": 1, "message": "success"})
}

func FiberCustom(c *fiber.Ctx, status int, errorCode string, errorMessage string) error {
	logDesc := getFiberInfo(c)
	logDesc += fmt.Sprintf("\n Return Code: %s", errorCode)
	logDesc += fmt.Sprintf("\n Return Message: %s", errorMessage)
	PrintError("FiberError", logDesc)
	return c.Status(status).JSON(fiber.Map{"status": 0, "code": errorCode, "message": errorMessage})
}

func FiberError(c *fiber.Ctx, errorCode string, errorMessage string, err ...error) error {
	log.Println("Error", err)
	return FiberCustom(c, fiber.StatusInternalServerError, errorCode, errorMessage)
}

func FiberQueryWithCustomDB(c *fiber.Ctx, db *sql.DB, sql string, values ...interface{}) error {
	jsonBytes, err := queryToJSON(db, sql, values...)
	if err != nil {
		PrintError(`SQL Error`, err.Error())
		return FiberError(c, "1001", "sql error")
	}
	return FiberSendData(c, string(jsonBytes))
}

func FiberQuery(c *fiber.Ctx, sql string, values ...interface{}) error {
	return FiberQueryWithCustomDB(c, DatabaseMysql, sql, values...)
}

func FiberSendData(c *fiber.Ctx, json string) error {
	message := `{"status":1, "message":"success", "data":` + json + `}`
	c.Set("Content-Type", "application/json")
	return c.SendString(string(message))
}

func FiberDeleteByID(c *fiber.Ctx, tableName string) error {
	type Delete struct {
		ID       string `json:"id"`
		DeleteBy string `json:"delete_by"`
	}

	var payload Delete
	err := c.BodyParser(payload)
	if err != nil {
		return FiberReviewPayload(c)
	}

	if tableName == `` || payload.ID == `` || payload.DeleteBy == `` {
		return FiberReviewPayload(c)
	}

	result := Database.Exec(`UPDATE ? SET deleted_at = now(), deleted_by = ? WHERE id = ?`, tableName, payload.DeleteBy, payload.ID)
	if result.Error != nil {
		PrintError(`FiberDelete`, result.Error.Error())
		return FiberError(c, "1001", "sql error")
	}

	return FiberSuccess(c)
}

func FiberDeletePermanentByID(c *fiber.Ctx, tableName string) error {
	type Delete struct {
		ID       string `json:"id"`
		DeleteBy string `json:"delete_by"`
	}

	var payload Delete
	err := c.BodyParser(payload)
	if err != nil {
		return FiberReviewPayload(c)
	}

	if tableName == `` || payload.ID == `` {
		return FiberReviewPayload(c)
	}

	result := Database.Exec(`DELETE FROM ? WHERE id = ?`, tableName, payload.ID)
	if result.Error != nil {
		PrintError(`FiberDeletePermanent`, result.Error.Error())
		return FiberError(c, "1001", "sql error")
	}

	return FiberSuccess(c)
}

func FiberWarmUp(app *fiber.App) {
	app.Get("/_ah/warmup", func(c *fiber.Ctx) error {
		message := "Warm-up request succeeded"
		fmt.Println(message)
		return c.Status(http.StatusOK).SendString(message)
	})
}

func queryToJSON(db *sql.DB, sql string, values ...interface{}) ([]byte, error) {
	list := []string{"INSERT ", "UPDATE ", "DELETE ", "CREATE ", "EMPTY ", "DROP ", "ALTER ", "TRUNCATE "}
	if StringExistsInList(strings.ToUpper(sql), list) {
		return nil, errors.New("NOT ALLOW: INSERT/UPDATE/DELETE/CREATE/EMPTY/DROP/ALTER/TRUNCATE")
	}

	rows, err := db.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		m := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}

			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				if val != nil {
					temp := fmt.Sprintf("%v", val)
					temp = strings.Replace(temp, " +0700 +07", "", -1)
					v = temp
					if len(temp) >= 10 {
						if temp[0:10] == "1900-01-01" {
							v = nil
						}
					}
				} else {
					v = val
				}
			}
			m[col] = v
		}
		result = append(result, m)
	}

	return json.Marshal(result)
}

func getFiberInfo(c *fiber.Ctx) string {
	logDesc := fmt.Sprintf("API Path: %s", c.Path())
	logDesc += fmt.Sprintf("\n Method: %s", c.Method())
	logDesc += fmt.Sprintf("\n Authorization: %s", c.Get("Authorization"))
	body, _ := getBodyJson(c)
	logDesc += fmt.Sprintf("\n Body: %s", body)
	return logDesc
}

func getBodyJson(c *fiber.Ctx) (string, error) {
	rawBody := c.Request().Body()

	var params map[string]interface{}
	if err := json.Unmarshal(rawBody, &params); err != nil {
		return "", errors.New("invalid json")
	}

	jsonString, err := json.Marshal(params)
	if err != nil {
		return "", errors.New("error marshaling json")
	}

	return string(jsonString), nil
}
