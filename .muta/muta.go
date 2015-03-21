package main

import (
	"./plugins/prettyurls"
	"github.com/leeola/muta"
	"github.com/leeola/muta-frontmatter"
	"github.com/leeola/muta-markdown"
	"github.com/leeola/muta-sass"
	"github.com/leeola/muta-template"
)

type Guide struct {
	Title  string   `yaml:"title"`
	Author string   `yaml:"author"`
	Date   string   `yaml:"date"`
	Tags   []string `yaml:"categories"`
}

type Page struct {
	Title string `yaml:"title"`
}

func fmTyper(st string) (t interface{}, err error) {
	switch st {
	case "page":
		t = &Page{}
	}
	return
}

func Markdown() (*muta.Stream, error) {
	//s := muta.Src("./**/*.md").
	// Manually loading a single template for now
	s := muta.Src("./guides/getting-started-kpm.md").
		Pipe(frontmatter.FrontMatter(fmTyper)).
		Pipe(markdown.Markdown()).
		Pipe(template.Template("./.muta/templates")).
		Pipe(prettyurls.Prettyurls()).
		Pipe(muta.Dest("./build"))
	return s, nil
}

func Sass() (*muta.Stream, error) {
	s := muta.Src("./.muta/sass/*.scss").
		Pipe(sass.Sass()).
		Pipe(muta.Dest("./build/css"))
	return s, nil
}

func main() {
	muta.Task("markdown", Markdown)
	muta.Task("sass", Sass)

	muta.Task("build", "markdown", "sass")
	muta.Task("default", "markdown")
	muta.Te()
}
