package routes

import (
	"uber/src/util"
)

func Authorizer(username, password string) bool {
	user, err := userRepository.Getuser(username)
	if err != nil {
		return false
	}

	return util.CheckPasswordHash(password, user.Password)
}

// func Authorizer(c *fiber.Ctx) error {
// 	auth := c.Get(fiber.HeaderAuthorization)

// 	// Check if the header contains content besides "basic".
// 	if len(auth) <= 6 || strings.ToLower(auth[:5]) != "basic" {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// Decode the header contents
// 	raw, err := base64.StdEncoding.DecodeString(auth[6:])
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// Get the credentials
// 	creds := utils.UnsafeString(raw)

// 	// Check if the credentials are in the correct form
// 	// which is "username:password".
// 	index := strings.Index(creds, ":")
// 	if index == -1 {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// Get the username and password
// 	username := creds[:index]
// 	password := creds[index+1:]

// 	user, err := userRepository.Getuser(username)
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusUnauthorized)

// 	}

// 	if util.CheckPasswordHash(password, user.Password) {
// 		c.Context().SetUserValue("UserID", user.ID)
// 		return c.Next()
// 	} else {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}
// }
