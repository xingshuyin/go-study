package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	user_type := c.GetString("user_type")
	err = nil
	if user_type != role {
		err = errors.New("Unathorized to access this resource")
	}
	return err
}
func MatchUserTypeToUid(c *gin.Context, user_id string) (err error) {
	user_type := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if user_type == "USER" && uid != user_id {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, user_type)
	return err
}
