package helpers

import (
	"jagratama-backend/internal/config"
	"jagratama-backend/internal/model"
)

// SendRequestDocumentApproveMail sends an email to request document approval to approver
func SendRequestDocumentApproveMail(toEmail, toName, fromName string, document *model.Document) error {
	// subject := "Document Approval Request"
	// body := templateDocumentApprovedMail(toName, fromName, document)

	// return resend.SendMail(toEmail, subject, body)
	return nil
}

func templateDocumentApprovedMail(toName, fromName string, document *model.Document) string {
	return `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Permintaan Approval Dokumen - Jagratama</title>
				<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap" rel="stylesheet">
				<style type="text/css">
					/* Reset CSS */
					body, table, td, a { -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; }
					table, td { mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
					img { -ms-interpolation-mode: bicubic; }

					/* General Styling */
					body {
						margin: 0;
						padding: 0;
						font-family: 'Inter', sans-serif;
						background-color: #f4f7f6;
						color: #333333;
						line-height: 1.6;
					}
					a {
						text-decoration: none;
						color: #1f939c; /* Green tone for links */
					}
					.container {
						max-width: 600px;
						margin: 0 auto;
						background-color: #ffffff;
						border-radius: 12px;
						overflow: hidden;
						box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
					}
					.header {
						background-color: #1f939c; /* Primary Green tone */
						padding: 24px;
						text-align: center;
						border-top-left-radius: 12px;
						border-top-right-radius: 12px;
					}
					.header img {
						max-width: 150px;
						height: auto;
					}
					.content {
						padding: 32px 24px;
					}
					.button {
						display: inline-block;
						background-color: #1f939c; /* Primary Green tone */
						color: #ffffff;
						padding: 14px 28px;
						border-radius: 8px;
						font-weight: 600;
						text-align: center;
						transition: background-color 0.3s ease;
					}
					.button:hover {
						background-color: #166d73; /* Darker green on hover */
					}
					.footer {
						background-color: #e6f5f6; /* Lighter green for footer */
						padding: 24px;
						text-align: center;
						font-size: 12px;
						color: #666666;
						border-bottom-left-radius: 12px;
						border-bottom-right-radius: 12px;
					}
					.footer a {
						color: #1f939c; /* Green tone for footer links */
					}
					.divider {
						border-bottom: 1px solid #eeeeee;
						margin: 24px 0;
					}

					/* Responsive Styles */
					@media only screen and (max-width: 600px) {
						.container {
							width: 100% !important;
							border-radius: 0 !important;
						}
						.header, .content, .footer {
							padding: 20px !important;
						}
						.button {
							padding: 12px 24px !important;
							font-size: 14px !important;
						}
					}
				</style>
			</head>
			<body>
				<table border="0" cellpadding="0" cellspacing="0" width="100%" style="border-collapse: collapse;">
					<tr>
						<td align="center" style="padding: 20px 0;">
							<table border="0" cellpadding="0" cellspacing="0" class="container" width="100%">
								<tr>
									<td class="header">
										<img src="https://placehold.co/150x50/1f939c/ffffff?text=JAGRATAMA" alt="Logo Jagratama">
									</td>
								</tr>
								<tr>
									<td class="content">
										<h2 style="margin-top: 0; margin-bottom: 20px; color: #333333; font-size: 24px; font-weight: 700;">
											Permintaan Approval Dokumen Baru
										</h2>
										<p style="margin-bottom: 15px;">Kepada ` + toName + `</p>
										<p style="margin-bottom: 15px;">
											Anda memiliki permintaan approval dokumen baru dari pengaju di sistem Jagratama.
											Mohon segera tinjau dan berikan persetujuan Anda.
										</p>
										<table border="0" cellpadding="0" cellspacing="0" width="100%" style="margin-bottom: 20px;">
											<tr>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee; width: 30%; font-weight: 600;">Pengaju:</td>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee; width: 70%;">` + fromName + `</td>
											</tr>
											<tr>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee; font-weight: 600;">Judul Dokumen:</td>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee;">` + document.Title + `</td>
											</tr>
											<tr>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee; font-weight: 600;">Tipe Dokumen:</td>
												<td style="padding: 8px 0; border-bottom: 1px solid #eeeeee;">` + document.Category.Name + `</td>
											</tr>
											<tr>
												<td style="padding: 8px 0; font-weight: 600;">Tanggal Pengajuan:</td>
												<td style="padding: 8px 0;">` + document.CreatedAt.Format("02 Jan 2006") + `</td>
											</tr>
										</table>

										<p style="margin-bottom: 25px;">
											Untuk melihat detail dokumen dan memberikan approval, silakan klik tombol di bawah ini:
										</p>

										<table border="0" cellpadding="0" cellspacing="0" width="100%">
											<tr>
												<td align="center" style="padding-bottom: 20px;">
													<a href="` + config.GetEnv("FRONTEND_URL", "") + `/jagratama/actions/documents-to-review" class="button" target="_blank">
														Tinjau Dokumen
													</a>
												</td>
											</tr>
										</table>

										<p style="margin-top: 20px;">Terima kasih atas perhatian dan kerja sama Anda.</p>
										<p>Hormat kami,<br>Tim Jagratama</p>
									</td>
								</tr>
								<tr>
									<td class="footer">
										<p style="margin-top: 0; margin-bottom: 5px;">
											Email ini dikirim secara otomatis oleh sistem Jagratama.
											Mohon tidak membalas email ini.
										</p>
										<p style="margin-bottom: 0;">
											Jika Anda memiliki pertanyaan, silakan hubungi
											<a href="mailto:support@jagratama.ac.id" style="color: #1f939c;">support@jagratama.ac.id</a>.
										</p>
									</td>
								</tr>
							</table>
						</td>
					</tr>
				</table>
			</body>
			</html>
	`
}
