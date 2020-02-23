package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
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

/*Construct Google Form URI*/
func generateUri(item *ParentInfo) string {
	//
	/* URI format: https://docs.google.com/forms/d/e/1FAIpQLSc9RnLOEwfZk4DT6cx_cUSfyx42lIraEciMFaMk0vl7aeudGw/viewform?usp=pp_url&
	entry.471162833=PF+EntryID(999)&
	entry.2005620554=PF+Nama+Anak&
	entry.1045781291=PF+Ibu&
	entry.1224974944=PF+Bapa&
	entry.1065046570=PF+Kelas&
	entry.1453029410=PF+Homeroom&
	entry.1166974658=PF+Tempat+Tinggal&
	entry.839337160=PF+Phone
	*/
	var uri strings.Builder
	uri.WriteString("https://docs.google.com/forms/d/e/1FAIpQLSc9RnLOEwfZk4DT6cx_cUSfyx42lIraEciMFaMk0vl7aeudGw/viewform?usp=pp_url&")
	fmt.Fprintf(&uri, "entry.471162833=%s&", strings.ReplaceAll(item.Id, " ", "+"))
	fmt.Fprintf(&uri, "entry.2005620554=%s&", strings.ReplaceAll(item.NamaAnak, " ", "+"))
	fmt.Fprintf(&uri, "entry.1045781291=%s&", strings.ReplaceAll(item.Ibu, " ", "+"))
	fmt.Fprintf(&uri, "entry.1224974944=%s&", strings.ReplaceAll(item.Bapa, " ", "+"))
	fmt.Fprintf(&uri, "entry.1065046570=%s&", strings.ToUpper(strings.ReplaceAll(item.Kelas, " ", "")))
	fmt.Fprintf(&uri, "entry.1453029410=%s&", strings.ReplaceAll(item.Homeroom, " ", "+"))
	fmt.Fprintf(&uri, "entry.1166974658=%s&", strings.ReplaceAll(item.TempatTinggal, " ", "+"))
	fmt.Fprintf(&uri, "entry.839337160=%s&", strings.ReplaceAll(item.Phone, " ", "+"))

	uri.WriteString("\n")
	return uri.String()
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

	/* Output URI to terminal */
	var wg sync.WaitGroup

	for _, item := range list.Root {

		uri := generateUri(&item)
		//if item.Id == single {
		go func(inner_uri string, child string) {
			defer wg.Done()

			wg.Add(1)
			fmt.Printf("Opening GoogleForm prefilled for %s\n", child)
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", inner_uri).Start()
			check(err)
		}(uri, item.NamaAnak)
		//}
		//fmt.Printf("uri [%s]", uri)
	}
	wg.Wait()
	/*Execute URI with web browser*/
}
