//go:generate fyne bundle -o bundled.go pieces
package main

import "fyne.io/fyne/v2"

func GetPieceFromRessource() fyne.Resource {
	return resourceBlackBishopSvg
}
