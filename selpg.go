package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
	"os/exec"
)

type selpgArgs struct {
	startPage int
	endPage int
	inputFileName string
	destination string
	number int
	delimited bool
}

var progname string

func initArgs(args *selpgArgs) {
	flag.Usage = Usage
	flag.IntVar(&args.startPage, "s", -1, "sepcify start page")
	flag.IntVar(&args.endPage, "e", -1, "specify end page")
	flag.IntVar(&args.number, "l", -1, "specify number of line per page")
	flag.BoolVar(&args.delimited, "f", false, "specify if using delimiter")
	flag.StringVar(&args.destination, "d", "", "specify the printer")
	flag.Parse()
}

func processArgs(args *selpgArgs) {
	if args.startPage < 0 && args.endPage < 0 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: first arg must be -s start page\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	endIndex := 2
	if len(os.Args[1]) == 2 {
		endIndex = 3
	}

	if os.Args[endIndex][0] != '-' || os.Args[endIndex][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: second arg must be -e end page\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if args.startPage > args.endPage || args.startPage < 0 || args.endPage < 0 {
		fmt.Fprintf(os.Stderr, "%s: Invalid arguments\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if args.delimited == false {
		if args.number == -1 {
			args.number = 72
		}
	}

	if args.delimited == true && args.number != -1 {
		fmt.Fprintf(os.Stderr, "%s: delimited and number can't coexist.\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

}

func processInput(args *selpgArgs) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd

	if args.destination != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	if flag.NArg() > 0 {
		args.inputFileName = flag.Arg(0)
		output, err := os.Open(args.inputFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(output)
		if args.delimited {
			readByF(reader, args, stdin)
		} else {
			readByLine(reader, args, stdin)
		}
	} else {
		readByTerminal(args, stdin)
	}

	if args.destination != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Usage()  {
	fmt.Printf("Usage of %s:\n\n", progname)
	fmt.Printf("%s is a tool to select pages from what you want.\n\n",
		progname)
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tselpg -s=Number -e=Number [options] [filename]\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t-s=Number\tStart from Page <number>.\n")
	fmt.Printf("\t-e=Number\tEnd to Page <number>.\n")
	fmt.Printf("\t-l=Number\t[options]Specify the number of line per page.Default is 72.\n")
	fmt.Printf("\t-d=lp number\t[options]Using cat to test.")
	fmt.Printf("\t-f\t\t[options]Specify that the pages are sperated by \\f.\n")
	fmt.Printf("\t[filename]\t[options]Read input from the file.\n\n")
	fmt.Printf("If no file specified, %s will read input from stdin. Control-D to end.\n\n", progname)
}



func readByF(reader *bufio.Reader, args *selpgArgs, stdin io.WriteCloser)  {
	for pageNum := 0; pageNum <= args.endPage; pageNum++ {
		line, err := reader.ReadString('\f')
		if err != io.EOF && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		printOrWrite(args, string(line), stdin)
	}
}

func readByLine(reader *bufio.Reader, args *selpgArgs, stdin io.WriteCloser)  {
	count := 0
	for {
		line, _, err := reader.ReadLine()
		if err != io.EOF && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		if count / args.number >= args.startPage {
			if count / args.number >= args.endPage {
				break
			} else {
				printOrWrite(args, string(line), stdin)
			}
		}
		count++
	}
}

func readByTerminal(args *selpgArgs, stdin io.WriteCloser) {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	target := ""
	for scanner.Scan() {
		line := scanner.Text()
		line += "\n"
		if count / args.number >= args.startPage {
			if count / args.number < args.endPage {
				target += line
			}
		}
		count++
	}
	printOrWrite(args, string(target), stdin)
}

func printOrWrite(args *selpgArgs, line string, stdin io.WriteCloser) {
	if args.destination != "" {
		stdin.Write([]byte(line + "\n"))
	} else {
		fmt.Println(line)
	}
}

func main() {
	progname = os.Args[0]
	var args selpgArgs
	initArgs(&args)
	processArgs(&args)
	processInput(&args)
}
