package businesses

type Business struct {
	Id            int
	Alias         string
	Name          string
	ImageUrl      string
	IsClosed      bool
	Url           string
	ReviewCount   int
	Rating        float64
	CoordinateId  int // ID of associated coordinate
	TransactionId int // ID of associated transaction
	Price         float64
	LocationId    int // ID of associated location
	Phone         string
	DisplayPhone  string
	Distance      float64
}

type BusinessUseCaseInterface interface {
	Add(business Business) (Business, error)
	Edit(business Business, id int) (Business, error)
}

type BusinessRepoInterface interface {
	Add(business Business) (Business, error)
	Edit(business Business, id int) (Business, error)
}
