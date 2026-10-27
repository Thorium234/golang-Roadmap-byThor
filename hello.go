package main 

import "fmt" 

// Main entry point
func main() {     
    fmt.Println("Hello World") 
}
/*
Explanation of Code Parts
Let us take a look at the various parts of the above program −

1. Package Declaration
The first line of the program package main defines the package name in which this program should lie. It is a mandatory statement, as Go programs run in packages. The main package is the starting point to run the program. Each package has a path and name associated with it.

2. Import Packages
The next line import "fmt" is a preprocessor command which tells the Go compiler to include the files lying in the package fmt.

3. Main Function
The next line func main() is the main function where the program execution begins.

4. Comments
The next line /*...*/ is ignored by the compiler and it is there to add comments in the program. Comments are also represented using // similar to Java or C++ comments.

5. Print Statement
The next line fmt.Println(...) is another function available in Go which causes the message "Hello, World!" to be displayed on the screen. Here fmt package has exported Println method which is used to display the message on the screen.

Notice the capital P of Println method. In Go language, a name is exported if it starts with capital letter. Exported means the function or variable/constant is accessible to the importer of the respective package.*/
