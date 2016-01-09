package main

import (
	"./swarm"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	var d swarm.Resjson
	dec.Decode(&d)
	fmt.Println(d)
	fmt.Println(d.Res.Checkins.Count)
	fmt.Println(d.Res.Checkins.Items[0].ID)
}
