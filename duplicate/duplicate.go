package duplicate

type Info struct {
	Province string `json:"Province"`
	Age      int    `json:"Age"`
}

type Response struct {
	Res []Info `json:"Data"`
}

var In Response

type Age struct {
	Age0to30  int `json:"0-30"`
	Age31to60 int `json:"31-60"`
	Age61up   int `json:"61+"`
	AgeNA     int `json:"N/A"`
}

func (listP Response) Dupcount() map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range listP.Res {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item.Province]

		if exist {
			duplicate_frequency[item.Province] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item.Province] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}
func (listA Response) Agecount() Age {
	AgeGroup := Age{
		Age0to30:  0,
		Age31to60: 0,
		Age61up:   0,
		AgeNA:     0,
	}
	for _, item := range listA.Res {
		if item.Age > 60 {
			AgeGroup.Age61up += 1
		} else if item.Age > 30 {
			AgeGroup.Age31to60 += 1
		} else if item.Age > 0 {
			AgeGroup.Age0to30 += 1
		} else {
			AgeGroup.AgeNA += 1
		}
	}
	return AgeGroup
}
