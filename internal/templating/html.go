package templating

import "html/template"

func HtmlTemplateWithComponents(name string, text string) (*template.Template, error) {
	log("Generating html template with components for " + name)
	tmpl, err := template.New(name).Funcs(funcMap).Parse(text)

	if err != nil {
		return nil, err
	}

	style, err := readComponentStyles()
	if err != nil {
		log("Error reading styles")
		return nil, err
	}

	tmpl, err = tmpl.Funcs(funcMap).Parse(style)
	if err != nil {
		log("Error parsing style string: \n" + style)
		return nil, err
	}

	components, err := readComponents()
	if err != nil {
		log("Error reading components")
		return nil, err
	}

	// Load all components and return the template
	for _, component := range components {
		tmpl, err = tmpl.Funcs(funcMap).Parse(component)
		if err != nil {
			log("Error parsing component string: \n" + component)
			return nil, err
		}
	}
	return tmpl, nil
}
