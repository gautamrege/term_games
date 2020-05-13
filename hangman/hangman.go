package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const MAX_CHANCES int = 8

var dev = flag.Bool("dev", false, "Development environment")

type Hangman struct{}

type Gamer interface {
	RenderGame([]string, map[string]bool) error
	GetEntry() string
}

func (h *Hangman) RenderGame(placeholder []string, entries map[string]bool) error {
	fmt.Println("\n")
	fmt.Println(placeholder)
	fmt.Printf("Chances left: %d\n", MAX_CHANCES-len(entries))
	fmt.Printf("Guesses: %s\n", get_keys(entries))
	fmt.Printf("Guess a letter or the word: ")

	return nil
}

func (h *Hangman) GetEntry() string {
	str := ""
	fmt.Scanln(&str)
	str = strings.TrimSuffix(str, "\n")

	return str
}

func get_keys(entries map[string]bool) []string {
	keys := []string{}
	for k, _ := range entries {
		keys = append(keys, k)
	}

	return keys
}

func play(h Gamer, word string) bool {
	entries := map[string]bool{}
	placeholder := make([]string, len(word), len(word))

	// initialize to '_'
	for i := range word {
		placeholder[i] = "_"
	}

	for {
		// evaluate a loss!
		if len(entries) == MAX_CHANCES {
			return false
		}

		// evaluate a win!
		if strings.Join(placeholder, "") == word {
			return true
		}

		h.RenderGame(placeholder, entries)
		str := h.GetEntry()
		if str == word {
			return true
		} else if str == "" {
			// someone pressed enter direclty!
			continue
		} else if len(str) > 1 {
			entries[str] = true
			continue
		}
		input := rune(str[0]) // take input and convert to rune (similar to char)

		// evalute the entry
		found := false
		for i, e := range word {
			if input == e {
				placeholder[i] = string(input)
				found = true
			}
		}
		if !found {
			entries[string(str)] = true
		}
	}

	// should never come here!
	return false
}

func get_word() string {
	res, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
	if err != nil {
		log.Fatal(err)
	}

	blob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var words []string
	err = json.Unmarshal(blob, &words)
	if err != nil {
		fmt.Println(err)
		return "elephant"
	}

	for _, word := range words {
		if len(word) > 4 && len(word) < 11 {
			// return a word between 5 and 10 characters
			return string(word)
		}
	}
	return "elephant" // default
}

func main() {
	flag.Parse()

	word := "elephant"
	if *dev == false {
		word = get_word()
	}

	game := &Hangman{}
	if play(game, word) == true {
		fmt.Println("You win! You've saved yourself from a hanging")
	} else {
		fmt.Println("Damn! You're hanged!!")
		fmt.Println("Word was: ", word)
	}
}
