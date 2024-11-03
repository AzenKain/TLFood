package services

import (
	"app_food/pkg/cache"
	"app_food/pkg/database"
	"app_food/pkg/dtos"
	"app_food/pkg/mail"
	"app_food/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var otpCollection *mongo.Collection = database.GetCollection("otps")

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyJWT(tokenString, typeJwt string) (*models.CustomClaims, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv(typeJwt)

	if jwtSecret == "" {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "missing JWT_SECRET in environment")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method!")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenToken(Id primitive.ObjectID, Email string) (*models.AuthResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	if jwtSecret == "" || jwtRefreshSecret == "" {
		return nil, errors.New("missing JWT secrets in environment")
	}

	claimsAccess := jwt.MapClaims{
		"Id":    Id,
		"Email": Email,
		"exp":   time.Now().Add(6 * time.Hour).Unix(),
	}

	claimsRefresh := jwt.MapClaims{
		"Id":    Id,
		"Email": Email,
		"exp":   time.Now().Add(15 * 24 * time.Hour).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	at, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := refreshToken.SignedString([]byte(jwtRefreshSecret))
	if err != nil {
		return nil, err
	}

	res := models.AuthResponse{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return &res, nil
}

func SaveNewRefreshToken(ctx context.Context, Id primitive.ObjectID, RefreshToken string) error {
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": Id}).Decode(&user)
	if err != nil {
		return err
	}
	user.RefreshToken = RefreshToken

	_, err = userCollection.UpdateOne(
		ctx,
		bson.M{"_id": Id},
		bson.M{"$set": bson.M{"refresh_token": RefreshToken}},
	)
	if err != nil {
		return err
	}

	return nil
}

func LoginService(ctx context.Context, dto *dtos.LoginDto) (*models.AuthResponse, error) {

	user, err := GetUserByEmail(ctx, dto.Email)

	if user == nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "User doesn't exists!"+err.Error())
	}

	if !CheckPasswordHash(dto.Password, user.Hash) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid identity or password!")

	}

	data, err := GenToken(user.Id, user.Email)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "")

	}

	err = SaveNewRefreshToken(ctx, user.Id, data.RefreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save refresh token")
	}

	return data, nil
}


func SignUpService(ctx context.Context, dto *dtos.SignUpDto) (*models.AuthResponse, error) {
	_, err := GetUserByEmail(ctx, dto.Email)

	if err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User already exists!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newUser := models.User{
		Id:           primitive.NewObjectID(),
		Email:        dto.Email,
		RefreshToken: "",
		Detail: &models.UserDetail{
			FirstName:   dto.FirstName,
			LastName:    dto.LastName,
			PhoneNumber: dto.PhoneNumber,
		},
		Role:      []string{"user"},
		Hash:      string(hashedPassword),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	data, err := GenToken(user.Id, user.Email)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error can't gen token")

	}

	err = SaveNewRefreshToken(ctx, user.Id, data.RefreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save refresh token")
	}

	return data, nil

}

func RefreshService(ctx context.Context, c *fiber.Ctx) (*models.AuthResponse, error) {
	userData := c.Locals("user")
	if userData == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	user, ok := userData.(*models.User)
	if !ok {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Invalid user data")
	}

	data, err := GenToken(user.Id, user.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate new token")
	}

	err = SaveNewRefreshToken(ctx, user.Id, data.RefreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save refresh token")
	}

	return data, nil
}

func CreateOtpService(ctx context.Context, dto *dtos.CreateOtpDto) (*models.OtpResponse, error) {
	user, err := GetUserByEmail(ctx, dto.Email)
	if err != nil && dto.OtpType != "SignUp" {
		return nil, err
	}
	newOtp := models.Otp{
		Id:        primitive.NewObjectID(),
		Email:     dto.Email,
		IsDisplay: true,
		Value:     false,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	if dto.OtpType == "SignUp" {
		otp := GenerateOTP()
		err := mail.ML.SendTemplateEmail(
			user.Email, "createAccount",
			"Otp Create Account TLFood",
			map[string]interface{}{
				"username": "",
				"code":     otp,
			},
		)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to send otp")
		}
		newOtp.OtpCode = otp
		newOtp.OtpType = "SignUp"

	} else if dto.OtpType == "ForgotPassword" {
		otp := GenerateOTP()
		err := mail.ML.SendTemplateEmail(
			user.Email, "forgetPassword",
			"Otp Forget Password TLFood",
			map[string]interface{}{
				"username": user.Detail.LastName + " " + user.Detail.FirstName,
				"code":     otp,
			},
		)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to send otp")
		}
		newOtp.OtpCode = otp
		newOtp.OtpType = "ForgotPassword"
	}

	if newOtp.OtpType == "" {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create otp")
	}

	result, err := otpCollection.InsertOne(ctx, newOtp)
	if err != nil {
		return nil, err
	}
	return &models.OtpResponse{
		Id:     result.InsertedID,
		Status: false,
	}, nil

}
func VerifyOtpService(ctx context.Context, dto *dtos.VerifyOtpDto) (*models.OtpResponse, error) {
	otp, err := GetOtpById(ctx, dto.OtpId)
	if err != nil {
		return nil, err
	}
	if otp.Email != dto.Email || otp.OtpType != dto.OtpType {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to verify otp")
	}

	if otp.OtpCode != dto.OtpCode {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Otp code not match")
	}

	updateResult, err := otpCollection.UpdateOne(
		ctx,
		bson.M{"_id": otp.Id},
		bson.M{
			"$set": bson.M{
				"value":      true,
				"is_display": false,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if updateResult.MatchedCount == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "OTP not found")
	}

	return &models.OtpResponse{
		Id:     otp.Id,
		Status: true,
	}, nil
}

func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp)
}

func GetOtpById(ctx context.Context, Id string) (*models.Otp, error) {
	var otp models.Otp
	keys := fmt.Sprintf("otps:id:%s", Id)
    value, err := cache.RI.Get(ctx, keys).Result()
    if err == nil {
        if err := json.Unmarshal([]byte(value), &otp); err == nil {
            return &otp, nil
        }
    }
	otpId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Error invalid ID format")
	}

	err = otpCollection.FindOne(ctx, bson.M{"_id": otpId}).Decode(&otp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fiber.NewError(fiber.StatusNotFound, "Otp not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error otp id not found")
	}
	otpData, _ := json.Marshal(otp)
    cache.RI.Set(ctx, keys, otpData, time.Minute * 3)
	return &otp, nil
}