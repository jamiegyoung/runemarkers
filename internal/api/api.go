package api

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
	"github.com/jamiegyoung/runemarkers-go/internal/pages"
)

var log = logger.New("api")

func Generate(ents []*entities.Entity) error {
	api_templates := []string{
		"templates/api/button.tmpl",
		"templates/api/tile_data_display.tmpl",
		"templates/api/tile_data_display_pretty.tmpl",
		"templates/api/tile_data_display_truncated.tmpl",
		"templates/api/tile_data_display_pretty_truncated.tmpl",
	}

	var err error = nil
	for _, tmpl := range api_templates {
		log("Generating entity specific api for " + tmpl)
		err = GenerateEntitySpecific(ents, tmpl)
		if err != nil {
			log("Error generating entity specific api: " + err.Error())
			break
		}
	}
	return err
}

func GenerateEntitySpecific(ents []*entities.Entity, tmpl_path string) error {
	api_page, err := pageio.ReadPageString(tmpl_path)
	if err != nil {
		return err
	}

	path_with_ext := filepath.Base(tmpl_path)
	api_name := strings.TrimSuffix(path_with_ext, filepath.Ext(path_with_ext))
	log("Generating " + api_name + " api")

	// create the api folder if it doesn't exist
	if _, err := os.Stat("public/api"); os.IsNotExist(err) {
		err = os.Mkdir("public/api", 0755)
		if err != nil {
			return err
		}
	}

	for _, entity := range ents {
		output, err := pageio.CreateOutFile(
			"public/api",
			"api/"+api_name+"_"+entity.ApiUri+".html",
		)
		if err != nil {
			return err
		}
		defer output.Close()

		data := pages.NewPage(
			map[string]interface{}{
				"TilesString":       entity.TilesString,
				"TilesStringPretty": entity.TilesStringPretty,
				"SafeApiUri":        entity.SafeApiUri,
			},
		)

		err = pageio.RenderHtml(
			api_name+"_"+entity.ApiUri,
			api_page,
			output,
			data,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
