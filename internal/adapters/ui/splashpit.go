// Copyright 2025.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

// A tiny mosh-pit physics toy for the splash screen: a crowd of ASCII moshers
// swirl in a circle, shove each other, and bounce off the pit edge — a nod to
// the app name. All positions are tracked in a unit disc (radius 1) so the
// physics stay perfectly circular, then mapped to the character grid on render.
const (
	moshPitW  = 21 // pit grid width, in cells
	moshPitH  = 7  // pit grid height, in cells
	moshCount = 11 // number of moshers — one per character of the IP address
)

// introHoldFrames is how long the static IP address is held before the
// moshers start shoving. Brief on purpose — long enough to read the IP, short
// enough that motion begins before anyone reaches for a key.
const introHoldFrames = 6 // ~0.3s at the splash frame rate

// rampDuration is how long the pit takes to build from the first calm shove to
// full chaos — a visible escalation rather than an instant explosion.
const rampDuration = 32 // ~1.75s at the splash frame rate

// moshGlyphs are the per-mosher characters — the digits and dots of an IP
// address, so the crowd reads as "192.168.1.1" before it breaks into the mosh.
var moshGlyphs = []rune("192.168.1.1")

type mosher struct {
	x, y      float64 // position in the unit disc
	vx, vy    float64 // velocity
	glyph     rune
	fallTimer int // frames remaining on the ground; 0 means upright
}

type moshSim struct {
	moshers     []mosher
	rng         *rand.Rand
	introFrames int // frames left holding the static IP address; 0 once moshing
	rampFrames  int // frames left ramping from calm to full chaos; 0 once full
}

// newMoshSim seeds a pit of moshers, one per character of the IP address,
// laid out as a static in-order line until kickoff breaks them into the mosh.
func newMoshSim() *moshSim {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // visual jitter, not security
	s := &moshSim{rng: rng, moshers: make([]mosher, moshCount), introFrames: introHoldFrames}
	for i := range s.moshers {
		s.moshers[i].glyph = moshGlyphs[i%len(moshGlyphs)]
		s.moshers[i].x = -0.82 + float64(i)*(1.64/float64(moshCount-1))
		s.moshers[i].y = 0
	}
	return s
}

// kickoff ends the static IP intro: every mosher gets a gentle sideways shove,
// alternating left/right so neighbors ram into each other. The ramp then
// builds that first calm jostle up into full chaos over rampDuration frames.
func (s *moshSim) kickoff() {
	s.rampFrames = rampDuration
	for i := range s.moshers {
		dir := 1.0
		if i%2 == 1 {
			dir = -1.0
		}
		s.moshers[i].vx = dir * 0.04
		s.moshers[i].vy = (s.rng.Float64()*2 - 1) * 0.03
	}
}

