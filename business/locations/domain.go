package locations

type Location struct {
	Id            int
	Address1      string
	Address2      string
	Address3      string
	City          string
	ZipCode       string
	Country       string
	State         string
	DisplayAdress string
}

type LocationUseCaseInterface interface {
	Add(location Location) (Location, error)
}

type LocationRepoInterface interface {
	Add(location Location) (Location, error)
}
