package auth

import (
	"context"
	"errors"
	"time"

	"github.com/StevenYAMBOS/wilo-api/config"
	"github.com/StevenYAMBOS/wilo-api/database"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string    `bson:"_id,omitempty"`
	Email          string    `bson:"email"`
	Password       string    `bson:"password"`
	ProfilePicture string    `bson:"profilePicture,omitempty" json:"profilePicture"`
	Role           string    `bson:"role"`
	CreatedAt      time.Time `bson:"createdAt,omitempty" json:"createdAt"`
}

// Inscription
func RegisterUser(email, password string) error {
	collection := database.DB.Collection("user")

	var existingUser User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existingUser)
	if err == nil {
		return errors.New("l'utilisateur existe déjà")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	newUser := User{
		Email:     email,
		Password:  string(hashedPassword),
		ProfilePicture:  "https://s2.qwant.com/thumbr/474x316/7/c/193c63927b6275cb2ddbc94c3b120204f41fc841071a71bee94cf4cd835e5a/th.jpg?u=https%3A%2F%2Ftse.mm.bing.net%2Fth%3Fid%3DOIP.vSAtWooVnzmYOpkpcTEyigHaE8%26pid%3DApi&q=0&b=1&p=0&a=0",
		Role:      "user",
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), newUser)
	return err
}

// Connexion
func LoginUser(email, password string) (string, error) {
	collection := database.DB.Collection("user")

	var user User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("utilisateur non trouvé")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("mot de passe incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	return tokenString, err
}
