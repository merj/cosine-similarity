package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/merj/gokogiri"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s n\n", path.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	re := regexp.MustCompile(`[\s\W[:space:]]+`)

	for i := 0; i < n; i++ {
		buf, err := ioutil.ReadFile(fmt.Sprintf("%d.html", i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		doc, err := gokogiri.ParseHtml(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		for _, blacklist := range []string{"title", "script", "style"} {
			nodeRm, _ := doc.Search(fmt.Sprintf("//%s/text()", blacklist))
			for j := range nodeRm {
				nodeRm[j].Remove()
			}
		}
		ns, err := doc.Search("//text()")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		var s string
		for _, n := range ns {
			s += n.String()
		}
		doc.Free()

		s = gsub(re, s)
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		ioutil.WriteFile(fmt.Sprintf("%d.txt", i), []byte(s), 0777)
	}
}

func gsub(re *regexp.Regexp, str string) string {
	s := ""
	i := 0
	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		g := []string{}
		for i := 0; i < len(v); i += 2 {
			g = append(g, str[v[i]:v[i+1]])
		}
		s += str[i:v[0]] + " "
		i = v[1]
	}
	return s + str[i:]
}
