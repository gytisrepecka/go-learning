// Copyright 2019 Gytis Repečka (gytis@repecka.com). All rights reserved.
// Use of this source code is governed by a GNU GPL license that can be
// found in the LICENSE file.

package main

import (
  "fmt"
  "time"
  "os"
  "encoding/csv"
  "io"
  // "log"
  // "bufio"
  // "strings"
  "strconv"
  "errors"
)

// Todo: push loaded data to this struct
type carDiagData struct {
  dataTimestamp string
  engineRpm float64
  fuelFlow float64
  intakeAirTemp uint
  speedObd uint
  massAirFlow float64
  throttlePos float64
  voltageObd float64
  engineCoolTemp uint
}

// Format date
func ConvertStrToDate (inStr string) (string, error) {
  // Timestamp layout of data in CSV file
  timestampLayoutFile := "Mon Jan 02 15:04:05 MST 2006"
  // Timestamp layout to output
  timestampLayoutOut := "2006-01-02 15:04:05"
  var timestampOut string

  timestamp, err := time.Parse(timestampLayoutFile, inStr)
  if err != nil {
    return "", errors.New("Can't convert string to date!")
  } else {
    timestampOut = timestamp.Format(timestampLayoutOut)
    return timestampOut, nil
  }
}

// String to float
func ConvertStrToFloat (inStr string) (float64, error) {
  flt, err := strconv.ParseFloat(inStr, 64)
  if err != nil {
    return 0, errors.New("Can't convert string to float!")
  } else {
    return flt, nil
  }
}

// Convert string to integer
func ConvertStrToInt(inStr string) (int, error) {
  intg, err := strconv.Atoi(inStr)
  if err != nil {
    return 0, errors.New("Can't convert string to int!")
  } else {
    return intg, nil
  }
}


func main() {
  currentTime := time.Now()
  inputFileName := "trackLog-sample.csv"

  fmt.Printf("File to process: %s\n", inputFileName)
  fmt.Printf("Started: %s.\n", currentTime.Format("2006-01-02 15:04:05.000 (MST Z07:00)"))
  fmt.Println("--------------------")

  csvFile, err := os.Open(inputFileName)
  // Close file in the end of main
  defer csvFile.Close()
  if err != nil {
    fmt.Printf("There was an error opening file: %s.\n", err)
    return
  }

  // Read CSV file
  fileReader := csv.NewReader(csvFile)
  // Reader options
  // - Field delimiter
  fileReader.Comma = ','
  // - Each record field count must match header (0) record field count
  fileReader.FieldsPerRecord = 0
  // / Reader options

  recordCount := 0

  for {
    // Read one line at one iteration
    record, err := fileReader.Read()

    // If end of file, stop reading file
    if err == io.EOF {
      break
    }

    // If error in record, stop reading file
    if err != nil {
      // log.Fatal(err)
      fmt.Printf("Error in record: %s.\n", err)
      break
    }

    // Perform actions with record
    // Record is an array of fields

    // If it is first record, treat it as header
    if recordCount == 0 {
      fmt.Printf("Fields: %d\n\n", len(record))
      fmt.Printf("\n---------------------------\n")

      // Iterate through record elements (fields)
      fmt.Printf("|")
      // for i := 0; i < len(record); i++ {
      // For experimenting only output 4 fiels
      for i := 0; i < 4; i++ {
        fmt.Printf(" %s |", record[i])
      }

      fmt.Printf("\n---------------------------\n")
    } else {
      // Process data record (not header)


      field0, err := ConvertStrToDate(record[0])
      if err != nil {
        // Instead of returned error output question mark
        fmt.Printf("(?)")
      } else {
        fmt.Printf("(%s) ", field0)
      }

      field1, err := ConvertStrToFloat(record[1])
      if err != nil {
        fmt.Printf("?.?? | ")
      } else {
        fmt.Printf("%.2f | ", field1)
      }

      field2, err := ConvertStrToFloat(record[2])
      if err != nil {
        fmt.Printf("?.?? | ")
      } else {
        fmt.Printf("%.2f | ", field2)
      }

      // Convert string (inputString) to integer (inputInt)
      field3, err := ConvertStrToInt(record[3])
      if err != nil {
        // fmt.Printf("%s", err)
        // Instead of ConvertStrToInt returned error output question mark
        fmt.Printf("? |")
      } else {
        fmt.Printf("%d |", field3)
      }


      fmt.Printf("\n")
      // fmt.Println(record)
    }
    // / Perform actions with record

    // Increase counter
    recordCount++
  }

  fmt.Println("--------------------")
  fmt.Printf("Total records processed: %d.\n", recordCount)

}
