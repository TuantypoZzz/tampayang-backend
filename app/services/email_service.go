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

// SendEmailStatusUpdateNotification sends email notification for report status updates
func SendEmailStatusUpdateNotification(recipientName, recipientEmail, reportNumber, newStatus, adminNotes string) {
	// Get configuration from environment variables
	fromName := os.Getenv("EMAIL_SENDER_NAME")
	fromAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Recipient
	to := []string{recipientEmail}

	// Create status-specific message
	var statusMessage string
	var statusColor string
	switch strings.ToLower(newStatus) {
	case "proses":
		statusMessage = "sedang dalam proses penanganan"
		statusColor = "#FFA500" // Orange
	case "selesai":
		statusMessage = "telah selesai ditangani"
		statusColor = "#28A745" // Green
	case "ditolak":
		statusMessage = "ditolak"
		statusColor = "#DC3545" // Red
	case "verifikasi":
		statusMessage = "sedang dalam tahap verifikasi"
		statusColor = "#007BFF" // Blue
	default:
		statusMessage = fmt.Sprintf("diperbarui dengan status: %s", newStatus)
		statusColor = "#6C757D" // Gray
	}

	// Create HTML body for email
	htmlBody := fmt.Sprintf(`
	<html>
	<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
		<div style="max-width: 600px; margin: 0 auto; padding: 20px;">
			<h3 style="color: #2C3E50; border-bottom: 2px solid #3498DB; padding-bottom: 10px;">
				Update Status Laporan - TAMPAYANG
			</h3>
			<p>Halo, <strong>%s</strong>.</p>
			<p>Laporan kerusakan infrastruktur Anda dengan nomor referensi <strong>%s</strong> %s pada %s WIT.</p>

			<div style="background-color: #F8F9FA; padding: 15px; border-left: 4px solid %s; margin: 20px 0;">
				<p style="margin: 0;"><strong>Status Terbaru:</strong>
					<span style="color: %s; font-weight: bold; text-transform: uppercase;">%s</span>
				</p>
			</div>`,
		recipientName,
		reportNumber,
		statusMessage,
		time.Now().Format("02 January 2006, 15:04"),
		statusColor,
		statusColor,
		strings.ToUpper(newStatus),
	)

	// Add admin notes if provided
	if adminNotes != "" {
		htmlBody += fmt.Sprintf(`
			<div style="background-color: #E8F4FD; padding: 15px; border-radius: 5px; margin: 20px 0;">
				<p style="margin: 0 0 10px 0;"><strong>Catatan Admin:</strong></p>
				<p style="margin: 0; font-style: italic;">%s</p>
			</div>`, adminNotes)
	}

	htmlBody += `
			<p>Anda dapat melacak progres laporan Anda melalui menu "Cek Status" di aplikasi atau situs web TAMPAYANG.</p>
			<br>
			<div style="border-top: 1px solid #E9ECEF; padding-top: 20px; margin-top: 30px;">
				<p style="margin: 0;">Hormat kami,</p>
				<p style="margin: 5px 0 0 0;"><strong>Tim Admin TAMPAYANG</strong></p>
			</div>
		</div>
	</body>
	</html>`

	// Combine headers and body for email message
	subject := fmt.Sprintf("Update Status Laporan #%s - %s", reportNumber, strings.ToUpper(newStatus))
	headers := []string{
		fmt.Sprintf("From: %s <%s>", fromName, fromAddress),
		fmt.Sprintf("To: %s", strings.Join(to, ", ")),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
	}
	message := strings.Join(headers, "\r\n") + "\r\n\r\n" + htmlBody

	// Authenticate and send email
	auth := smtp.PlainAuth("", fromAddress, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAddress, to, []byte(message))
	if err != nil {
		log.Printf("ERROR: Gagal mengirim email update status ke %s: %v", recipientEmail, err)
		return
	}

	log.Printf("INFO: Email update status berhasil dikirim ke %s untuk laporan %s dengan status %s", recipientEmail, reportNumber, newStatus)
}
