# Create PDF from HTML in Go

This repository demonstrates how to create PDF files from HTML templates using the Go programming language. The primary library used in this project is `go-wkhtmltopdf`, which is a Go wrapper for the `wkhtmltopdf` command line tool.

## go-wkhtmltopdf

The `go-wkhtmltopdf` library provides an easy-to-use interface to generate PDF files from HTML content. The project's repository can be found here: https://github.com/SebastiaanKlippert/go-wkhtmltopdf

### Dependencies

- wkhtmltopdf

#### Usage

For usage details of the `wkhtmltopdf` command line tool, visit: https://wkhtmltopdf.org/usage/wkhtmltopdf.txt

#### Installation

On macOS, you can install `wkhtmltopdf` using Homebrew:

```bash
$ brew install wkhtmltopdf
```

#### Go Dependencies

When using Go v1.18, you need to install a specific version of go-wkhtmltopdf:

```bash
$ go get github.com/SebastiaanKlippert/go-wkhtmltopdf@v1.7.2
```

#### Docker

The `wkhtmltopdf` package is known to be broken in the Alpine 3.14 image.
Therefore, this project uses the Ubuntu 22.04 image for the Docker container.

#### Getting Started

To run the project locally, first install the required dependencies, then build
and run the Go web server. Alternatively, you can build and run the project
using Docker.

#### Running the Project Locally

Install wkhtmltopdf and the Go dependencies as described above.
Build and run the Go web server:

```bash
$ go build
$ ./main
```

#### Running the Project with Docker

Build the Docker image:

```bash
$ docker build -t go-web-server .
```

Run the Docker container, exposing port 8080:

```bash
$ docker run -p 8080:8080 --name my-go-web-server go-web-server
```
