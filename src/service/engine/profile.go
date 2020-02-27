package engine

import "encoding/json"

type Profile struct {
	Name string
}

type Profiles struct {
}

func LoadProfiler() Profiles {
	profileObject := Profiles{}
	json.Unmarshal([]byte("file"), &profileObject)
	return profileObject
}
