package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s n a b\n", path.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 4 {
		usage()
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil || n <= 0 {
		usage()
	}

	a, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil || a < 0 {
		usage()
	}

	b, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil || b <= a || b > n {
		usage()
	}

	fn := fmt.Sprintf("h-%d-%d.csv", a, b)
	hf, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(3)
	}
	defer hf.Close()
	fn = fmt.Sprintf("m-%d-%d.csv", a, b)
	mf, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(3)
	}
	defer mf.Close()
	fn = fmt.Sprintf("l-%d-%d.csv", a, b)
	lf, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(3)
	}
	defer lf.Close()

	vis := make([]map[string]float32, n)
	sis := make([]float32, n)

	for i := int64(0); i < n; i++ {
		p := fmt.Sprintf("%d.txt.tfidf", i)
		f, err := os.OpenFile(p, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(3)
		}
		vi := make(map[string]float32)
		si := float32(0.0)
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
			vi[k] = float32(v)
			si += float32(v) * float32(v)
		}
		f.Close()
		vis[i] = vi
		sis[i] = si
	}

	for i := a; i < b; i++ {
		v1 := vis[i]
		s1 := sis[i]
		for j := i + 1; j < n; j++ {
			v2 := vis[j]
			s2 := sis[j]
			m := float32(math.Sqrt(float64(s1)) * math.Sqrt(float64(s2)))
			if m > 0 {
				vs := v1
				vl := v2
				if len(vs) > len(vl) {
					t := vl
					vl = vs
					vs = t
				}
				d := float32(0.0)
				for k, v := range vs {
					d += v * vl[k]
				}
				c := d / m
				if c >= 0.80 {
					fmt.Fprintf(hf, "%d,%d,%0.4f\n", i, j, c)
				} else if c >= 0.60 {
					fmt.Fprintf(mf, "%d,%d,%0.4f\n", i, j, c)
				} else if c >= 0.40 {
					fmt.Fprintf(lf, "%d,%d,%0.4f\n", i, j, c)
				}
			}
		}
	}
}
