package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
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

	for i := int64(0); i < n; i++ {
		fn := fmt.Sprintf("%d.txt", i)
		f, err := os.OpenFile(fn, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(3)
		}
		tf := make(map[string]float32)
		s := float32(0.0)
		r := bufio.NewScanner(f)
		r.Split(bufio.ScanWords)
		for r.Scan() {
			k := r.Text()
			tf[k]++
			s++
		}
		f.Close()
		fn = fmt.Sprintf("%d.txt.tf", i)
		f, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(4)
		}
		for k, v := range tf {
			/* term frequency adjusted for document length */
			v /= s
			fmt.Fprintf(f, "%s,%0.10f\n", k, v)
		}
		f.Close()
	}
}
