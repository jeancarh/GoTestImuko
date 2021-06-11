package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Organization struct {
	Organization string `json:"organization"`
	Users        []User `json:"users"`
}

type User struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func main() {

	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	type usercsv struct {
		Org  string
		name string
		rol  string
	}
	m := make(map[string][]string)
	roles := []string{}
	for _, line := range csvLines {
		roles = append(roles, line[2])
		m[line[0]+"-"+line[1]] = roles
		roles = []string{}
	}
	validateData(m)

	// final, _ := json.Marshal(users)
	// fmt.Println(string(final))
}

func validateData(m map[string][]string) {
	// roles := []string{}
	// // userStack := []User{}
	// // user := User{}
	// // org := Organization{}
	// for _, value := range m {
	// 	res1 := strings.Split(value, "-")
	// 	roles = append(roles, res1[2])
	// 	m[res1[2]] =
	// }
	fmt.Println(m)

}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func stringInSlice(usr string, list []string) bool {
	for _, v := range list {
		if v == usr {
			return true
		}
	}
	return false
}

func stringInSliceRole(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func stringInSliceOrg(org Organization, list []Organization) bool {
	for _, v := range list {
		if v.Organization == org.Organization {
			fmt.Println(v.Organization + org.Organization)
			return true
		}
	}
	return false
}
