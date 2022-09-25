package token

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/subosito/gotenv"
	"octaviusfarrel.dev/latihan_web/models"
	"octaviusfarrel.dev/latihan_web/repositories/pgsql"
)

var secretKey = func() (result string) {
	gotenv.Load()
	result = os.Getenv("TOKEN_SECRET")
	return
}()

type TokenClaim struct {
	CreatedAt int64  `json:"created_at"`
	Salt      string `json:"salt"`
	HashToken string `json:"hash_token"`
	UserID    uint   `json:"user_id"`
	jwt.RegisteredClaims
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type TokenUtil struct {
	tokenRepo pgsql.ITokenRepo
}

func NewTokenUtil() *TokenUtil {
	return &TokenUtil{
		tokenRepo: pgsql.NewTokenRepo(),
	}
}

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func (tokenUtil *TokenUtil) CreateToken(user models.UserModel) (token string, err error) {
	hashToken := randomString(32)
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaim{
		CreatedAt: time.Now().Unix(),
		Salt:      randomString(24),
		HashToken: hashToken,
		UserID:    user.ID,
	})

	tokenModel := models.TokenModel{
		User:          user,
		HashToken:     hashToken,
		IsTokenActive: true,
	}

	err = tokenUtil.tokenRepo.InsertTokenByUser(tokenModel)

	if err != nil {
		return
	}

	token, err = tokenJwt.SignedString([]byte(secretKey))
	return
}

func (tokenUtil *TokenUtil) ValidateToken(inputToken string, permission string) (err error) {
	token, err := jwt.ParseWithClaims(inputToken, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return
	}

	claims, okClaim := token.Claims.(*TokenClaim)

	if !okClaim || claims.UserID == 0 {
		err = errors.New("invalid token type")
		return
	}

	err = tokenUtil.tokenRepo.ValidateToken(claims.HashToken, permission, context.Background())
	return
}
