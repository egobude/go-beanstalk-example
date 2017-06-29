package main

import (
	"github.com/iwanbk/gobeanstalk"
	"log"
	"os"
)

func main() {
	conn, err := gobeanstalk.Dial(os.Getenv("BEANSTALKD_HOST") + ":" + os.Getenv("BEANSTALKD_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		j, err := conn.Reserve()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id:%d, body:%s\n", j.ID, string(j.Body))
		err = conn.Delete(j.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}