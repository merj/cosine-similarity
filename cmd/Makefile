BINS=tf idf tfidf cs0 cs1 cs ih

PREFIX?=/usr/local
BINDIR?=${PREFIX}/bin

.SUFFIXES:

.ifndef NOGOLANG
.SUFFIXES: .go
.go:
	go build -o ${.TARGET} ${.IMPSRC}
.endif

.SUFFIXES: .rb
.rb:
	cp ${.IMPSRC} ${.TARGET}

.SUFFIXES: .sh
.sh:
	cp ${.IMPSRC} ${.TARGET}

all: ${BINS}

.PHONY: clean
clean:
.for BIN in ${BINS}
	-rm ${BIN}
.endfor

.PHONY: install
install: all
.for BIN in ${BINS}
	install ${BIN} ${BINDIR}/${BIN}
.endfor

.PHONY: uninstall
uninstall:
.for BIN in ${BINS}
	-rm ${BINDIR}/${BIN}
.endfor
