/*

	Mark Ira Goldberg 2016

*/


/*
	Tweak bigdigits (optionally starting with preexisting code) to
	1) Output a usage message if param -h or --help is used
	2) Add a row of asterisks before and after the bigdigits if
	param -b or --bar is used
	3) Retain prior functionality if neither new param is used
	
	usage bigdigits [-b|--bar] <whole-number>
	-b --bar  draw an underbar and an overbar
	
	Advice:
	
	a) import strings and use strings.repeat(string, int),
	probably to create the bars (returns string repeated int times)
	b) use either go's basic "flag" package to handle X11 options, 
	or get something better from godashboard.appspot.com/project

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
	//for every row of the '0' bigDigit representation
	//since height is constant for all bigdigits,
	//for the height of the bigDigit representation
	
	for row := range bigDigits[0] {
		line := ""
		//for every digit in the stringOfDigits input
		//and, correspondingly, column of bigDigit output
		for column := range stringOfDigits {
			//convert char input to corresponding integer
			//used as index for bigDigits slice
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		//-b: add the overbar
		if row == 0 && *barFlag {
			//subtract 2 to account for trailing "  " in line
			fmt.Println(strings.Repeat("*", len(line)-2))
		}
		//only print once all digits in the row are handled
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
	
	
	
	
	
	
	
	
	
	
	
	
	

