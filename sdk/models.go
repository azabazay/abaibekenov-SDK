package sdk

type Response struct {
	Docs []map[string]interface{} `json:"docs"`
	// Docs   []interface{} `json:"docs"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
}

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Chapter struct {
	ID          string `json:"_id"`
	Book        string `json:"book"`
	ChapterName string `json:"chapterName"`
}

type Movie struct {
	ID                         string `json:"_id"`
	Name                       string `json:"name"`
	AcademyAwardNominations    int    `json:"academyAwardNominations"`
	AcademyAwardWins           int    `json:"academyAwardWins"`
	RuntimeInMinutes           int    `json:"runtimeInMinutes"`
	RottenTomatoesScore        int    `json:"rottenTomatoesScore"`
	BudgetInMillions           int    `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions int    `json:"boxOfficeRevenueInMillions"`
}

type Character struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Death  string `json:"death"`
	Gender string `json:"gender"`
	Hair   string `json:"hair"`
	Height string `json:"height"`
	Race   string `json:"race"`
	Realm  string `json:"realm"`
	Spouse string `json:"spouse"`
}

type Quote struct {
	ID        string `json:"_id"`
	Character string `json:"character"`
	Movie     string `json:"movie"`
	Dialog    string `json:"dialog"`
}
