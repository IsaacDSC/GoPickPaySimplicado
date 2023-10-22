package contracts

type NotificationMailerInterface interface {
	SentMailer(mailer string) (err error)
}
