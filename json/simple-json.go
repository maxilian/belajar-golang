package main

import (
	"encoding/json"
	"fmt"
)

type Mobil struct {
  Nama string
  Tahun uint
  Manufaktur string
}

func main() {
  var ferrari Mobil

  ferrariJson := `{
      "nama": "Ferrari California",
      "tahun": 2008,
      "manufaktur": "Ferrari"
    }`

  json.Unmarshal([]byte(ferrariJson), &ferrari)
	fmt.Println("nama :", ferrari.Nama)
	fmt.Println("tahun :", ferrari.Tahun)
	fmt.Println("manufaktur :", ferrari.Manufaktur)
  
}
