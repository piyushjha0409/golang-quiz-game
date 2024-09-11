package main

import (
   "encoding/csv"
   "fmt"
   "flag"
   "os"
)
func main (){
	csvFileName := flag.String("csv", "sample.csv", "here is the csv file for the quiz")
	//add a timeline here 
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to Open a CSV file, %s\n", &csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	// fmt.Printf(lines)
	problems := parseLines(lines);
	// add timer here 
	correct := 0
	for i, p := range problems {
	fmt.Printf("Problem #%d %s = \n", i+1, p.q)
	answerCh := make(chan string)
	var answer string
	fmt.Scanf("%S\n", &answer)

	if answer == p.a {
		correct++
	}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
  ret := make([]problem, len(lines))
  //so make the array and iterator  
  for i, line := range lines {
     ret[i] = problem{
        q: line[0],
		a: line[1],
	 }
  }
  return ret
}

type problem struct {
	q string;
	a string;
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}