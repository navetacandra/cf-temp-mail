[Baca dalam Bahasa Indonesia](./README.id.md)

# CF Temp Mail

A minimalist and efficient backend system for a disposable/temporary email service, built on Cloudflare and Go.

## About This Project

This project provides a solution for creating a temporary email service without the need to manage complex and expensive email servers. By leveraging Cloudflare Workers and Email Routing, this system can catch incoming emails and display them in real-time.

The goal is to provide a backend foundation that allows users to receive emails on their own domain and view their content instantly, suitable for registration, testing, or privacy purposes.

## How It Works

The system's workflow is simple and efficient:

1.  **Email Reception:** Cloudflare Email Routing is configured to catch all emails sent to the target domain.
2.  **Edge Processing:** Each incoming email triggers a Cloudflare Worker. This worker is responsible for parsing the email content to extract key information such as sender, subject, and the HTML body.
3.  **Webhook Forwarding:** After processing, the Worker sends the email data in JSON format to a backend endpoint via an HTTP POST request (webhook).
4.  **Backend Display:** A lightweight Go server receives the data from the webhook and immediately displays it in the server's console/terminal.

## Technical Stack

*   **Edge Logic:** Cloudflare Worker (JavaScript)
*   **Backend:** HTTP Server (Go)
*   **Email Infrastructure:** Cloudflare Email Routing