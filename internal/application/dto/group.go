package dto

type Group struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ExerciseGroupTypeList struct {
	Groups []Group `json:"groups"`
}
