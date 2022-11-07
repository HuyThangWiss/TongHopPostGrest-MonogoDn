package information

type Humans struct{
	Firstname string `json:"Firstname" bson:"Firstname" binding:"required"`
	Lastname  string `json:"Lastname" bson:"Lastname" binding:"required"`
	Gender    string `json:"Gender" bson:"Gender" binding:"required"`
	Age       int	`json:"Age" bson:"age" Binding:"required"`
	Address   string `json:"Address" bson:"Address" binding:"required"`
	Gmail     string `json:"Gmail" bson:"Gmail" binding:"required"`
}

//Year int64 `json:"Year" bson:"Year" binding:"required `
