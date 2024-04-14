package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	type PayloadImage struct {
		Image    string `json:"image"`
		ImageUrl string `json:"image_url"`
	}

	type PayloadVariant struct {
		Description   string         `json:"description"`
		Price         int            `json:"price"`
		CategoryID    int            `json:"category_id"`
		SubCategoryID int            `json:"sub_category_id"`
		ColorName     string         `json:"color_name"`
		ColorCode     string         `json:"color_code"`
		SizeName      string         `json:"size_name"`
		Size          string         `json:"size"`
		Quantity      int            `json:"quantity"`
		Images        []PayloadImage `json:"image"`
	}

	type Payload struct {
		Name          string           `json:"name"`
		Description   string           `json:"description"`
		CategoryID    int              `json:"category_id"`
		SubCategoryID int              `json:"sub_category_id"`
		Image         string           `json:"image"`
		ImageUrl      string           `json:"image_url"`
		Varaints      []PayloadVariant `json:"varaint"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productOrm := new(model.Product)
	productOrm.Name = payload.Name
	productOrm.Description = payload.Description
	productOrm.CategoryID = payload.CategoryID
	productOrm.SubCategoryID = payload.SubCategoryID
	productOrm.Image = payload.Image
	productOrm.ImageUrl = payload.ImageUrl

	if errSaveProduct := common.Database.Save(productOrm).Error; errSaveProduct != nil {
		common.PrintError(" save product error ", errSaveProduct.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save product")
		return errSaveProduct
	}

	for _, variant := range payload.Varaints {
		productVariantOrm := new(model.ProductVariant)
		productVariantOrm.ProductID = int(productOrm.ID)
		productVariantOrm.Description = variant.Description
		productVariantOrm.Price = variant.Price
		productVariantOrm.CategoryID = variant.CategoryID
		productVariantOrm.SubCategoryID = variant.SubCategoryID
		productVariantOrm.SizeName = variant.SizeName
		productVariantOrm.Size = variant.Size
		productVariantOrm.ColorName = variant.ColorName
		productVariantOrm.ColorCode = variant.ColorCode
		productVariantOrm.Quantity = variant.Quantity

		if errSaveProductVariant := common.Database.Save(productVariantOrm).Error; errSaveProductVariant != nil {
			common.PrintError(" save product variant error ", errSaveProductVariant.Error())
			common.FiberError(c, fiber.StatusBadRequest, "can't save product variant")
			return errSaveProductVariant
		}

		for _, image := range variant.Images {
			productImgOrm := new(model.ProductImage)
			productImgOrm.ProductVariantID = int(productVariantOrm.ID)
			productImgOrm.Image = image.Image
			productImgOrm.ImageUrl = image.ImageUrl

			if errSaveProductImg := common.Database.Save(productImgOrm).Error; errSaveProductImg != nil {
				common.PrintError(" save product img error ", errSaveProductImg.Error())
				common.FiberError(c, fiber.StatusBadRequest, "can't save product img")
				return errSaveProductImg
			}
		}
	}

	return common.FiberSuccess(c)
}
