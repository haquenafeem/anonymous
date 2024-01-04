package internal

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/fogleman/gg"
)

type ShareMessage struct {
	ctx           *gg.Context
	width, height int
}

func NewShareMessage() *ShareMessage {
	return &ShareMessage{
		ctx:    gg.NewContext(1000, 1000),
		width:  1000,
		height: 1000,
	}
}

func (sm *ShareMessage) clear() {
	sm.ctx.SetRGBA(0, 0, 0, 0)
	sm.ctx.Clear()
}

func (sm *ShareMessage) GetImageBytes(message string) ([]byte, error) {
	sm.clear()
	sm.setGradient()
	sm.drawRecangle()
	sm.drawRoundedRectangle()
	err := sm.loadFont()
	if err != nil {
		return nil, err
	}

	sm.writeMessage(message)
	return sm.getPngBytes()
}

func (sm *ShareMessage) setGradient() {
	gradient := gg.NewLinearGradient(0, 0, float64(sm.width), float64(sm.height))
	gradient.AddColorStop(0, color.RGBA{255, 0, 151, 255})
	gradient.AddColorStop(1, color.RGBA{64, 64, 128, 255})

	sm.ctx.SetFillStyle(gradient)
}

func (sm *ShareMessage) drawRecangle() {
	sm.ctx.DrawRoundedRectangle(0, 0, float64(sm.width), float64(sm.height), 50)
	sm.ctx.Fill()
}

func (sm *ShareMessage) drawRoundedRectangle() {
	sm.ctx.DrawRoundedRectangle(100, 100, 800, 800, 50)
	sm.ctx.SetRGB(1, 1, 1)
	sm.ctx.Fill()
}

func (sm *ShareMessage) loadFont() error {
	return sm.ctx.LoadFontFace("assets/font.ttf", 40)
}

func (sm *ShareMessage) writeMessage(text string) {
	sm.ctx.SetRGB(255, 255, 255)
	sm.setText(text)
}

func (sm *ShareMessage) getPngBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, sm.ctx.Image())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (sm *ShareMessage) splitString(input string, chunkSize int) []string {
	var result []string

	runes := []rune(input)

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		result = append(result, string(runes[i:end]))
	}

	return result
}

func (sm *ShareMessage) setText(text string) {
	strs := sm.splitString(text, 35)

	startHeight := 200
	for _, str := range strs {
		sm.ctx.DrawString(str, 170, float64(startHeight))
		startHeight += 40
	}
}
