package lastmod

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/jamiegyoung/runemarkers/internal/entities"
	"github.com/jamiegyoung/runemarkers/internal/logger"
)

type EntityMod struct {
	Entity        entities.Entity
	ModTime       time.Time
	ModTimeString string
}

var log = logger.New("lastmod")

func BuildEntityMods(entities []*entities.Entity) ([]*EntityMod, error) {
	modTimes, err := gitLastModified()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	mods := make([]*EntityMod, len(entities))
	for i, ent := range entities {
		modTime, ok := modTimes[ent.SourcePath]
		if !ok {
			log(fmt.Sprintf(
				"no git history for %v, using current time",
				ent.SourcePath,
			))
			modTime = now
		}
		mods[i] = &EntityMod{
			Entity:        *ent,
			ModTime:       modTime,
			ModTimeString: Format(modTime),
		}
	}

	return mods, nil
}

// Output is NUL-separated per commit, e.g.:
//
//	\02026-07-23T10:12:49+01:00
//	entities/doom (sidewalk).json
//
//	\02024-07-13T19:39:54+01:00
//	entities/vorkath.json
//	entities/zulrah.json
func gitLastModified() (map[string]time.Time, error) {
	cmd := exec.Command(
		"git", "log", "--name-only", "--no-merges",
		"--format=format:%x00%aI",
		"--", "entities/*.json",
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git log failed: %w", err)
	}

	return parseGitLog(out), nil
}

func parseGitLog(out []byte) map[string]time.Time {
	modTimes := make(map[string]time.Time)

	for _, block := range bytes.Split(out, []byte{0}) {
		lines := strings.Split(strings.TrimSpace(string(block)), "\n")
		if len(lines) == 0 || lines[0] == "" {
			continue
		}

		commitTime, err := time.Parse(time.RFC3339, lines[0])
		if err != nil {
			continue
		}

		for _, path := range lines[1:] {
			path = strings.TrimSpace(path)
			if path == "" {
				continue
			}
			if _, exists := modTimes[path]; !exists {
				modTimes[path] = commitTime
			}
		}
	}

	return modTimes
}

func FindLastMod(mods []*EntityMod) time.Time {
	if len(mods) == 0 {
		return time.Now().UTC()
	}

	lastmod := mods[0].ModTime
	for _, mod := range mods {
		if mod.ModTime.After(lastmod) {
			lastmod = mod.ModTime
		}
	}

	return lastmod
}

func Format(t time.Time) string {
	return t.Format(time.RFC3339)
}
