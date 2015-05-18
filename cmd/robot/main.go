package main

import "os"
import "fmt"
import "github.com/tomlanyon/toy-gobot/robot"

func main() {
	r := robot.NewRobot(0,5,0,5)

	err := r.Place(1,1,robot.South)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = r.Move()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Left()

	err = r.Move()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = r.Move()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Report()

	r.Right()

	err = r.Move()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Report()

}
