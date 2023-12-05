package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: ./wget <URL>")
		return
	}

	url := os.Args[1]
	downloadSite(url)
}

func downloadSite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении GET-запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: неверный статус ответа:", resp.Status)
		return
	}

	baseURL := getBaseURL(url)
	basePath := getBasePath(url)

	err = os.Mkdir(basePath, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}

	parseHTML(resp.Body, baseURL, basePath)
}

func getBaseURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[0] + "//" + parts[2]
}

func getBasePath(url string) string {
	parts := strings.Split(url, "/")
	return parts[2]
}

func parseHTML(reader io.Reader, baseURL, basePath string) {
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "a" || token.Data == "link" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if strings.HasPrefix(link, "/") {
							link = baseURL + link
						}

						downloadFile(link, basePath)
					}
				}
			}
		}
	}
}

func downloadFile(url, basePath string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении GET-запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: неверный статус ответа:", resp.Status)
		return
	}

	filePath := basePath + "/" + getFileName(url)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return
	}

	fmt.Println("Скачан файл:", filePath)
}

func getFileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
