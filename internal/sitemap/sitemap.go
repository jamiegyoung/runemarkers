package sitemap

import (
	"fmt"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/lastmod"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	"github.com/jamiegyoung/runemarkers-go/internal/pageio"
	"github.com/jamiegyoung/runemarkers-go/internal/pages"
)

var log = logger.New("sitemap")

func Generate(ents []*entities.Entity) error {
	err := lastmod.GenerateDb()
	if err != nil {
		return err
	}

	modded, err := lastmod.EntitiesDiff(ents)
	if err != nil {
		return err
	}

	log(fmt.Sprintf("%v modded files found", len(modded)))

	err = lastmod.UpdateEntities(modded)
	if err != nil {
		return err
	}

	entMods, err := lastmod.GetEntites()
	if err != nil {
		return err
	}

	log(fmt.Sprintf("%v entities for sitemap found", len(entMods)))
	return processTemplate(entMods)
}

func processTemplate(entMods []*lastmod.EntityMod) error {
	sitemap_page, err := pageio.ReadPageString("templates/sitemap/sitemap.tmpl")
	if err != nil {
		return err
	}

	output, err := pageio.CreateOutFile(
		"public/",
		"sitemap.xml",
	)
	if err != nil {
		return err
	}
	defer output.Close()

	data := pages.NewPage(
		map[string]interface{}{
			"EntityMods": entMods,
		},
	)

	err = pageio.RenderPage(
		"sitemap",
		sitemap_page,
		output,
		data,
	)
	if err != nil {
		return err
	}

	return nil
}
