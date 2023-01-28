//go:generate fyne bundle -o data.go Icon.png

package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"
)

func main() {
	a := app.New()
	w := a.NewWindow("chess")
	game := chess.NewGame()
	grid := createGrid(game.Position().Board())

	go func() {
		rand.Seed(time.Now().Unix())
		for game.Outcome() == chess.NoOutcome {
			time.Sleep(time.Microsecond * 50000)
			valid := game.ValidMoves()
			m := valid[rand.Intn(len(valid))]
			move(m, game, grid)
		}
	}()

	w.SetContent(grid)
	w.Resize(fyne.NewSize(480, 480))
	w.ShowAndRun()
}

func createGrid(b *chess.Board) *fyne.Container {

	grid := container.NewGridWithColumns(8)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0xE0}
			}
			p := b.Piece(chess.Square(x + (7-y)*8))
			img := canvas.NewImageFromResource(ressourceForPiece(p))
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))
		}
	}

	return grid

}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container) {
	game.Move(m)
	refreshGrid(grid, game.Position().Board())
}

func refreshGrid(grid *fyne.Container, b *chess.Board) {
	y, x := 7, 0
	for _, cell := range grid.Objects {
		p := b.Piece(chess.Square(x + y*8))
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		img.Resource = ressourceForPiece(p)
		img.Refresh()

		x++
		if x == 8 {
			x = 0
			y--
		}
	}
}
