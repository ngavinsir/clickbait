package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/ngavinsir/clickbait/model"
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
)

type WriteCounter struct {
	Total uint64
	Size  uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc WriteCounter) printProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

var cmdArticle = &cobra.Command{
	Use:  "article",
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := model.InitDB()
		if err != nil {
			panic(err)
		}
		log.Println("connected to db")
		defer db.Close()

		fileName := ksuid.New().String() + ".csv"

		if err := downloadFile(fileName, args[0]); err != nil {
			fmt.Println(err)
			return
		}

		inputDataset(fileName, db)
	},
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	size, _ := getURLSize(url)

	counter := &WriteCounter{
		Size: size,
	}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	fmt.Print("\n")

	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func getURLSize(url string) (uint64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	return uint64(size), nil
}

func inputDataset(datasetPath string, db *model.DB) {
	csvfile, err := os.Open(datasetPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	articleRepository := &model.ArticleDatastore{DB: db}

	i := 0
	ch, headerMap := processCSV(csvfile)
	for row := range ch {
		i++
		fmt.Printf("\r%s", strings.Repeat(" ", 35))
		fmt.Printf("\rInserting... %d", i)
		articleRepository.InsertArticle(
			context.Background(),
			row[headerMap["Judul"]],
			row[headerMap["URL"]],
			row[headerMap["Tanggal"]],
			row[headerMap["Sumber"]],
			row[headerMap["Konten_pisah"]],
		)
	}
	fmt.Print("\n")

	if err = os.Remove(datasetPath); err != nil {
		log.Fatal(err)
	}
}

func processCSV(rc io.Reader) (ch chan []string, headerMap map[string]uint8) {
	ch = make(chan []string, 10)

	r := csv.NewReader(rc)
	headers, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	headerMap = make(map[string]uint8)
	for i, header := range headers {
		headerMap[header] = uint8(i)
	}

	go func() {
		defer close(ch)

		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)

			}
			ch <- rec
		}
	}()
	return
}
