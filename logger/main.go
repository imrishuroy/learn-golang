package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"Id":  "11111",
			"Vin": "4Y1SL65848Z411439",
		},
	).Info("Logging at Info")

	log.SetLevel(log.DebugLevel)
	log.WithFields(
		log.Fields{
			"Id":  "22222",
			"Vin": "4Y1SL65848Z411439",
		},
	).Debug("Logging at Debug")

	log.WithFields(
		log.Fields{
			"Id":  "33333",
			"Vin": "4Y1SL65848Z411439",
		},
	).Warn("Logging at Warn")
}

//
//func main() {
//	logFile, err := os.Create("app.log")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func(logFile *os.File) {
//		err := logFile.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(logFile)
//	log.SetOutput(logFile)
//	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
//
//	log.Println("Normal log message as Info")
//	log.Printf("Debug message: %s", "some debug info")
//	log.Println("Fatal error message")
//
//	//logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Lmicroseconds|log.Ltime)
//	//
//	//logger.Println("Normal log message as Info")
//	//
//	//logger.Fatalln("fatal message")
//	//logger.Panicln("panic message")
//}

// custom logger
/*

var (
 WarningLogger *log.Logger
 InfoLogger    *log.Logger
 ErrorLogger   *log.Logger
 DebugLogger   *log.Logger
)

func init() {
 file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
 if err != nil {
  log.Fatal(err)
 }

 InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
 WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
 ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
 DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
 InfoLogger.Println("Starting the application...")
 InfoLogger.Println("Something Info Level to Log")
 DebugLogger.Printf("Debug message: %s", "Some debug Info")
 WarningLogger.Println("There is something you should be warned about")
 ErrorLogger.Println("We have an error")
}

*/
