package entity

type Payment struct {
	Name string 
	PhoneNumber int 
	Email string
	BalanceUser float64

	Status string
	//selesain cart and review


}

type Cart struct {
	ReviewID int
	Review Review
	Name string
	PhoneNumber int
	Email string
	StudioPrice float64
	Tax float64
	Total float64
}

type Review struct {
	CartTotal float64
}