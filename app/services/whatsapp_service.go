package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func SendFonnteNotification(recipientName, recipientPhone, reportNumber string) {
	apiUrl := os.Getenv("WHATSAPP_API_URL")
	token := os.Getenv("WHATSAPP_API_KEY")

	var formattedPhone string
	if strings.HasPrefix(recipientPhone, "08") {
		formattedPhone = "62" + strings.TrimPrefix(recipientPhone, "0")
	} else {
		formattedPhone = recipientPhone
	}

	message := fmt.Sprintf(
		"Halo, *%s*.\n\nTerima kasih atas partisipasi Anda. Laporan kerusakan infrastruktur dengan nomor *%s* telah kami terima pada tanggal %s WIT.\n\nAnda dapat melacak progres laporan melalui menu \"Cek Status\" di aplikasi TAMPAYANG.",
		recipientName,
		reportNumber,
		time.Now().Format("02 January 2006, 15:04"),
	)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("target", formattedPhone)
	writer.WriteField("message", message)

	err := writer.Close()
	if err != nil {
		log.Printf("ERROR: Fonnte - Gagal membuat form data: %v\n", err)
		return
	}

	// 4. Membuat HTTP request.
	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		log.Printf("ERROR: Fonnte - Gagal membuat request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: Fonnte - Gagal mengirim notifikasi ke %s: %v\n", formattedPhone, err)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		log.Printf("INFO: Fonnte - Notifikasi berhasil dikirim ke %s. Respons: %s\n", formattedPhone, string(responseBody))
	} else {
		log.Printf("ERROR: Fonnte - Gagal mengirim notifikasi ke %s, Status: %s, Respons: %s\n", formattedPhone, resp.Status, string(responseBody))
	}
}
