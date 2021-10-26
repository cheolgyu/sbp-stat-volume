package local_log

import (
	"io"
	"os"

	"log"
	"time"
)

func ElapsedTime(tag string, msg string) func() {
	if msg != "" {
		log.Printf("[%s] %s", tag, msg)
	}

	start := time.Now()
	return func() { log.Printf("[%s] Elipsed Time: %s", tag, time.Since(start)) }
}
func OpenLogFile(logfile string) {

	if logfile != "" {
		t := time.Now()
		dirname := t.Format("2006-01-02")
		filename := t.Format("2006-01-02_15_04_05_000Z")
		os.MkdirAll("./logs/"+dirname, 0755)
		lf, err := os.OpenFile("./logs/"+dirname+"/"+filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		multi := io.MultiWriter(lf, os.Stdout)
		log.SetOutput(multi)
	}
}
