package main

// Account type to manage account information
type Account struct {
	CardIsActive  bool     `json:"cardIsActive"`
	Limit         float64  `json:"limit"`
	Blacklist     []string `json:"blacklist"`
	IsWhitelisted bool     `json:"isWhitelisted"`
}
