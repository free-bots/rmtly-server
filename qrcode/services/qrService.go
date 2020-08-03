package services

import (
	"github.com/mdp/qrterminal/v3"
	"os"
)

func ShowQr(text string) {
	qrterminal.Generate(text, qrterminal.L, os.Stdout)
}
