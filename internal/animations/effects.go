// Package animations contains spinner animation generators.
// This is an internal package - not part of the public API.
package animations

import (
	"math"

	"github.com/benoitpetit/agent-spinner/internal/braille"
)

// Generator is a function that creates animation frames.
type Generator func() []string

// Registry maps spinner names to their generators.
var Registry = map[string]Generator{
	"scan":         genScan,
	"rain":         genRain,
	"scanline":     genScanLine,
	"pulse":        genPulse,
	"snake":        genSnake,
	"sparkle":      genSparkle,
	"cascade":      genCascade,
	"columns":      genColumns,
	"orbit":        genOrbit,
	"breathe":      genBreathe,
	"waverows":     genWaveRows,
	"checkerboard": genCheckerboard,
	"helix":        genHelix,
	"fillsweep":    genFillSweep,
	"diagswipe":    genDiagonalSwipe,
	"dots":         genDots,
	"bounce":       genBounce,
	"ripple":       genRipple,
	"bars":         genBars,
	"radar":        genRadar,
	"typing":       genTyping,
	"loading":      genLoading,
	"pingpong":     genPingPong,
	"star":         genStar,
	"arrow":        genArrow,
	"wave":         genWave,
	"progress":     genProgress,
	"circle":       genCircle,
	"cross":        genCross,
	"zigzag":       genZigzag,
	"diamond":      genDiamond,
	"tiles":        genTiles,
	"expand":       genExpand,
	"shrink":       genShrink,
	"scandual":     genScanDual,
	"wavevertical": genWaveVertical,
	"randomfill":   genRandomFill,
	"border":       genBorder,
	"matrix":       genMatrix,
}

// Generate creates frames for a named animation.
func Generate(name string) []string {
	if gen, ok := Registry[name]; ok {
		return gen()
	}
	return nil
}

