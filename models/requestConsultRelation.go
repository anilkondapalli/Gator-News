package models

/*RequestConsultRelation has true or false that is obtained from querying the relationship between 2 users*/
type RequestConsultRelation struct {
	Status bool `json:"status"`
}
