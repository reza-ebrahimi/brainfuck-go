package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"github.com/urfave/cli"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	l := len(*s)
	if l > 0 {
		op := (*s)[l-1]
		*s = (*s)[:l-1]
		return op
	}
	return ""
}

func (s *Stack) Top() string {
	n := len(*s) - 1
	return (*s)[n]
}

func (s Stack) Len() int {
	return len(s)
}

func Execute(op string, indexPtr *int, program *[]uint32) {

	switch op {
	// Move the pointer to the right
	case ">":
		if *indexPtr == 32000 {
			break
		}
		(*indexPtr)++
	// Move the pointer to the left
	case "<":
		if *indexPtr == 0 {
			*indexPtr = 0
			break
		}
		(*indexPtr)--
	// Increment the memory cell under the pointer
	case "+":
		(*program)[*indexPtr]++
	// Decrement the memory cell under the pointer
	case "-":
		if (*program)[*indexPtr] == 0 {
			(*program)[*indexPtr] = 255
			break
		}
		(*program)[*indexPtr]--
	// Output the character signified by the cell at the pointer
	case ".":
		character := string((*program)[*indexPtr])
		fmt.Printf("%s", character)
	// Input a character and store it in the cell at the pointer
	case ",":
		scanner := bufio.NewScanner(os.Stdin)
		input, err := strconv.ParseUint(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		(*program)[*indexPtr] = uint32(input)
	default:
	}
}

func Interpret(stream io.Reader) {
	var mainStack Stack
	var loopStack Stack
	var op string
	buf := make([]byte, 1)
	program := make([]uint32, 32000)
	stackPtr := 0

	for {
		if loopStack.Len() > 0 {
			op = loopStack.Pop()
		} else {
			_, err := io.ReadFull(stream, buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
				break
			}
			op = string(buf)
		}
		switch op {
		// Execute on operators
		case ">", "<", "+", "-", ".", ",":
			Execute(op, &stackPtr, &program)
			mainStack.Push(op)
			break
		case "[":
			mainStack.Push(op)
			break
		case "]":
			mainStack.Push(op)
			if program[stackPtr] > 0 {
				innerLoop := 0
				firsttimehit := false
				for {
					operation := mainStack.Pop()
					if operation == "" {
						break
					}
					loopStack.Push(operation)
					// nested loops
					if operation == "]" && firsttimehit {
						innerLoop++
					}
					if operation == "[" {
						if innerLoop == 0 {
							break
						} else {
							innerLoop--
						}
					}
					firsttimehit = true
				}
			}
		default:
		}
	}
}

func main() {
	app := &cli.App{
		Name:  "Brainfuck Interpreter",
		Usage: "A Brainfuck cli interpreter",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				file, err := os.Open(c.Args().Get(0))
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				Interpret(file)
			} else {
				log.Fatal("Fatal error: No input file\n")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
