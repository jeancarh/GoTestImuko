package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
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
	for _, line := range csvLines {
		m[line[0]+"-"+line[1]] = append(m[line[0]+"-"+line[1]], line[2])
	}
	validateData(m)
}

func validateData(m map[string][]string) {
	var users []User
	var orgs []Organization
	for k, v := range m {
		res1 := strings.Split(k, "-")
		user := User{
			Username: res1[1],
			Roles:    v,
		}
		users = append(users, user)
		org := Organization{
			Organization: res1[0],
			Users:        users,
		}
		orgs = append(orgs, org)
		users = []User{}
	}
	sort.Slice(orgs, func(i, j int) bool {
		return orgs[i].Organization < orgs[j].Organization
	})
	final, _ := json.Marshal(orgs)
	fmt.Println(string(final))
}
