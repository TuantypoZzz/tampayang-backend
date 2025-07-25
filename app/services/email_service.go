package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

// SendEmailNotification mengirim email konfirmasi laporan.
func SendEmailNotification(recipientName, recipientEmail, reportNumber string) {
	// Ambil konfigurasi dari getenv
	fromName := os.Getenv("EMAIL_SENDER_NAME")
	fromAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Penerima
	to := []string{recipientEmail}

	// Buat pesan HTML untuk body email
	htmlBody := fmt.Sprintf(`
	<html>
	<body>
		<h3>Laporan Diterima - TAMPAYANG</h3>
		<p>Halo, <strong>%s</strong>.</p>
		<p>Terima kasih atas partisipasi Anda. Laporan kerusakan infrastruktur dengan nomor referensi <strong>%s</strong> telah berhasil kami terima pada %s WIT.</p>
		<p>Anda dapat melacak progres laporan Anda melalui menu "Cek Status" di aplikasi atau situs web TAMPAYANG.</p>
		<br>
		<p>Hormat kami,</p>
		<p><strong>Tim Admin TAMPAYANG</strong></p>
	</body>
	</html>`, recipientName, reportNumber, time.Now().Format("02 January 2006, 15:04"))

	// Gabungkan header dan body untuk pesan email
	subject := fmt.Sprintf("Konfirmasi Laporan Diterima #%s", reportNumber)
	headers := []string{
		fmt.Sprintf("From: %s <%s>", fromName, fromAddress),
		fmt.Sprintf("To: %s", strings.Join(to, ", ")),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
	}
	message := strings.Join(headers, "\r\n") + "\r\n\r\n" + htmlBody

	// Otentikasi dan kirim email
	auth := smtp.PlainAuth("", fromAddress, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAddress, to, []byte(message))
	if err != nil {
		log.Printf("ERROR: Gagal mengirim email notifikasi ke %s: %v", recipientEmail, err)
		return
	}

	log.Printf("INFO: Email notifikasi berhasil dikirim ke %s untuk laporan %s", recipientEmail, reportNumber)
}
