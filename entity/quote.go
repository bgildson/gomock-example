package entity

// Quote represents a finalspace quote
type Quote struct {
	ID        int    `json:"id"`
	Quote     string `json:"quote"`
	By        string `json:"by"`
	Character string `json:"character"`
	Image     string `json:"image"`
}
