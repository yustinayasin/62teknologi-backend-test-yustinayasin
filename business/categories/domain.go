package categories

type Category struct {
	Id    int    
	Alias string
	Title string 
}

type CategoryUseCaseInterface interface {
	Add(category Category) (Category, error)
}

type CategoryRepoInterface interface {
	Add(category Category) (Category, error)
}
