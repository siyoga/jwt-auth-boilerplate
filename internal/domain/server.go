package domain

type Mode string

const (
	Local Mode = "local"
	Dev   Mode = "dev"
	Prod  Mode = "prod"
)
