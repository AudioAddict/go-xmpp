package xmpp

import (
	"fmt"
)

type HtmlChat struct {
	Remote    string
	Type      string
	PlainText string
	Html      string
}

// TODO: filter html
func (c *Client) SendHtml(chat HtmlChat) {
	fmt.Fprintf(c.conn, "<message to='%s' type='%s' xml:lang='en'>"+
		"<body>%s</body><html xmlns='http://jabber.org/protocol/xhtml-im'>"+
		"<body xmlns='http://www.w3.org/1999/xhtml'>%s</body></html></message>",
		xmlEscape(chat.Remote), xmlEscape(chat.Type), xmlEscape(chat.PlainText), chat.Html)
}
