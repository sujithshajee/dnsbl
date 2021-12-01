package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	secure "github.com/elithrar/simple-scrypt"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/user"
)

var params secure.Params

func init() {
	params = secure.DefaultParams
}

// Generate user creates a new user in the authentication system
func GenerateUser(ctx context.Context, cl *ent.Client, username, password string) error {
	hash, err := secure.GenerateFromPassword([]byte(password), params)
	if err != nil {
		return fmt.Errorf("user creation: %w", err)
	}

	_, err = cl.User.Create().
		SetUsername(username).
		SetPassword(hash).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("user creation: %w", err)
	}

	return nil
}

//check if incoming password matches the DB value
func CheckPassword(hash []byte, password []byte) error {
	return secure.CompareHashAndPassword(hash, []byte(password))
}

//Basic Auth implementation
func BasicAuth(cl *ent.Client, next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := basicAuth(ctx, cl, r); err != nil {
			http.Error(rw, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, r)
	}
}

func basicAuth(ctx context.Context, cl *ent.Client, r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		return fmt.Errorf("unable to parse basic auth")
	}

	u, err := cl.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return fmt.Errorf("user query: %w", err)
	}

	return CheckPassword(u.Password, []byte(strings.TrimSpace(password)))
}
