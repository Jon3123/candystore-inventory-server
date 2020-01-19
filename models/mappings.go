package models

//Mapping mapping struct
type Mapping struct {
	Source string
	Target string
	Count  int
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
