package update

type Update struct {
	Phone   int    `json:"phone_number"`
	DOB     string   `json:"date_of_birth"`
	Address string `json:"address"`
	Bio     string `json:"bio"`
}