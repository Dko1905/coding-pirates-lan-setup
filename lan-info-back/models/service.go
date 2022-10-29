package models

import "github.com/google/uuid"

type Service struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Icon          string    `json:"icon"`
	DarkSupported bool      `json:"darkSupported" db:"dark_supported"`
	Online        bool      `json:"online"`
	Hostname      string    `json:"hostname"`
	Port          string    `json:"port"`
}
