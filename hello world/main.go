package main

/*
   Package is the project
   It can have serval file
   different files should have same package name if included in the same

   Two types of package:
   1. executable package : generates a file that we can run
   2. Reusable package : generates a file that we can import in other projects

   // how to know which  package is executable package?

  	 name of the package is "main" then it is an executable package
	 executable package must have main function
*/

import "fmt"

/*
  used to format text, including printing to the console
  fmt is a standard library package

*/

/*
   main function is the entry point for the executable package
   when we run the program, execution starts from main function
*/

func main() {
	fmt.Println("Hello, World!")
}
