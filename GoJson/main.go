package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ParentInfo struct {
	Id            string `json:"id"`
	NamaAnak      string `json:"nama anak"`
	Ibu           string `json:"ibu"`
	Bapa          string `json:"bapa"`
	Kelas         string `json:"kelas"`
	Homeroom      string `json:"homeroom"`
	TempatTinggal string `json:"tempat tinggal"`
	Phone         string `json:"phone"`
}

type ParentInfoList struct {
	Root []ParentInfo `json:"root"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/* Read file */
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	//fmt.Print(string(dat))

	/* Unmarshall Json data*/
	var list ParentInfoList
	err = json.Unmarshal(dat, &list)
	check(err)

	/* Output array to terminal */
	for _, item := range list.Root {
		fmt.Printf("Id [%s], Name[%s], Ibu[%s], Bapa[%s]\n", item.Id, item.NamaAnak, item.Ibu, item.Bapa)
	}
}