func genScan() []string {
	const w, h = 8, 4
	var frames []string
	for pos := -1; pos < w+1; pos++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if c == pos || c == pos-1 {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genRain() []string {
	const w, h, totalFrames = 8, 4, 12
	var frames []string
	offsets := []int{0, 3, 1, 5, 2, 7, 4, 6}
	for f := 0; f < totalFrames; f++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			row := (f + offsets[c]) % (h + 2)
			if row < h {
				g[row][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genScanLine() []string {
	const w, h = 6, 4
	var frames []string
	positions := []int{0, 1, 2, 3, 2, 1}
	for _, row := range positions {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			g[row][c] = true
			if row > 0 {
				g[row-1][c] = (c%2 == 0)
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genPulse() []string {
	const w, h = 6, 4
	var frames []string
	cx := float64(w)/2 - 0.5
	cy := float64(h)/2 - 0.5
	radii := []float64{0.5, 1.2, 2, 3, 3.5}
	for _, r := range radii {
		g := braille.MakeGrid(h, w)
		for row := 0; row < h; row++ {
			for col := 0; col < w; col++ {
				dist := math.Hypot(float64(col)-cx, float64(row)-cy)
				if math.Abs(dist-r) < 0.9 {
					g[row][col] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genSnake() []string {
	const w, h = 4, 4
	var path [][2]int
	for r := 0; r < h; r++ {
		if r%2 == 0 {
			for c := 0; c < w; c++ {
				path = append(path, [2]int{r, c})
			}
		} else {
			for c := w - 1; c >= 0; c-- {
				path = append(path, [2]int{r, c})
			}
		}
	}
	var frames []string
	for i := 0; i < len(path); i++ {
		g := braille.MakeGrid(h, w)
		for t := 0; t < 4; t++ {
			idx := (i - t + len(path)) % len(path)
			g[path[idx][0]][path[idx][1]] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genSparkle() []string {
	patterns := [][]int{
		{1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
	}
	const w, h = 8, 4
	var frames []string
	for _, pat := range patterns {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				g[r][c] = pat[r*w+c] == 1
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genCascade() []string {
	const w, h = 8, 4
	var frames []string
	for offset := 0; offset < w+h; offset++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				diag := c + r
				if diag == offset || diag == offset-1 {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genColumns() []string {
	const w, h = 6, 4
	var frames []string
	for col := 0; col < w; col++ {
		for fillTo := h - 1; fillTo >= 0; fillTo-- {
			g := braille.MakeGrid(h, w)
			for pc := 0; pc < col; pc++ {
				for r := 0; r < h; r++ {
					g[r][pc] = true
				}
			}
			for r := fillTo; r < h; r++ {
				g[r][col] = true
			}
			frames = append(frames, braille.GridToBraille(g))
		}
	}
	frames = append(frames, braille.GridToBraille(braille.MakeGrid(h, w)))
	return frames
}

func genOrbit() []string {
	const w, h = 2, 4
	path := [][2]int{
		{0, 0}, {0, 1}, {1, 1}, {2, 1}, {3, 1},
		{3, 0}, {2, 0}, {1, 0},
	}
	var frames []string
	for i := 0; i < len(path); i++ {
		g := braille.MakeGrid(h, w)
		g[path[i][0]][path[i][1]] = true
		t1 := (i - 1 + len(path)) % len(path)
		g[path[t1][0]][path[t1][1]] = true
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genBreathe() []string {
	stages := [][][2]int{
		{},
		{{1, 0}},
		{{0, 1}, {2, 0}},
		{{0, 0}, {1, 1}, {3, 0}},
		{{0, 0}, {1, 1}, {2, 0}, {3, 1}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {3, 1}},
		{{0, 0}, {0, 1}, {1, 0}, {2, 1}, {3, 0}, {3, 1}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {3, 0}, {3, 1}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1}, {3, 0}, {3, 1}},
	}
	var frames []string
	sequence := append(stages, reverseStages(stages[:len(stages)-1])...)
	for _, dots := range sequence {
		g := braille.MakeGrid(4, 2)
		for _, d := range dots {
			g[d[0]][d[1]] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genWaveRows() []string {
	const w, h, totalFrames = 8, 4, 16
	var frames []string
	for f := 0; f < totalFrames; f++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			phase := float64(f) - float64(c)*0.5
			row := int(math.Round((math.Sin(phase*0.8) + 1) / 2 * float64(h-1)))
			g[row][c] = true
			if row > 0 {
				g[row-1][c] = (f+c)%3 == 0
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genCheckerboard() []string {
	const w, h = 6, 4
	var frames []string
	for phase := 0; phase < 4; phase++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if phase < 2 {
					g[r][c] = (r+c+phase)%2 == 0
				} else {
					g[r][c] = (r+c+phase)%3 == 0
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genHelix() []string {
	const w, h, totalFrames = 8, 4, 16
	var frames []string
	for f := 0; f < totalFrames; f++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			phase := float64(f+c) * (math.Pi / 4)
			y1 := int(math.Round((math.Sin(phase) + 1) / 2 * float64(h-1)))
			y2 := int(math.Round((math.Sin(phase+math.Pi) + 1) / 2 * float64(h-1)))
			g[y1][c] = true
			g[y2][c] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genFillSweep() []string {
	const w, h = 4, 4
	var frames []string
	for row := h - 1; row >= 0; row-- {
		g := braille.MakeGrid(h, w)
		for r := row; r < h; r++ {
			for c := 0; c < w; c++ {
				g[r][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	for row := 1; row < h; row++ {
		g := braille.MakeGrid(h, w)
		for r := row; r < h; r++ {
			for c := 0; c < w; c++ {
				g[r][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	frames = append(frames, braille.GridToBraille(braille.MakeGrid(h, w)))
	return frames
}

func genDiagonalSwipe() []string {
	const w, h = 4, 4
	var frames []string
	maxDiag := w + h - 2
	for d := 0; d <= maxDiag; d++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if r+c <= d {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	for d := 1; d <= maxDiag; d++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if r+c > d {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func reverseStages(stages [][][2]int) [][][2]int {
	result := make([][][2]int, len(stages))
	for i := range stages {
		result[i] = stages[len(stages)-1-i]
	}
	return result
}
