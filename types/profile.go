package types

// KohhiProfile is Type of Profile JSON
type KohhiProfile struct {
	Type string `json:"type"`
	Name struct {
		Fullname   string `json:"fullname"`
		Surname    string `json:"surname"`
		Firstname  string `json:"firstname"`
		EFullname  string `json:"e_fullname"`
		ESurname   string `json:"e_surname"`
		EFirstname string `json:"e_firstname"`
	} `json:"name"`
	Age      int `json:"age"`
	Birthday struct {
		Full  string `json:"full"`
		Year  int    `json:"year"`
		Month int    `json:"month"`
		Day   int    `json:"day"`
	} `json:"birthday"`
	Size struct {
		Height int `json:"height"`
		Weight int `json:"weight"`
		BWH    struct {
			B int `json:"B"`
			W int `json:"W"`
			H int `json:"H"`
		} `json:"BWH"`
		BMI float64 `json:"BMI"`
	} `json:"size"`
	Birthplace struct {
		Ja string `json:"ja"`
		En string `json:"en"`
	} `json:"birthplace"`
	DominantHand struct {
		Ja string `json:"ja"`
		En string `json:"en"`
	} `json:"dominant_hand"`
	BloodType  string `json:"blood_type"`
	VoiceActor struct {
		Fullname   string `json:"fullname"`
		Surname    string `json:"surname"`
		Firstname  string `json:"firstname"`
		EFullname  string `json:"e_fullname"`
		ESurname   string `json:"e_surname"`
		EFirstname string `json:"e_firstname"`
	} `json:"voice_actor"`
}
