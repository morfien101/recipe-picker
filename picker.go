package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
)

func pick(from string, numberOfRecipes int) ([]string, error) {
	recipes, err := collectFilesNames(from)
	if err != nil {
		return nil, err
	}
	return selectFiles(numberOfRecipes, recipes), nil
}

func collectFilesNames(dir string) ([]string, error) {
	pattern := fmt.Sprintf("%s/*\\.pdf", dir)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) < 1 {
		return nil, fmt.Errorf("Failed to find recipes in given directory")
	}
	return matches, nil
}

func selectFiles(howMany int, from []string) []string {
	seed()
	positions := []int{}
	l := len(from)

	alreadyPicked := func(incoming int) bool {
		if len(positions) == 0 {
			return false
		}
		for _, number := range positions {
			if incoming == number {
				return true
			}
		}
		return false
	}

	for i := 0; i < howMany; i++ {
		randInt := rand.Intn(l)
		if alreadyPicked(randInt) {
			i--
			continue
		}
		positions = append(positions, randInt)
	}

	selection := []string{}
	for _, p := range positions {
		selection = append(selection, from[p])
	}
	return baseStripper(selection)
}

func baseStripper(files []string) []string {
	output := []string{}
	for _, file := range files {
		f := strings.Split(file, "/")
		output = append(output, f[len(f)-1])
	}
	return output
}
