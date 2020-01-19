package models

//Mapping mapping struct
type Mapping struct {
	Source string `json:"Source"`
	Target string `json:"Target"`
	Count  int    `json:"Count"`
}

//NewMapping Create a new mapping
func NewMapping(source string, target string, count int) *Mapping {
	m := &Mapping{
		Source: source,
		Target: target,
		Count:  count,
	}

	return m
}
