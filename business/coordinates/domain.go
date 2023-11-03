package coordinates

type Coordinate struct {
	Id         int
	Latitude   float64
	Longtitude float64
}

type CoordinateUseCaseInterface interface {
	Add(coordinate Coordinate) (Coordinate, error)
}

type CoordinateRepoInterface interface {
	Add(coordinate Coordinate) (Coordinate, error)
}
