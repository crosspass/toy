package tool

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"toy/lib/orm"
	"toy/model"
)

type hospital struct {
	Name string
	Id   int
}

func fetchHospitals() ([]hospital, error) {
	// Get http body
	url := "https://www.zhyygh.cn/web/hospital_list.jsp"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Parse http body
	// <h6 style="font-weight:bold;"><a href="hospital.jsp?hospital_id=1" title="">珠海市人民医院</a></h6>
	hospitalIdReg := regexp.MustCompile(`hospital_id=(?P<id>\d+)`)
	hospitalNameReg := regexp.MustCompile(`<a.+>(?P<name>[^<]+)`)

	reg := regexp.MustCompile(`\<h6.+?\<\/h6\>`)
	var hospitals []hospital
	for _, hospital_tag := range reg.FindAllString(string(body), -1) {
		id, err := strconv.Atoi(hospitalIdReg.FindStringSubmatch(hospital_tag)[1])
		if err != nil {
			return hospitals, err
		}
		hospitals = append(hospitals, hospital{hospitalNameReg.FindStringSubmatch(hospital_tag)[1], id})
	}
	// "hospital_id=1"
	return hospitals, nil
}

func updateHospitals() {
	hospitals, err := fetchHospitals()
	if err != nil {
		fmt.Println("%s", err)
		return
	}
	for _, hospital = range hospitals {
		h := Hospita{}
		model.FindOrCreateBy(hospital)
	}
}
