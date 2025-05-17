package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/RavenStorm-bit/toy-compiler/evaluator"
	"github.com/RavenStorm-bit/toy-compiler/lexer"
	"github.com/RavenStorm-bit/toy-compiler/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" || line == "quit" {
			fmt.Println("Goodbye!")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		expression := p.ParseExpression()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if expression != nil {
			result := evaluator.Eval(expression)
			fmt.Fprintf(out, "%d\n", result)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "Parser errors:\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}