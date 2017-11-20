package tool

import (
	"testing"
	"toy/lib/orm"
	"toy/model"
)

func TestFetchHospitals(t *testing.T) {
	hospitals, err := fetchHospitals()
	if err != nil {
		t.Error(err)
	}
	expectedHospitals := []hospital{
		hospital{"珠海市人民医院", 1},
		hospital{"珠海市中西医结合医院", 2},
		hospital{"广东省中医院珠海医院", 3},
		hospital{"珠海市妇幼保健院", 4},
		hospital{"中山大学附属第五医院", 5},
		hospital{"遵义医学院第五附属医院", 6},
		hospital{"珠海市高新区人民医院", 7},
		hospital{"珠海市香洲区人民医院", 9},
		hospital{"珠海市慢性病防治中心", 10},
		hospital{"珠海市斗门区侨立中医院", 11},
		hospital{"珠海市斗门区妇幼保健院", 12},
		hospital{"珠海市香洲区第二人民医院", 13},
		hospital{"珠海市第五人民医院（平沙医院）", 14},
	}
	for _, expectedHospital := range expectedHospitals {
		var match = false
		for _, hospital := range hospitals {
			if hospital.Name == expectedHospital.Name && hospital.Id == expectedHospital.Id {
				match = true
			}
		}
		if !match {
			t.Errorf("%v not in actual %v", expectedHospital, hospitals)
		}
	}
}

func TestUpdateHospitals(t *testing.T) {
	updateHospitals()
	if actual, err := orm.Count(new(model.Hospital)); err != nil || actual != 13 {
		t.Errorf("expected count %d, actual: %d, err: %v", 13, actual, err)
	}
}

func BenchmarkFetchHospitals(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fetchHospitals()
	}
}
