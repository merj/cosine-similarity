package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s o n m p\n", path.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 5 {
		usage()
	}

	o, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil || o < 0 {
		usage()
	}

	n, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil || n <= 0 {
		usage()
	}

	m, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil || m <= 0 || m > n {
		usage()
	}

	p, err := strconv.ParseInt(os.Args[4], 10, 32)
	if err != nil || p < m || p > n {
		usage()
	}

	for i := int64(0); i < n; i++ {
		p := fmt.Sprintf("%d.txt.tfidf", i)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(3)
		}
	}

	csns := make([]*exec.Cmd, m)
	pc := 0
	pp := 0
	pq := 0
	ps := 0

	i := int64(0)
	j := int64(math.Ceil(float64(float32(n) / float32(p))))
	for i < n {
		if int64(pc) < m {
			k := i + j
			a := i
			b := int64(math.Min(float64(k), float64(n)))
			c := fmt.Sprintf("cs%d %d %d %d", o, n, a, b)
			d := strings.Replace(c, " ", "_", -1)
			csn := exec.Command("/bin/sh", "-c", c)
			fn := fmt.Sprintf("%s.out", d)
			of, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(4)
			}
			csn.Stdout = of
			fn = fmt.Sprintf("%s.err", d)
			ef, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(4)
			}
			csn.Stderr = ef
			err = csn.Start()
			if err == nil {
				of.Close()
				ef.Close()
				csns[pq%cap(csns)] = csn
				pq++
				fmt.Fprintln(os.Stderr, c)
				pc++
				i = k
			}
		}
		if int64(pc) == m {
			csn := csns[pp%cap(csns)]
			pp++
			err := csn.Wait()
			if err != nil {
				ps = 5
			}
			pc--
		}
	}

	for int64(pc) > 0 {
		csn := csns[pp%cap(csns)]
		pp++
		err := csn.Wait()
		if err != nil {
			ps = 5
		}
		pc--
	}

	os.Exit(ps)
}
