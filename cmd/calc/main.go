package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	op := os.Args[1]

	reader := bufio.NewReader(os.Stdin)

	switch op {
	case "avg":
		if err := avg(reader); err != nil {
			log.Fatalln(err)
		}
	}
}

func avg(r *bufio.Reader) error {
	var count, total float64
	defer func() { fmt.Fprint(os.Stdout, total/count) }()

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		count++
		f, err := strconv.ParseFloat(strings.TrimSuffix(line, "\n"), 64)
		if err != nil {
			return err
		}

		total += f
	}

	return nil
}
