package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s n\n", path.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil || n <= 0 {
		usage()
	}

	fn := "idf"
	f, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(3)
	}
	idf := make(map[string]float32)
	r := bufio.NewScanner(f)
	for r.Scan() {
		l := strings.Split(r.Text(), ",")
		if len(l) != 2 {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(3)
		}
		k := l[0]
		v, err := strconv.ParseFloat(l[1], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(2)
		}
		idf[k] = float32(v)
	}
	f.Close()

	for i := int64(0); i < n; i++ {
		fn := fmt.Sprintf("%d.txt.tf", i)
		f, err := os.OpenFile(fn, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(3)
		}
		tfidf := make(map[string]float32)
		r := bufio.NewScanner(f)
		for r.Scan() {
			l := strings.Split(r.Text(), ",")
			if len(l) != 2 {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(3)
			}
			k := l[0]
			v, err := strconv.ParseFloat(l[1], 32)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(3)
			}
			/* term frequency * inverse document frequency */
			tfidf[k] = float32(v) * idf[k]
		}
		f.Close()
		fn = fmt.Sprintf("%d.txt.tfidf", i)
		f, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(4)
		}
		for k, v := range tfidf {
			fmt.Fprintf(f, "%s,%0.10f\n", k, v)
		}
		f.Close()
	}
}
