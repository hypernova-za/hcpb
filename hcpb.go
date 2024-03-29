package hcpb

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	pieces  = 40
	refresh = 0.2
)

type ProgressBar struct {
	started    time.Time
	max        int
	pos        int
	lastDraw time.Time
	longestLine int
}

func (b *ProgressBar) Inc() {
	b.pos += 1
	b.draw()
}

func New(max int) *ProgressBar {
	b := &ProgressBar{}
	b.started = time.Now()
	b.max = max
	b.pos = 0
	b.lastDraw = time.Now().AddDate(0, 0, -1)
	b.longestLine = 0
	return b
}

func (b *ProgressBar) draw() {
	secs := time.Since(b.lastDraw).Seconds()
	if secs >= refresh || b.max == b.pos {
		b.lastDraw = time.Now()
		elapsed := time.Since(b.started).Seconds()
		perSecond := 0.0
		if elapsed != 0 {
			perSecond = float64(b.pos) / elapsed
		}
		donePieces := int(math.Round(float64(b.pos) / float64(b.max) * pieces))
		var line string
		line = " ["
		for i := 1; i <= pieces; i++ {
			if donePieces < i {
				line += string(" ")
			} else {
				line += string(">")
			}
		}
		line += string("]")

		e := secondsToMinutes(int(elapsed))
		remaining := math.Round((float64(b.max) - float64(b.pos)) / perSecond)
		r := secondsToMinutes(int(remaining))
		if len(r) > 6 {
			r = "~"
		}

		out := fmt.Sprintf("\r%s %s/%s @ %.0f/s in %s ETA %s\r", line, toHumanReadable(b.pos), toHumanReadable(b.max), perSecond, e, r)
		if len(out) > b.longestLine {
			b.longestLine = len(out)
		}
		for {
			if len(out) < b.longestLine {
				out += " "
			} else {
				break
			}

		}
		fmt.Print(out)
	}
}

func toHumanReadable(n int) string {
	v := 0.0
	switch {
	case n >= 1000000000:
		// Billions
		v = float64(n) / 1000000000
		return fmt.Sprintf("%.1fb", v)
	case n >= 1000000:
		// Millions
		v = float64(n) / 1000000
		return fmt.Sprintf("%.1fm", v)
	case n >= 1000:
		// Millions
		v = float64(n) / 1000
		return fmt.Sprintf("%.1fk", v)
	default:
		return strconv.Itoa(n)
	}
}

func secondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}
