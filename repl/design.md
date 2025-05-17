# REPL Package Design

The Read-Eval-Print Loop (REPL) provides an interactive way to test the
language. It repeatedly reads a line from the user, parses and evaluates it, and
prints the result.

Features:

- Uses a `bufio.Scanner` to read input line by line.
- Recognizes `exit` or `quit` commands to terminate the session.
- Delegates lexing, parsing, and evaluation to their respective packages.

The REPL is intended for experimentation and will later evolve to support
executing source files directly.