// step advances the simulation by one frame.
func (s *moshSim) step() {
	if s.introFrames > 0 {
		s.introFrames--
		if s.introFrames == 0 {
			s.kickoff() // sideways shove that breaks the IP into the mosh
		}
		return
	}

	const (
		swirl   = 0.011 // tangential push → circular motion
		jitter  = 0.020 // random chaos per frame
		damp    = 0.92  // velocity decay
		maxV    = 0.090 // speed cap
		collide = 0.40  // contact distance between moshers
		rebound = 0.014 // extra kick apart when moshers collide
		bandR   = 0.58  // radius the crowd is drawn toward

		fallChance    = 0.010 // per-frame odds someone goes down, when nobody is
		fallMinFrames = 20    // shortest time on the ground (~1s)
		fallMaxFrames = 45    // longest time on the ground (~2.5s)
	)

	// Intensity ramps from 0 (calm) to 1 (full chaos) over rampDuration frames
	// after the intro, so the pit escalates instead of exploding at once.
	intensity := 1.0
	if s.rampFrames > 0 {
		intensity = 1.0 - float64(s.rampFrames)/float64(rampDuration)
		s.rampFrames--
	}

	// Occasionally send one upright mosher to the ground — but only ever one
	// at a time, and only once the pit is at full chaos.
	anyDown := false
	for i := range s.moshers {
		if s.moshers[i].fallTimer > 0 {
			anyDown = true
			break
		}
	}
	if s.rampFrames == 0 && !anyDown && s.rng.Float64() < fallChance {
		idx := s.rng.Intn(len(s.moshers))
		s.moshers[idx].fallTimer = fallMinFrames + s.rng.Intn(fallMaxFrames-fallMinFrames)
		s.moshers[idx].vx, s.moshers[idx].vy = 0, 0
	}

	// Apply swirl, an inward pull toward the mosh band, and random chaos.
	for i := range s.moshers {
		m := &s.moshers[i]
		if m.fallTimer > 0 {
			continue // a downed mosher takes no forces
		}
		if r := math.Hypot(m.x, m.y); r > 1e-4 {
			m.vx += -m.y / r * swirl * intensity
			m.vy += m.x / r * swirl * intensity
			pull := (bandR - r) * 0.013 * intensity
			m.vx += m.x / r * pull
			m.vy += m.y / r * pull
		}
		m.vx += (s.rng.Float64()*2 - 1) * jitter * intensity
		m.vy += (s.rng.Float64()*2 - 1) * jitter * intensity
	}

	// Integrate, damp, and cap speed.
	for i := range s.moshers {
		m := &s.moshers[i]
		if m.fallTimer > 0 {
			// Stumble to the middle of the pit and stay down.
			m.x *= 0.82
			m.y *= 0.82
			m.vx, m.vy = 0, 0
			m.fallTimer--
			if m.fallTimer == 0 {
				// Helped back up — rejoin the swirl with an outward nudge.
				ang := s.rng.Float64() * 2 * math.Pi
				m.vx = math.Cos(ang) * 0.05
				m.vy = math.Sin(ang) * 0.05
			}
			continue
		}
		m.vx *= damp
		m.vy *= damp
		if sp := math.Hypot(m.vx, m.vy); sp > maxV {
			m.vx = m.vx / sp * maxV
			m.vy = m.vy / sp * maxV
		}
		m.x += m.vx
		m.y += m.vy
	}

	// Bounce off the pit edge.
	for i := range s.moshers {
		m := &s.moshers[i]
		r := math.Hypot(m.x, m.y)
		if r > 1 {
			nx, ny := m.x/r, m.y/r
			m.x, m.y = nx, ny
			if dot := m.vx*nx + m.vy*ny; dot > 0 {
				m.vx -= 2 * dot * nx
				m.vy -= 2 * dot * ny
			}
			m.vx *= 0.6
			m.vy *= 0.6
		}
	}

	// Moshers collide — push apart and shove velocity along the contact normal.
	for i := 0; i < len(s.moshers); i++ {
		for j := i + 1; j < len(s.moshers); j++ {
			a, b := &s.moshers[i], &s.moshers[j]
			if a.fallTimer > 0 || b.fallTimer > 0 {
				continue // downed moshers don't get shoved around
			}
			dx, dy := b.x-a.x, b.y-a.y
			d := math.Hypot(dx, dy)
			if d >= collide {
				continue
			}
			if d < 1e-4 {
				dx, dy, d = s.rng.Float64()-0.5, s.rng.Float64()-0.5, 0.1
			}
			ux, uy := dx/d, dy/d
			// Scale the collision response by intensity so the tightly packed
			// IP address eases apart instead of bursting on the first frame.
			push := (collide - d) / 2 * intensity
			a.x -= ux * push
			a.y -= uy * push
			b.x += ux * push
			b.y += uy * push
			av := a.vx*ux + a.vy*uy
			bv := b.vx*ux + b.vy*uy
			diff := (bv - av) * intensity
			a.vx += diff * ux
			a.vy += diff * uy
			b.vx -= diff * ux
			b.vy -= diff * uy
			// Extra rebound so the bump visibly kicks the pair apart.
			a.vx -= ux * rebound * intensity
			a.vy -= uy * rebound * intensity
			b.vx += ux * rebound * intensity
			b.vy += uy * rebound * intensity
		}
	}
}

// render rasterizes the pit to colored text lines for the splash body.
func (s *moshSim) render() []string {
	type cell struct {
		r     rune
		color string
	}
	grid := make([][]cell, moshPitH)
	for y := range grid {
		grid[y] = make([]cell, moshPitW)
		for x := range grid[y] {
			grid[y][x] = cell{r: ' '}
		}
	}

	cx, cy := float64(moshPitW-1)/2, float64(moshPitH-1)/2
	ax, ay := cx, cy

	// Faint ring marking the edge of the pit.
	ring := Hex(ActiveTheme.Surface2)
	for a := 0.0; a < 2*math.Pi; a += 0.18 {
		gx := int(math.Round(cx + ax*math.Cos(a)))
		gy := int(math.Round(cy + ay*math.Sin(a)))
		if gx >= 0 && gx < moshPitW && gy >= 0 && gy < moshPitH {
			grid[gy][gx] = cell{r: '·', color: ring}
		}
	}

	// The moshers themselves, drawn over the ring. A downed mosher shows as
	// a dim '_' in the middle until the crowd helps them back up.
	for i := range s.moshers {
		m := &s.moshers[i]
		gx := int(math.Round(cx + m.x*ax))
		gy := int(math.Round(cy + m.y*ay))
		if gx < 0 || gx >= moshPitW || gy < 0 || gy >= moshPitH {
			continue
		}
		glyph, color := m.glyph, moshColor(i)
		if m.fallTimer > 0 {
			glyph, color = '_', Hex(ActiveTheme.Overlay0)
		}
		grid[gy][gx] = cell{r: glyph, color: color}
	}

	lines := make([]string, moshPitH)
	for y := range grid {
		var b strings.Builder
		for _, c := range grid[y] {
			if c.r == ' ' {
				b.WriteByte(' ')
				continue
			}
			b.WriteString("[" + c.color + "]")
			b.WriteRune(c.r)
			b.WriteString("[-]")
		}
		lines[y] = b.String()
	}
	return lines
}

// moshColor returns a theme accent color for mosher i, so the crowd is as
// colorful as the active theme allows.
func moshColor(i int) string {
	palette := []tcell.Color{
		ActiveTheme.Text, ActiveTheme.Green, ActiveTheme.Teal,
		ActiveTheme.Peach, ActiveTheme.Mauve, ActiveTheme.Blue,
		ActiveTheme.Yellow, ActiveTheme.Sapphire, ActiveTheme.Lavender,
	}
	return Hex(palette[i%len(palette)])
}
