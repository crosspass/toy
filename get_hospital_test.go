package main

import (
  "net/http"
  "io/ioutil"
)

type hospital struct {
  Name string
  Id int
}

func fetchHospitals() []hospital, error {
   // Get http body
   url := "https://www.zhyygh.cn/web/hospital-list.jsp"
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
   // Update db
}
