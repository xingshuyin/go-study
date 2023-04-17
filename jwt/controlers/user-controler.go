package controlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"project/database"
	helpers "project/helpers"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(b)
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providePassword), []byte(userPassword))
	if err != nil {
		return false, "password is incrrect"
	}
	return true, ""
}

// func Signup() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var ctx_, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		var user models.User
// 		if err := ctx.BindJSON(&user); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}
// 		fmt.Printf("user: %v\n", user)
// 		validateErr := validate.Struct(user)

//			if validateErr != nil {
//				ctx.JSON(http.StatusBadRequest, gin.H{"Error": validateErr.Error()})
//			}
//			password := HashPassword(*user.Password)
//			user.Password = &password
//			count, err := userCollection.CountDocuments(ctx_, bson.M{"email": user.Email})
//			defer cancel()
//			if err != nil {
//				log.Fatal(err)
//				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checked for the"})
//			}
//			if count > 0 {
//				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exprience"})
//			}
//			user.Create_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
//			user.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
//			user.ID = primitive.NewObjectID()
//			user.User_id = user.ID.Hex()
//			token, refresh, _ := helpers.GenerateAllToken(*user.Email, *user.Name)
//			user.Token = &token
//			user.Refresh = &refresh
//			r, err := userCollection.InsertOne(ctx_, user)
//			if err != nil {
//				msg := fmt.Sprintf("User item was not created")
//				ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
//			}
//			defer cancel()
//			ctx.JSON(http.StatusOK, r)
//		}
//	}
func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// fmt.Printf("user: %v\n", user)
		validationErr := validate.Struct(user)
		fmt.Printf("validationErr: %v\n", validationErr)
		// if validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		// 	return
		// }

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		// count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		// defer cancel()
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		// }

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
		}

		user.Create_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllToken(*user.Email, *user.Name)
		user.Token = &token
		user.Refresh = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ctx_, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx_, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email of passerword is incorrect"})
		}
		valid, e := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !valid {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e})
			return
		}
		if foundUser.Email == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not founded"})

		}
		signedToken, signedRefresh, err := helpers.GenerateAllToken(*foundUser.Email, *foundUser.Name)
		helpers.UpdateAllTokens(signedToken, signedRefresh, foundUser.User_id)
		err = userCollection.FindOne(ctx_, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, foundUser)

	}
}
func Get_users() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helpers.CheckUserType(ctx, "ADMIN"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		// var ctx_, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// limit, err := strconv.Atoi(ctx.Query("limit"))
		// if err != nil {
		// 	limit = 10
		// }
		// page, err := strconv.Atoi(ctx.Query("page"))
		// if err != nil {
		// 	page = 1
		// }
		ctx.JSON(http.StatusOK, gin.H{"all user": "aaaaaaaaaaaa"})
	}
}
func Get_user() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		if err := helpers.MatchUserTypeToUid(ctx, user_id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx_, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx_, bson.M{"user_id": user_id}).Decode(&user)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, user)
	}
}
