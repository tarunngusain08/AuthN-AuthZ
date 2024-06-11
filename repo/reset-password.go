package repo

import (
	"AuthN-AuthZ/contracts"
	"crypto/rand"
	"github.com/jmoiron/sqlx"
	"github.com/jordan-wright/email"
	"math/big"
	"net/smtp"
)

type ResetPasswordRepo struct {
	db       *sqlx.DB
	Email    string
	Password string
}

const subject = "OTP for Password Reset"

func NewResetPasswordRepo(db *sqlx.DB) *ResetPasswordRepo {
	return &ResetPasswordRepo{db: db}
}

func GenerateOTP() (string, error) {
	const digits = "0123456789"
	otpLength := 6
	otp := make([]byte, otpLength)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}

func (l *ResetPasswordRepo) ResetPassword(userDetails *contracts.ResetPassword) error {

	otp, err := GenerateOTP()
	if err != nil {
		return err
	}
	e := email.NewEmail()
	e.From = l.Email
	e.To = []string{userDetails.Email}
	e.Subject = subject
	e.Text = []byte(otp)

	// SMTP server configuration
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	smtpUsername := l.Email
	smtpPassword := l.Password

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	return e.Send(smtpHost+":"+smtpPort, auth)
}
