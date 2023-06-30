package mongodb

type Port struct {
	Key         string    `bson:"_id"`
	Name        string    `bson:"name"`
	City        string    `bson:"city"`
	Country     string    `bson:"country"`
	Alias       []string  `bson:"alias"`
	Regions     []string  `bson:"regions"`
	Coordinates []float64 `bson:"coordinates"`
	Province    string    `bson:"province"`
	Timezone    string    `bson:"timezone"`
	Unlocs      []string  `bson:"unlocs"`
	Code        string    `bson:"code"`
}
