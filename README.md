# Human Duration [![Build Status](https://travis-ci.org/davidbanham/human_duration.svg?branch=master)](https://travis-ci.org/davidbanham/human_duration) [![GoDoc](https://godoc.org/github.com/davidbanham/human_duration?status.svg)](https://godoc.org/github.com/davidbanham/human_duration)
A little Go util to print duration strings in a human-friendly format

## Docs

The [String](https://godoc.org/github.com/davidbanham/human_duration#String) function takes a Duration and the precision that's important to the user.

The allowed precisions are year, day, hour, minute and second

## Usage

```go
import "github.com/davidbanham/human_duration"

example := time.Hour * 25 + time.Minute * 4 + time.Second * 8

fmt.Println(human_duration.String(example, "second")) // 1 day 4 minutes 8 seconds
fmt.Println(human_duration.String(example, "minute")) // 1 day 4 minutes
fmt.Println(human_duration.String(example, "day"))    // 1 day

day := time.Hour * 24
year := day * 365

longExample := year * 4 + day * 2

fmt.Println(human_duration.String(longExample, "second")) // 4 years 2 days
```

There are more examples in the [tests](https://github.com/davidbanham/human_duration/blob/master/main_test.go).

## Credit

Adapted and extended from (this gist)[https://gist.github.com/harshavardhana/327e0577c4fed9211f65]
