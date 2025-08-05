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

func SendFonnteNotification(recipientName, recipientPhone, reportNumber, villageName string) {
	apiUrl := os.Getenv("WHATSAPP_API_URL")
	token := os.Getenv("WHATSAPP_API_KEY")

	var formattedPhone string
	if strings.HasPrefix(recipientPhone, "08") {
		formattedPhone = "62" + strings.TrimPrefix(recipientPhone, "0")
	} else {
		formattedPhone = recipientPhone
	}

	message := fmt.Sprintf(
		"Halo, *%s*.\n\nTerima kasih atas partisipasi Anda. Laporan kerusakan infrastruktur dengan nomor *%s* pada desa *%s* telah kami terima pada tanggal %s WIT.\n\nAnda dapat melacak progres laporan melalui menu \"Cek Status\" di aplikasi TAMPAYANG.",
		recipientName,
		reportNumber,
		villageName,
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

// SendFonnteStatusUpdateNotification sends WhatsApp notification for report status updates
func SendFonnteStatusUpdateNotification(recipientName, recipientPhone, reportNumber, newStatus, adminNotes, villageName string) {
	apiUrl := os.Getenv("WHATSAPP_API_URL")
	token := os.Getenv("WHATSAPP_API_KEY")

	var formattedPhone string
	if strings.HasPrefix(recipientPhone, "08") {
		formattedPhone = "62" + strings.TrimPrefix(recipientPhone, "0")
	} else {
		formattedPhone = recipientPhone
	}

	// Create status-specific message
	var statusMessage string
	switch strings.ToLower(newStatus) {
	case "proses":
		statusMessage = "sedang dalam proses penanganan"
	case "selesai":
		statusMessage = "telah selesai ditangani"
	case "ditolak":
		statusMessage = "ditolak"
	case "verifikasi":
		statusMessage = "sedang dalam tahap verifikasi"
	default:
		statusMessage = fmt.Sprintf("diperbarui dengan status: %s", newStatus)
	}

	message := fmt.Sprintf(
		"Halo, *%s*.\n\nLaporan kerusakan infrastruktur Anda dengan nomor *%s* pada desa *%s* %s pada tanggal %s WIT.",
		recipientName,
		reportNumber,
		villageName,
		statusMessage,
		time.Now().Format("02 January 2006, 15:04"),
	)

	// Add admin notes if provided
	if adminNotes != "" {
		message += fmt.Sprintf("\n\n*Catatan Admin:*\n%s", adminNotes)
	}

	message += "\n\nAnda dapat melacak progres laporan melalui menu \"Cek Status\" di aplikasi TAMPAYANG."

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("target", formattedPhone)
	writer.WriteField("message", message)

	err := writer.Close()
	if err != nil {
		log.Printf("ERROR: Fonnte Status Update - Gagal membuat form data: %v\n", err)
		return
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		log.Printf("ERROR: Fonnte Status Update - Gagal membuat request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: Fonnte Status Update - Gagal mengirim notifikasi ke %s: %v\n", formattedPhone, err)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		log.Printf("INFO: Fonnte Status Update - Notifikasi berhasil dikirim ke %s. Respons: %s\n", formattedPhone, string(responseBody))
	} else {
		log.Printf("ERROR: Fonnte Status Update - Gagal mengirim notifikasi ke %s, Status: %s, Respons: %s\n", formattedPhone, resp.Status, string(responseBody))
	}
}
