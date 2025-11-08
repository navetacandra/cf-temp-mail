[Read in English](./README.md)

# CF Temp Mail

Sistem backend untuk layanan email sementara (disposable/temporary email) yang minimalis dan efisien, dibangun di atas infrastruktur Cloudflare dan Go.

## Tentang Proyek Ini

Proyek ini menyediakan solusi untuk membuat layanan email sementara tanpa perlu mengelola server email yang rumit dan mahal. Dengan memanfaatkan Cloudflare Workers dan Email Routing, sistem ini dapat menangkap email yang masuk dan menampilkannya secara *real-time*.

Tujuannya adalah menyediakan fondasi backend yang memungkinkan pengguna menerima email di domain mereka sendiri dan melihat isinya secara instan, cocok untuk keperluan registrasi, pengujian, atau menjaga privasi.

## Cara Kerja

Alur kerja sistem ini sederhana dan efisien:

1.  **Penerimaan Email:** Cloudflare Email Routing dikonfigurasi untuk menangkap semua email yang dikirim ke domain target.
2.  **Pemrosesan di Edge:** Setiap email yang masuk akan memicu Cloudflare Worker. Worker ini bertugas mem-parsing konten email untuk mengekstrak informasi penting seperti pengirim, subjek, dan isi email (HTML).
3.  **Penerusan via Webhook:** Setelah diproses, Worker mengirimkan data email dalam format JSON ke sebuah endpoint backend melalui HTTP POST request (webhook).
4.  **Tampilan di Backend:** Sebuah server Go yang ringan menerima data dari webhook dan langsung menampilkannya di konsol/terminal server.

## Komponen Teknis

*   **Edge Logic:** Cloudflare Worker (JavaScript)
*   **Backend:** Server HTTP (Go)
*   **Infrastruktur Email:** Cloudflare Email Routing
