package podcastrss

import (
	"os"
	"text/template"
)

// CreateFeeds Generates the XML Template for the RSS Feed
func CreateFeeds(cast Podcast) {
	tmpl, err := template.New(cast.Web.URL).Parse(rsstemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, cast)
	if err != nil {
		panic(err)
	}
}
