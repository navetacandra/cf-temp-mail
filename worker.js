function getSimpleBody(raw = "") {
  try {
    const bodyStart = raw.indexOf("\r\n\r\n");
    if (bodyStart < 0) return "";

    const body = raw.slice(bodyStart + 4).trim();
    const bPos = raw.indexOf('boundary="');
    if (bPos < 0) return body;

    const bEnd = raw.indexOf('"', bPos + 10);
    if (bEnd < 0) return body;

    const token = "--" + raw.slice(bPos + 10, bEnd);
    const parts = body.split(token);

    for (let i = 0, p; i < parts.length; i++) {
      p = parts[i];
      const h = p.indexOf("Content-Type: text/html");
      if (h < 0) continue;

      const s = p.indexOf("\r\n\r\n", h);
      if (s < 0) continue;

      let c = p.slice(s + 4).trimEnd();
      if (c.endsWith("--")) c = c.slice(0, -2).trimEnd();
      return c;
    }
    return null;
  } catch {
    return null;
  }
}

const WEBHOOK_ENDPOINT = "";
const WORKER_ID = "";
const TMPMAIL_DOMAIN = "".replace(/\./g, "\\.");

export default {
  async email(msg) {
    if (!new RegExp(`^[a-z0-9.]{3,}@${TMPMAIL_DOMAIN}$`).test(msg.to))
      return msg.setReject(
        `<h1>Recipient <b>&lt;${msg.to}&gt;<b> not found!</h1>`,
      );

    const h = msg.headers;
    const raw = await new Response(msg.raw).text();
    const body = getSimpleBody(raw) || "<h1>Failed to parse body!</h1>";

    return fetch(`${WEBHOOK_ENDPOINT}/mail-worker-webhook`, {
      method: "POST",
      headers: { "Content-Type": "application/json", "X-Worker-Id": WORKER_ID },
      body: JSON.stringify({
        id: h.get("message-id"),
        from: msg.from,
        to: msg.to,
        subject: h.get("subject"),
        body,
      }),
    });
  },
};
