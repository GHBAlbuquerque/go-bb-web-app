package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(writer http.ResponseWriter, name string) {
	// create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal("Error parsing page:", err)
	}

	// get requested template from cache
	template, ok := templateCache[name]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(writer)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// initialize template map
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl
	// Glob returns the names of all files matching pattern or nil if there is no matching file
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all

	for _, page := range pages {
		// Base returns the last element of path (file name)
		name := filepath.Base(page)

		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get the templates
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// associates the templates with the template I got
		if len(matches) > 0 {
			// ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with t.
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		// adds the template to the map
		myCache[name] = templateSet
	}

	return myCache, nil
}
