# Build ./genfile

go build .

# Usage of ./genfile:

Generate n random integer numbers greater than zero to a file, separated by "\t"

    -file string

    	where you want to save your file (default "../../files/input.txt")

    -n int
    
    	amount of numbers you want to generate to the file (default 10000)