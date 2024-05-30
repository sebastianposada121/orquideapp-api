package auth_usecase

import (
	"errors"
	"orquideapp/src/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	UserRepository        domain.UserRepository
	EmployeeRepository    domain.EmployeesRepository
	BeneficiaryRepository domain.BeneficiaryRepository
}

func (uc *UseCase) authenticate(credentials domain.Login, storedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password))
}

func (uc *UseCase) generateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func (uc *UseCase) login(credentials domain.Login, getData func(email string) (storePassword string, claims jwt.Claims, err error)) (string, error) {
	storePassword, claims, err := getData(credentials.Email)
	if err != nil {
		return "", err
	}

	if err := uc.authenticate(credentials, storePassword); err != nil {
		return "", err
	}

	return uc.generateToken(claims)
}

func getClaims(id int, name string, email string, extraClaims map[string]interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{
		"ID":    id,
		"Name":  name,
		"Email": email,
		"RegisteredClaims": jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	for k, v := range extraClaims {
		claims[k] = v
	}

	return claims
}

func (uc *UseCase) UserLogin(credentials domain.Login) (string, error) {
	getUserPasswordAndId := func(email string) (string, jwt.Claims, error) {
		user, err := uc.UserRepository.GetByEmail(email)
		if err != nil {
			return "", nil, err
		}
		claims := getClaims(user.ID, user.Name, user.Email, nil)
		return user.Password, claims, nil
	}

	return uc.login(credentials, getUserPasswordAndId)
}

func (uc *UseCase) EmployeeLogin(credentials domain.Login) (string, error) {
	getEmployeePasswordAndId := func(email string) (string, jwt.Claims, error) {
		employee, err := uc.EmployeeRepository.GetByEmail(email)
		if err != nil {
			return "", nil, err
		}
		extraClaims := map[string]interface{}{"IpsID": employee.IpsId}
		claims := getClaims(employee.ID, employee.Name, employee.Email, extraClaims)
		return employee.Password, claims, nil
	}

	return uc.login(credentials, getEmployeePasswordAndId)
}

func (uc *UseCase) BeneficiaryLogin(credentials domain.Login) (string, error) {
	getClientPasswordAndId := func(email string) (string, jwt.Claims, error) {
		beneficiary, err := uc.BeneficiaryRepository.GetByEmail(email)
		if err != nil {
			return "", nil, err
		}
		claims := getClaims(beneficiary.ID, beneficiary.Name, beneficiary.Email, nil)
		return beneficiary.Password, claims, nil
	}

	return uc.login(credentials, getClientPasswordAndId)
}

func (uc *UseCase) UserSignup(user domain.User) error {

	if _, err := uc.UserRepository.GetByEmail(user.Email); err == nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.UserRepository.Create(user)
}

func (uc *UseCase) UpdateEmployeePassword(credentials domain.Login) error {
	return uc.EmployeeRepository.UpdatePassword(credentials)
}

func (uc *UseCase) UpdateBeneficiaryPassword(credentials domain.Login) error {
	return uc.BeneficiaryRepository.UpdatePassword(credentials)
}
