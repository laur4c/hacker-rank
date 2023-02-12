package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func buildResponse(queries []string, mapQueryCount map[string]int32) []int32 {
    var result []int32
    for _, query := range queries {
        fmt.Println(query, mapQueryCount[query])
        result = append(result, mapQueryCount[query])
    }
    return result
}

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries
 */

func matchingStrings(stringList []string, queries []string) []int32 {
    n := len(stringList)
    q := len(queries)
    
    mapQueryCount := make(map[string]int32, len(queries))
    
    if n < 1 || n > 1000 || q < 1 || q > 1000 {
        return buildResponse(queries, mapQueryCount)
    }

    evaluated := make(map[string]bool)
    for _, query := range queries {
        if evaluated[query] {
            continue
        }

        // validate len of query
        if len(query) > 20 {
            continue
        }
        // initialize map with query
        if _, present := mapQueryCount[query]; !present {
            mapQueryCount[query] = 0
        }

        for _, str := range stringList {
            // validate len of string
            if len(str) < 1 {
                continue
            }

            if query == str {
                mapQueryCount[query] += 1
            }
        }

        if _, ok := evaluated[query]; !ok {
            evaluated[query] = true
        }
    }
    
    return buildResponse(queries, mapQueryCount)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    stringListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var stringList []string

    for i := 0; i < int(stringListCount); i++ {
        stringListItem := readLine(reader)
        stringList = append(stringList, stringListItem)
    }

    queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var queries []string

    for i := 0; i < int(queriesCount); i++ {
        queriesItem := readLine(reader)
        queries = append(queries, queriesItem)
    }

    res := matchingStrings(stringList, queries)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

