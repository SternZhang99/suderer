package service

import (
	"github.com/gin-gonic/gin"
	"strings"
	"sunderer/models"
	"sunderer/utils"
)

// 基本的census Url
const basicUrl = "https://census.daybreakgames.com/get/ps2:v2/character/?name.first_lower="
const itemUrl = "https://census.daybreakgames.com/get/ps2/item?item_id="
const statUrl = "https://census.daybreakgames.com/get/ps2:v2/single_character_by_id?character_id="

// GetUserInfo 获取用户信息handler
func GetUserInfo(c *gin.Context) {
	userName := c.Query("userName") //
	userName = strings.ToLower(userName)
	basicInfoUrl := utils.UrlMaker(basicUrl, userName)
	userBasic := models.FetchBasicInfo(basicInfoUrl)

	character := userBasic.CharacterList[0]
	var resp models.Response
	userId := character.CharacterId
	resp.Name = character.Name.First

}
