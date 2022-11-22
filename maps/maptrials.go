package main

import "fmt"

func main() {
	kv := map[string]string{
		"Wanda":         "WandaVision",
		"Will Robinson": "Lost In Space",
		"Doron":         "Fauda",
	}
	fmt.Println(kv)
	// Random order
	kv["Grace"] = "Grace and Frankie"
	for key, value := range kv {
		fmt.Println(key, value)
	}

	delete(kv, "Grace")
	fmt.Println(kv)
	fmt.Println(kv["Wandarw"])

}
