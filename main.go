package main

import (
	"log"
	"os"
	"time"

	"github.com/curi0s/twirgo"
	"github.com/micmonay/keybd_event"
)

// acc		=
// back		=
// left		= left + acc
// right	= right + acc
// boost
// flip
// ---
// halfflip
// was = what a save x3

type car struct {
	kb *keybd_event.KeyBonding
}

func initCar() *car {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	return &car{
		kb: &kb,
	}
}

func (c *car) event(sleep int, keys ...int) {
	for _, key := range keys {
		c.kb.AddKey(key)
	}

	c.kb.Press()
	time.Sleep(500 * time.Millisecond)
	c.kb.Release()

	c.kb.Clear()
}

func (c *car) acc() {
	c.event(500, keybd_event.VK_W)
}

func (c *car) back() {
	c.event(500, keybd_event.VK_S)
}

func (c *car) bleft() {
	c.event(500, keybd_event.VK_S, keybd_event.VK_A)
}

func (c *car) bright() {
	c.event(500, keybd_event.VK_S, keybd_event.VK_D)
}

func (c *car) fleft() {
	c.event(500, keybd_event.VK_W, keybd_event.VK_A)
}

func (c *car) fright() {
	c.event(500, keybd_event.VK_W, keybd_event.VK_D)
}

func (c *car) boost() {
	c.event(500, keybd_event.VK_B)
}

func (c *car) halfflip() {
	c.flip(keybd_event.VK_S)

	time.Sleep(200 * time.Millisecond)

	c.kb.SetKeys(keybd_event.VK_W)
	c.kb.Press()
	time.Sleep(200 * time.Millisecond)

	c.kb.HasSHIFT(true)
	c.kb.AddKey(keybd_event.VK_A)
	c.kb.Press()
	time.Sleep(500 * time.Millisecond)
	c.kb.Release()

	c.kb.Clear()

	time.Sleep(300 * time.Millisecond)
}

func (c *car) flip(key int) {
	c.kb.SetKeys(keybd_event.VK_J)
	c.kb.Press()
	time.Sleep(10 * time.Millisecond)
	c.kb.Release()

	time.Sleep(200 * time.Millisecond)

	c.kb.SetKeys(key, keybd_event.VK_J)
	c.kb.Press()
	time.Sleep(10 * time.Millisecond)
	c.kb.Release()

	c.kb.Clear()
}

func (c *car) frontflip() {
	c.flip(keybd_event.VK_W)
}

func (c *car) rightflip() {
	c.flip(keybd_event.VK_D)
}

func (c *car) leftflip() {
	c.flip(keybd_event.VK_A)
}

func (c *car) chatWhatASave() {
	for i := 1; i < 4; i++ {
		c.kb.SetKeys(keybd_event.VK_2)
		c.kb.Launching()

		time.Sleep(10 * time.Millisecond)

		c.kb.SetKeys(keybd_event.VK_4)
		c.kb.Launching()

		if i < 3 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	c.kb.Clear()
}

func handleEvents(t *twirgo.Twitch, ch chan interface{}) error {
	car := initCar()
	for event := range ch {
		switch ev := event.(type) {
		case twirgo.EventMessageReceived:
			switch ev.Message.Content {
			case "acc":
				car.acc()
			case "back":
				car.back()
			case "left":
				fallthrough
			case "fleft":
				car.fleft()
			case "right":
				fallthrough
			case "fright":
				car.fright()
			case "bleft":
				car.bleft()
			case "bright":
				car.bright()
			case "boost":
				car.boost()
			case "frontflip":
				fallthrough
			case "fflip":
				fallthrough
			case "jump":
				fallthrough
			case "flip":
				car.frontflip()
			case "lflip":
				car.leftflip()
			case "rflip":
				car.rightflip()
			case "hflip":
				fallthrough
			case "halfflip":
				car.halfflip()
			case "was":
				car.chatWhatASave()
			}

		case twirgo.EventConnectionError:
			return ev.Err
		}
	}

	return nil
}

func main() {
	// kb, err := keybd_event.NewKeyBonding()
	// if err != nil {
	// 	panic(err)
	// }

	// flip
	// kb.SetKeys(keybd_event.VK_J)
	// kb.Press()
	// time.Sleep(10 * time.Millisecond)
	// kb.Release()

	// time.Sleep(100 * time.Millisecond)

	// kb.SetKeys(keybd_event.VK_W, keybd_event.VK_J)
	// kb.Press()
	// time.Sleep(10 * time.Millisecond)
	// kb.Release()

	options := twirgo.Options{
		Username:       "curi0sde_bot",
		Token:          os.Getenv("TOKEN"),
		Channels:       []string{"curi0sde"},
		DefaultChannel: "curi0sde",
	}

	t := twirgo.NewTwirgo(options)

	ch, err := t.Connect()
	if err == twirgo.ErrInvalidToken {
		log.Fatal(err)
	}

	err = handleEvents(t, ch)
	if err != nil {
		log.Fatal(err)
	}
}
