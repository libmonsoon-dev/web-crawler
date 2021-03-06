package tor_test

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/html"

	"github.com/libmonsoon-dev/web-crawler/http/tor"
	"github.com/libmonsoon-dev/web-crawler/logger/logrus"
)

func TestClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping tor client test in short mode.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	factory, err := tor.NewClientFactory(ctx, logrus.NewFactory())
	if err != nil {
		t.Fatal(err)
	}
	defer factory.Close()

	resp, err := factory.NewClient().Get(ctx, "https://check.torproject.org")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	parsed, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	const wantTitle = "Congratulations. This browser is configured to use Tor."
	if title := getTitle(parsed); title != wantTitle {
		t.Errorf("Got:\n%q\nWant:\n%q", title, wantTitle)
	}
}

func getTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		var title bytes.Buffer
		if err := html.Render(&title, n.FirstChild); err != nil {
			panic(err)
		}
		return strings.TrimSpace(title.String())
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := getTitle(c); title != "" {
			return title
		}
	}
	return ""
}
