package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:28139/house-sitting/?page=1&available=&available=1&location=0&city%5Bid%5D=0&city%5Blocality%5D=&city%5Blocality_text%5D=&city%5Badministrative_area_level_2%5D=&city%5Badministrative_area_level_2_text%5D=&city%5Badministrative_area_level_1%5D=&city%5Badministrative_area_level_1_text%5D=&city%5Bcountry%5D=&city%5Bcountry_text%5D=&city%5Blatitude%5D=&city%5Blongitude%5D=&city%5Bzoom%5D=&city%5Bname%5D=&region%5Bid%5D=0&region%5Blocality%5D=&region%5Blocality_text%5D=&region%5Badministrative_area_level_2%5D=&region%5Badministrative_area_level_2_text%5D=&region%5Badministrative_area_level_1%5D=&region%5Badministrative_area_level_1_text%5D=&region%5Bcountry%5D=&region%5Bcountry_text%5D=&region%5Blatitude%5D=&region%5Blongitude%5D=&region%5Bzoom%5D=&region%5Bname%5D=&country=&environment=&population=&period=0&date=2017-03-03&datestart=2017-03-03&dateend=2017-06-24&season=&duration=&isfd=&stopover=", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
