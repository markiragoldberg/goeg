/*
	Mark Ira Goldberg 2016
	Apache License 2.0
	
	Added and implemented:
		-h, --help option
		-b, --bar option
	to bigdigits.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	//note to self: I didn't like this one
//	"github.com/jessevdk/go-flags"
	"github.com/pborman/getopt"
)

func main() {

	barFlag := getopt.BoolLong("bar", 'b', "display an additional overbar and underbar")
	getopt.Lookup("bar").SetOptional()
	
	helpFlag := getopt.BoolLong("help", 'h', "display the help information")
	getopt.Lookup("help").SetOptional()
	
	getopt.Parse()
	args := getopt.Args()
	
	if len(args) == 0 || *helpFlag {
		getopt.PrintUsage(os.Stdout)
		os.Exit(1)
	}
	
	stringOfDigits := args[0]
	//Every row in the final output (all digits are same height)
	for row := range bigDigits[0] {
		line := ""
		//Every digit in the row of output
		for column := range stringOfDigits {
			//convert character of a digit to corresponding integer
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
			//don't print until all digits are added to the row
		}
		//-b: add the overbar
		if row == 0 && *barFlag {
			//subtract 2 to account for trailing "  " in line
			fmt.Println(strings.Repeat("*", len(line)-2))
		}
		fmt.Println(line)
		//-b: add the underbar
		if row == len(bigDigits[0])-1 && *barFlag {
			fmt.Println(strings.Repeat("*", len(line)-2))
		}
	}
}

var bigDigits = [][]string{
	{"  000  ",
	 " 0   0 ",
	 "0     0",
	 "0     0",
	 "0     0",
	 " 0   0 ",
	 "  000  "},
	{" 1 ",
	 "11 ",
	 " 1 ",
	 " 1 ",
	 " 1 ",
	 " 1 ",
	 "111"},
	{" 222 ",
	 "2   2",
	 "   2 ",
	 "  2  ",
	 " 2   ",
	 "2    ",
	 "22222"},
    {" 333 ",
     "3   3", 
     "    3", 
     "  33 ", 
     "    3", 
     "3   3", 
     " 333 "},
    {"   4  ",
     "  44  ", 
     " 4 4  ", 
     "4  4  ", 
     "444444", 
     "   4  ",
     "   4  "},
    {"55555",
     "5    ", 
     "5    ", 
     " 555 ", 
     "    5", 
     "5   5", 
     " 555 "},
    {" 666 ",
     "6    ", 
     "6    ", 
     "6666 ", 
     "6   6", 
     "6   6", 
     " 666 "},
    {"77777",
     "    7", 
     "   7 ", 
     "  7  ", 
     " 7   ", 
     "7    ", 
     "7    "},
    {" 888 ",
     "8   8", 
     "8   8", 
     " 888 ", 
     "8   8", 
     "8   8", 
     " 888 "},
	{" 9999",
	 "9   9",
	 "9   9",
	 " 9999",
	 "    9",
	 "    9",
	 "    9"},
}
	
	
	
	
	
	
	
	
	
	
	
	
	

