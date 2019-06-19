package main

import (
	_ "api/database"
	orm "api/database"
	"api/router"
)
type Config struct {
	Enabled bool
	Path    string
}

/*var conf Config
func init()  {
	file, err := os.Open("conf/conf.json")
	if err !=nil {
		fmt.Println("打开失败",err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf = Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

}*/

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}