package adventofcode2017

import (
	"os"
	"log"
)

func ReadFile(filename string) string {


   file, err := os.Open(filename) // For read access.
   if err != nil {
      log.Fatal(err)
   }

   var contents string

   if (file != nil) {
      buffer := make([]byte, 1024)
      for {

         count, err := file.Read(buffer)
         if err != nil {
            break
         } else {
            contents += string(buffer[0:count])
         }

      }
   }

   return contents
}