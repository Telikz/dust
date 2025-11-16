package cmd

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/telikz/dust/internal/model"
	"github.com/telikz/dust/internal/systems"
)

type Model struct {
	world        *model.World
	systems      []model.System
	ticker       *time.Ticker
	currentType  string
	gridWidth    int
	gridHeight   int
	screenWidth  int
	screenHeight int
}

type particleType string

const (
	Sand  particleType = "Sand"
	Water particleType = "Water"
	Oil   particleType = "Oil"
)

func NewModel() *Model {
	world := model.NewWorld()

	// Initialize systems
	sysList := []model.System{
		&systems.GravitySystem{Gravity: 9.8},
		&systems.PhysicsSystem{},
		systems.NewCollisionSystem(2.0),
		&systems.FlowSystem{},
	}

	return &Model{
		world:        world,
		systems:      sysList,
		ticker:       time.NewTicker(16 * time.Millisecond), // ~60 FPS
		currentType:  "Sand",
		gridWidth:    80,
		gridHeight:   24,
		screenWidth:  80,
		screenHeight: 24,
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

type tickMsg struct{}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height
		m.gridWidth = msg.Width
		m.gridHeight = msg.Height - 4 // Reserve space for header
		m.world.Map.Width = msg.Width
		m.world.Map.Height = m.gridHeight
		m.world.FloorY = float64(m.gridHeight - 1)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.currentType = "Sand"
		case "2":
			m.currentType = "Water"
		case "3":
			m.currentType = "Oil"
		}
	case tea.MouseMsg:
		// Spawn particle at mouse position
		x := float64(msg.X)
		y := float64(msg.Y - 3) // Offset for header

		if y >= 0 && y < float64(m.gridHeight) {
			switch m.currentType {
			case "Sand":
				model.CreateSand(m.world, x, y)
			case "Water":
				model.CreateWater(m.world, x, y)
			case "Oil":
				model.CreateOil(m.world, x, y)
			}
		}
	case tickMsg:
		m.world.Update(m.systems, 0.016) // 16ms delta time
		return m, tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
			return tickMsg{}
		})
	}
	return m, nil
}

func (m *Model) View() string {
	if m.gridHeight <= 0 || m.gridWidth <= 0 {
		return "Loading..."
	}

	// Create grid to track particles and colors
	type Cell struct {
		hasParticle bool
		r, g, b     uint8
	}
	grid := make([][]Cell, m.gridHeight)
	for i := range grid {
		grid[i] = make([]Cell, m.gridWidth)
	}

	// Draw particles with their colors
	for p := range m.world.Particles {
		comps := m.world.Components[p]
		if pos, ok := comps["Position"].(*model.Position); ok {
			x := int(pos.X)
			y := int(pos.Y)

			if x >= 0 && x < m.gridWidth && y >= 0 && y < m.gridHeight {
				cell := &grid[y][x]
				cell.hasParticle = true

				// Get color from Color component
				if color, ok := comps["Color"].(*model.Color); ok {
					cell.r = uint8(color.R * 255)
					cell.g = uint8(color.G * 255)
					cell.b = uint8(color.B * 255)
				}
			}
		}
	}

	// Build output with ANSI colors
	output := "Dust Particles - Draw Mode\n"
	output += "Press: 1=Sand, 2=Water, 3=Oil | Click to draw | Q to quit\n"
	output += fmt.Sprintf("Current: %s | Particles: %d\n\n", m.currentType, len(m.world.Particles))

	floorRow := int(m.world.FloorY)
	for y, row := range grid {
		for _, cell := range row {
			if cell.hasParticle {
				// ANSI 24-bit true color: \033[38;2;R;G;Bm
				output += fmt.Sprintf("\033[38;2;%d;%d;%dm●\033[0m", cell.r, cell.g, cell.b)
			} else if y == floorRow {
				// Draw floor
				output += "\033[38;2;100;100;100m─\033[0m"
			} else {
				output += " "
			}
		}
		output += "\n"
	}

	return output
}

func RunTUI() error {
	p := tea.NewProgram(NewModel(), tea.WithMouseCellMotion(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
