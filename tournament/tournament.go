package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name    string
	matches int
	won     int
	tied    int
	lost    int
	points  int
}

type scoreBoard map[string]*team

func calcScore(stats scoreBoard, team1, team2, result string) error {
	var ok bool
	_, ok = stats[team1]
	if !ok {
		var t team
		t.name = team1
		stats[team1] = &t
	}

	_, ok = stats[team2]
	if !ok {
		var t team
		t.name = team2
		stats[team2] = &t
	}

	stats[team1].matches++
	stats[team2].matches++

	switch result {
	case "win":
		stats[team1].points += 3
		stats[team1].won++
		stats[team2].lost++
	case "draw":
		stats[team1].points++
		stats[team2].points++
		stats[team1].tied++
		stats[team2].tied++
	case "loss":
		stats[team2].points += 3
		stats[team2].won++
		stats[team1].lost++
	default:
		return fmt.Errorf("Invalid game result %s", result)
	}
	return nil
}

func printStats(stats []*team, w *bufio.Writer) {
	fmtString := "%-31v| %2v | %2v | %2v | %2v | %2v\n"
	s := fmt.Sprintf(fmtString, "Team", "MP", "W", "D", "L", "P")
	w.WriteString(s)
	for _, t := range stats {
		s = fmt.Sprintf(fmtString, t.name, t.matches, t.won, t.tied, t.lost, t.points)
		w.WriteString(s)
	}
	w.Flush()
}

func sortStats(stats scoreBoard) []*team {
	var list []*team
	for _, t := range stats {
		list = append(list, t)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].points == list[j].points {
			return list[i].name < list[j].name
		}
		return list[i].points > list[j].points
	})
	return list
}

// Tally reads from rd and writes to wr
func Tally(rd io.Reader, wr io.Writer) error {
	stats := make(scoreBoard)
	scanner := bufio.NewScanner(rd)
	w := bufio.NewWriter(wr)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 || s[0] == '#' {
			continue
		}

		parts := strings.Split(s, ";")
		if len(parts) != 3 {
			return fmt.Errorf("Invalid format: %s", s)
		}

		err := calcScore(stats, parts[0], parts[1], parts[2])
		if err != nil {
			return err
		}
	}
	sortedStats := sortStats(stats)
	printStats(sortedStats, w)
	return nil
}
