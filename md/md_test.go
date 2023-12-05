package md

import (
	"fmt"
	"testing"
)

func TestRenderString(t *testing.T) {
	for _, tc := range []struct {
		text string
		want string
	}{
		{
			text: "a",
			want: fmt.Sprintln("<p>a</p>"),
		},
		{
			text: "http://gnu.org",
			want: fmt.Sprintln(`<p><a href="http://gnu.org" target="_blank" rel="nofollow noopener">http://gnu.org</a></p>`),
		},
		{
			text: "gnu.org",
			want: fmt.Sprintln(`<p><a href="https://gnu.org" target="_blank" rel="nofollow noopener">gnu.org</a></p>`),
		},
		{
			text: "the link is https://gnu.org, and another link is https://example.com/abc",
			want: fmt.Sprintln(`<p>the link is <a href="https://gnu.org" target="_blank" rel="nofollow noopener">https://gnu.org</a>, and another link is <a href="https://example.com/abc" target="_blank" rel="nofollow noopener">https://example.com/abc</a></p>`),
		},
	} {
		t.Run(tc.text, func(t *testing.T) {
			got := RenderString(tc.text)

			if string(got) != tc.want {
				t.Errorf("expected:\n'%s', got:\n'%s'\n", tc.want, got)
			}
		})
	}
}
