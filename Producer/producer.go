package main

import (
	"github.com/iwanbk/gobeanstalk"
	"log"
	"time"
	"os"
)

func main() {
	conn, err := gobeanstalk.Dial(os.Getenv("BEANSTALKD_HOST") + ":" + os.Getenv("BEANSTALKD_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		id, err := conn.Put([]byte("hello"), 0, 10*time.Second, 30*time.Second)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Job id %d inserted\n", id)

		time.Sleep(time.Second * 2)
	}
}