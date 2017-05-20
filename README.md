acol
====

[![Build Status][travis-badge]][travis-url]
[![Go Report Card][report-badge]][report-url]
[![MIT License][license-badge]](LICENSE.txt)

Styles standard input in a tabular format like `ls`.

- [Installation](#installation)
- [Testing](#installation)
- [Usage](#usage)
- [Alternatives](#alternatives)

[travis-badge]: https://travis-ci.org/frickiericker/acol.svg?branch=master
[travis-url]: https://travis-ci.org/frickiericker/acol
[report-badge]: https://goreportcard.com/badge/github.com/frickiericker/acol
[report-url]: https://goreportcard.com/report/github.com/frickiericker/acol
[license-badge]: http://img.shields.io/badge/license-MIT-blue.svg

## Installation

[Install Go](https://golang.org/doc/install) and run this command:

    go get github.com/frickiericker/acol

Now acol is placed in GOPATH/bin (which is ~/go/bin by default).

## Testing

Move to the repository root (or ~/go/github.com/frickiericker/acol) and

    go test ./...

## Usage

acol reads list of lines and formats the list into multiple columns. Like this:

    % head -n 30 /usr/share/dict/words | acol
    A      aam       Aaronic    Ab       abaca        abaciscus
    a      Aani      Aaronical  aba      abacate      abacist
    aa     aardvark  Aaronite   Ababdeh  abacay       aback
    aal    aardwolf  Aaronitic  Ababua   abacinate    abactinal
    aalii  Aaron     Aaru       abac     abacination  abactinally

The output is column-major by default. Specifying -r option changes it to
row-major:

    % head -n 30 /usr/share/dict/words | acol -r
    A          a            aa         aal          aalii      aam       Aani
    aardvark   aardwolf     Aaron      Aaronic      Aaronical  Aaronite  Aaronitic
    Aaru       Ab           aba        Ababdeh      Ababua     abac      abaca
    abacate    abacay       abacinate  abacination  abaciscus  abacist   aback
    abactinal  abactinally

Of course, input may contain spaces.

    % sysctl dev.pci | acol
    dev.pci.5.wake: 0              dev.pci.2.wake: 0
    dev.pci.5.%parent: pcib5       dev.pci.2.%parent: pcib2
    dev.pci.5.%pnpinfo:            dev.pci.2.%pnpinfo:
    dev.pci.5.%location:           dev.pci.2.%location:
    dev.pci.5.%driver: pci         dev.pci.2.%driver: pci
    dev.pci.5.%desc: ACPI PCI bus  dev.pci.2.%desc: ACPI PCI bus
    dev.pci.4.wake: 0              dev.pci.1.wake: 0
    dev.pci.4.%parent: pcib4       dev.pci.1.%parent: pcib1
    dev.pci.4.%pnpinfo:            dev.pci.1.%pnpinfo:
    dev.pci.4.%location:           dev.pci.1.%location:
    dev.pci.4.%driver: pci         dev.pci.1.%driver: pci
    dev.pci.4.%desc: ACPI PCI bus  dev.pci.1.%desc: ACPI PCI bus
    dev.pci.3.wake: 0              dev.pci.0.%parent: pcib0
    dev.pci.3.%parent: pcib3       dev.pci.0.%pnpinfo:
    dev.pci.3.%pnpinfo:            dev.pci.0.%location:
    dev.pci.3.%location:           dev.pci.0.%driver: pci
    dev.pci.3.%driver: pci         dev.pci.0.%desc: ACPI PCI bus
    dev.pci.3.%desc: ACPI PCI bus  dev.pci.%parent:

### Options

    -h, --help   display help information
    -r           use row-major ordering
    -s[=2]       space between columns

## Alternatives

The column command ([Linux][linux-column], [BSD][bsd-column]) works similarly to
acol, although the output is less packed than that of acol:

    % head -n 30 /usr/share/dict/words | column
    A               Aani            Aaronite        Ababua          abacination
    a               aardvark        Aaronitic       abac            abaciscus
    aa              aardwolf        Aaru            abaca           abacist
    aal             Aaron           Ab              abacate         aback
    aalii           Aaronic         aba             abacay          abactinal
    aam             Aaronical       Ababdeh         abacinate       abactinally

You can still make it dense by piping the output to `column -t`:

    % head -n 30 /usr/share/dict/words | column | column -t
    A      Aani       Aaronite   Ababua     abacination
    a      aardvark   Aaronitic  abac       abaciscus
    aa     aardwolf   Aaru       abaca      abacist
    aal    Aaron      Ab         abacate    aback
    aalii  Aaronic    aba        abacay     abactinal
    aam    Aaronical  Ababdeh    abacinate  abactinally

While acol packs more words in each row:

    % head -n 30 /usr/share/dict/words | acol
    A      aam       Aaronic    Ab       abaca        abaciscus
    a      Aani      Aaronical  aba      abacate      abacist
    aa     aardvark  Aaronite   Ababdeh  abacay       aback
    aal    aardwolf  Aaronitic  Ababua   abacinate    abactinal
    aalii  Aaron     Aaru       abac     abacination  abactinally

[linux-column]: https://linux.die.net/man/1/column
[bsd-column]: https://www.freebsd.org/cgi/man.cgi?query=column
