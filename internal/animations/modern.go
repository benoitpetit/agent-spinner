package animations

import (
	"math"

	"github.com/benoitpetit/agent-spinner/internal/braille"
)

func genDots() []string {
	const w, h = 6, 4
	patterns := [][]int{
		{1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1},
		{0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
		{1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0},
	}
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

func genBounce() []string {
	const w, h = 4, 4
	var frames []string
	positions := [][2]int{
		{0, 1}, {1, 1}, {2, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {0, 1},
	}
	for _, pos := range positions {
		g := braille.MakeGrid(h, w)
		row, col := pos[0], pos[1]
		g[row][col] = true
		g[row][col+1] = true
		if row == 3 {
			g[3][0] = true
			g[3][3] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genRipple() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	for radius := 0.5; radius <= 3.5; radius += 0.5 {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				dist := math.Hypot(float64(c)-centerC, float64(r)-centerR)
				if math.Abs(dist-radius) < 0.6 {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genBars() []string {
	const w, h = 6, 4
	var frames []string
	for phase := 0; phase < 8; phase++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			height := ((c + phase) % (h + 1))
			for r := h - height; r < h && r >= 0; r++ {
				g[r][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genRadar() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	for angle := 0.0; angle < 2*math.Pi; angle += math.Pi / 4 {
		g := braille.MakeGrid(h, w)
		for dist := 0.0; dist < 2.5; dist += 0.3 {
			c := int(math.Round(centerC + dist*math.Cos(angle)))
			r := int(math.Round(centerR + dist*math.Sin(angle)))
			if r >= 0 && r < h && c >= 0 && c < w {
				g[r][c] = true
			}
		}
		g[int(centerR)][int(centerC)] = true
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genTyping() []string {
	const w, h = 8, 4
	var frames []string
	for pos := 0; pos < w; pos++ {
		g := braille.MakeGrid(h, w)
		g[h-1][pos] = true
		for c := 0; c < pos; c++ {
			g[h-1][c] = true
			if c%2 == 0 {
				g[h-2][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	for i := 0; i < 4; i++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			g[h-1][c] = true
			if c%2 == 0 {
				g[h-2][c] = true
			}
		}
		if i%2 == 0 {
			g[h-1][w-1] = false
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genLoading() []string {
	const w, h = 6, 4
	var frames []string
	for progress := 1; progress <= w*h; progress++ {
		g := braille.MakeGrid(h, w)
		count := 0
		for r := h - 1; r >= 0 && count < progress; r-- {
			for c := 0; c < w && count < progress; c++ {
				g[r][c] = true
				count++
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genPingPong() []string {
	const w, h = 6, 4
	var frames []string
	positions := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1}
	for _, pos := range positions {
		g := braille.MakeGrid(h, w)
		g[1][0] = true
		g[2][0] = true
		g[2][pos] = true
		g[1][w-1] = true
		g[2][w-1] = true
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genStar() []string {
	const w, h = 6, 4
	var frames []string
	for angle := 0.0; angle < 2*math.Pi; angle += math.Pi / 3 {
		g := braille.MakeGrid(h, w)
		centerC, centerR := float64(w-1)/2, float64(h-1)/2
		for i := 0; i < 4; i++ {
			a := angle + float64(i)*math.Pi/2
			for dist := 0.0; dist < 2.5; dist += 0.5 {
				c := int(math.Round(centerC + dist*math.Cos(a)))
				r := int(math.Round(centerR + dist*math.Sin(a)))
				if r >= 0 && r < h && c >= 0 && c < w {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genArrow() []string {
	const w, h = 6, 4
	patterns := [][][]int{
		{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 0}, {0, 0, 1, 0, 0, 0}},
		{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 0, 1, 1, 1, 0}, {0, 0, 0, 1, 0, 0}},
		{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 0}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 0, 1, 0}},
		{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 0, 1, 1, 1, 0}, {0, 0, 0, 1, 0, 0}},
	}
	var frames []string
	for _, pat := range patterns {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				g[r][c] = pat[r][c] == 1
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genWave() []string {
	const w, h, totalFrames = 8, 4, 12
	var frames []string
	for f := 0; f < totalFrames; f++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			phase := float64(f+c) * 0.5
			row := int(math.Round((math.Sin(phase) + 1) / 2 * float64(h-1)))
			g[row][c] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genProgress() []string {
	const w, h = 8, 4
	var frames []string
	for filled := 1; filled <= w; filled++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < filled; c++ {
			for r := 1; r < h-1; r++ {
				g[r][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genCircle() []string {
	const w, h = 4, 4
	var frames []string
	centerC, centerR := 1.5, 1.5
	for endAngle := math.Pi / 4; endAngle < 2*math.Pi; endAngle += math.Pi / 4 {
		g := braille.MakeGrid(h, w)
		for angle := 0.0; angle <= endAngle; angle += 0.1 {
			r := int(math.Round(centerR + 1.5*math.Sin(angle)))
			c := int(math.Round(centerC + 1.5*math.Cos(angle)))
			if r >= 0 && r < h && c >= 0 && c < w {
				g[r][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	full := braille.MakeGrid(h, w)
	full[0][0], full[1][0], full[2][0], full[0][1] = true, true, true, true
	full[3][0], full[3][1] = true, true
	full[0][2], full[0][3], full[1][3], full[2][3] = true, true, true, true
	full[3][2], full[3][3] = true, true
	frames = append(frames, braille.GridToBraille(full))
	return frames
}

func genCross() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	for angle := 0.0; angle < math.Pi; angle += math.Pi / 8 {
		g := braille.MakeGrid(h, w)
		for dist := -2.5; dist <= 2.5; dist += 0.5 {
			c1 := int(math.Round(centerC + dist*math.Cos(angle)))
			r1 := int(math.Round(centerR + dist*math.Sin(angle)))
			c2 := int(math.Round(centerC + dist*math.Cos(angle+math.Pi/2)))
			r2 := int(math.Round(centerR + dist*math.Sin(angle+math.Pi/2)))
			if r1 >= 0 && r1 < h && c1 >= 0 && c1 < w {
				g[r1][c1] = true
			}
			if r2 >= 0 && r2 < h && c2 >= 0 && c2 < w {
				g[r2][c2] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genZigzag() []string {
	const w, h = 8, 4
	var frames []string
	for offset := 0; offset < 4; offset++ {
		g := braille.MakeGrid(h, w)
		for c := 0; c < w; c++ {
			row := ((c + offset) % 4)
			if row < h {
				g[row][c] = true
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genDiamond() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	sizes := []float64{1.0, 1.5, 2.0, 1.5, 1.0}
	for _, size := range sizes {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				dist := math.Abs(float64(c)-centerC) + math.Abs(float64(r)-centerR)
				if math.Abs(dist-size) < 0.6 {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genTiles() []string {
	const w, h = 6, 4
	var frames []string
	for phase := 0; phase < 4; phase++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if (r+c+phase)%4 == 0 {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genExpand() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	for maxDist := 0.8; maxDist <= 3; maxDist += 1.1 {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				dist := math.Max(math.Abs(float64(c)-centerC), math.Abs(float64(r)-centerR))
				if dist <= maxDist {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genShrink() []string {
	const w, h = 6, 4
	var frames []string
	centerC, centerR := float64(w-1)/2, float64(h-1)/2
	for minDist := 2.2; minDist >= 0; minDist -= 1.1 {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				dist := math.Max(math.Abs(float64(c)-centerC), math.Abs(float64(r)-centerR))
				if dist >= minDist {
					g[r][c] = true
				}
			}
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genScanDual() []string {
	const w, h = 8, 4
	var frames []string
	for pos := 0; pos < w/2; pos++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			g[r][pos] = true
			g[r][w-1-pos] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genWaveVertical() []string {
	const w, h, totalFrames = 6, 4, 12
	var frames []string
	for f := 0; f < totalFrames; f++ {
		g := braille.MakeGrid(h, w)
		for r := 0; r < h; r++ {
			phase := float64(f+r) * 0.5
			col := int(math.Round((math.Sin(phase) + 1) / 2 * float64(w-1)))
			g[r][col] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genRandomFill() []string {
	const w, h = 6, 4
	order := []int{0, 23, 5, 18, 10, 3, 15, 8, 21, 12, 1, 19, 6, 14, 11, 22, 4, 17, 9, 2, 20, 7, 13, 16}
	var frames []string
	g := braille.MakeGrid(h, w)
	for _, idx := range order {
		r := idx / w
		c := idx % w
		if r < h && c < w {
			g[r][c] = true
			frames = append(frames, braille.GridToBraille(braille.CloneGrid(g)))
		}
	}
	return frames
}

func genBorder() []string {
	const w, h = 6, 4
	var frames []string
	borderCells := [][2]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5},
		{1, 5}, {2, 5}, {3, 5},
		{3, 4}, {3, 3}, {3, 2}, {3, 1}, {3, 0},
		{2, 0}, {1, 0},
	}
	for offset := 0; offset < len(borderCells); offset++ {
		g := braille.MakeGrid(h, w)
		for i := 0; i < 4; i++ {
			idx := (offset + i) % len(borderCells)
			cell := borderCells[idx]
			g[cell[0]][cell[1]] = true
		}
		frames = append(frames, braille.GridToBraille(g))
	}
	return frames
}

func genMatrix() []string {
	const w, h = 8, 4
	patterns := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
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
