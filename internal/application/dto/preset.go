package dto

type PresetListDTO struct {
	Exercises []*ExerciseDTO `json:"exercises"`
}

type ExerciseDTO struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Sets []*SetDTO `json:"sets"`
}

type SetDTO struct {
	Reps    int     `json:"reps"`
	Weight  float32 `json:"weight"`
	Minutes int     `json:"minutes"`
	Meters  int     `json:"meters"`
}
