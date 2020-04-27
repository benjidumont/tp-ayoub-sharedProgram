package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Hopital struct {
	LitsPris        int `json:"litsPris"`
	LitsDisponibles int `json:"litsDisponibles"`
}

func main() {

	var hopitaux []Hopital
	file, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(file, &hopitaux)

	showHospitals(hopitaux)
	choice, err := hospitalList()
	if err != nil {
		fmt.Printf("%v", err)
	}
	file, _ = ioutil.ReadFile("config.json")
	json.Unmarshal(file, &hopitaux)
	if choice == "1" {
		hopitaux[0].incrementHospital()
	} else if choice == "2" {
		hopitaux[1].incrementHospital()
	} else if choice == "3" {
		hopitaux[2].incrementHospital()
	} else {
		fmt.Println("The terminal entry is invalid")
	}
	file, _ = json.MarshalIndent(hopitaux, "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
	showHospitals(hopitaux)
}

func showHospitals(hopitaux []Hopital) {
	fmt.Print("\n")
	fmt.Println("Liste d'hôpitaux:")
	fmt.Print("\n")
	for i := 0; i < len(hopitaux); i++ {
		fmt.Printf("Hôpital 1:\nlits dispo: %d  -  lits pris %d\n", hopitaux[i].LitsDisponibles, hopitaux[i].LitsPris)
	}
}

func hospitalList() (string, error) {
	fmt.Print("\n")
	fmt.Println("Choix de l'hôpital:")
	fmt.Print("\n")
	fmt.Println("1 - Hopital un")
	fmt.Println("2 - Hopital deux")
	fmt.Println("3 - Hopital trois")
	var resp string
	_, err := fmt.Scanf("%s", &resp)
	if err != nil {
		return "", errors.New("Error while reading the terminal entry")
	}
	return resp, nil
}

func (h *Hopital) incrementHospital() {
	h.LitsDisponibles--
	h.LitsPris++
}
