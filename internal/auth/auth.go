package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
CLAIMS SEGUNDO CHAT GPT

Issuer (iss):
- Representa a entidade que emitiu o JWT.
- Esse claim serve para verificar a origem do token. Geralmente, é o servidor de autenticação ou um serviço responsável por gerar o token.

Subject (sub):
- Identifica o "assunto" ou "usuário" ao qual o token se refere.
-  Esse claim pode ser utilizado para identificar de maneira única a pessoa ou entidade a quem o token foi emitido.

Audience (aud):
- O que é: Indica a quem o token se destina. Ou seja, qual(s) serviço(s) ou aplicação(ões) deve(m) aceitar o JWT.
- Esse claim é útil para garantir que o token só será aceito por uma aplicação ou conjunto de aplicações específicas.

Expiration Time (exp):
- Indica a data e hora em que o token expira.
- Esse claim serve para definir o período de validade do token. Após esse tempo, o token não será mais válido.

Not Before (nbf):
- Define o horário em que o token pode começar a ser considerado válido.
- Esse claim impede que o token seja aceito antes de uma determinada data e hora.

Issued At (iat):
- Indica o horário em que o token foi emitido.
- Esse claim ajuda a determinar quando o token foi gerado. Ele pode ser usado para verificar a idade do token e compará-lo com a data de expiração ou a data de início de validade.

JWT ID (jti):
- Um identificador único para o token.
- Esse claim pode ser usado para evitar que o mesmo token seja reutilizado (no caso de revogação de tokens, por exemplo). Serve como um identificador único para o JWT.
*/

func NewToken(userID string) (string, error) {
	mySigningKey := os.Getenv("AUTH_SIGNING_KEY")
	if mySigningKey == "" {
		return "", errors.New("auth signing key not found")
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		Issuer:    "neofarma",
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(mySigningKey))
}

func ParseToken(token string) (string, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SIGNING_KEY")), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		return "", fmt.Errorf("error to parse token: %w", err)
	}

	if !jwtToken.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	return claims.Subject, nil
}

type cookieName string

const cookieAuth cookieName = "neofarma-auth"

func newCooke(value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     string(cookieAuth),
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &cookie
}

func SetCookie(w http.ResponseWriter, userID string) error {
	token, err := NewToken(userID)
	if err != nil {
		return fmt.Errorf("error to generate token to cookie: %w", err)
	}

	cookie := newCooke(token)
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode

	http.SetCookie(w, cookie)

	return nil
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := newCooke("")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

func ReadCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(string(cookieAuth))
	if err == http.ErrNoCookie {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("error to read cookie: %w", err)
	}
	return cookie.Value, nil
}
