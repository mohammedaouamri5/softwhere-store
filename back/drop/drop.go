package main

import (
	"back-end/initializers"
 
	"fmt"
	"log"
 
 
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}


func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
    var tables []string
	initializers.DB.Raw("SELECT tablename FROM pg_tables WHERE schemaname='public'").Scan(&tables)
	var err  = 0 	
	for _ , v := range tables {

		info :=  initializers.DB.Exec("DROP TABLE public."+v ) 
		if info.Error != nil { err++ } 
		fmt.Print(info)

	} 

	if err != 0 { main() }

}
