package main


import (
	"fmt"
	"go/frame/config"

)


func main(){



	dd:=config.GetConfigDir()

	fmt.Println(dd)

}