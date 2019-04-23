package web

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	templates map[string]*template.Template
	bufpool   *bpool.BufferPool
	loaded    = false
)

func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	loadTemplates()
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	buf := bufpool.Get()
	defer bufpool.Put(buf)
	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return nil
	}
	w.Header().Set("X-Frame-Options", "deny")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}
func loadTemplates() {
	if loaded {
		return
	}
	templates = make(map[string]*template.Template)
	bufpool = bpool.NewBufferPool(64)
	layoutTemplates := map[string][]string{
		"web/layouts/outside.html": {
			"./web/includes/register.html",
			"./web/includes/login.html",
		},
		"web/layouts/inside.html": {
			"./web/includes/authorize.html",
		},
	}
	for layout, includes := range layoutTemplates {
		for _, include := range includes {
			files := []string{include, layout}
			templates[filepath.Base(include)] = template.Must(template.ParseFiles(files...))
		}
	}
	loaded = true
}
