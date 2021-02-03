package main

var rsstemplate = `
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom">
<channel>
<title>{{.Show.Required.Title}}</title>
{{if .Show.Recommended.Website}}<link>{{.Show.Recommended.Website}}</link>{{ end }}
<language>{{.Show.Required.Language}}</language>
{{if .Show.Recommended.Author}}<itunes:author>{{.Show.Recommended.Author}}</itunes:author> {{end}}
<itunes:summary>{{.Show.Required.Description}}</itunes:summary>
<description>{{.Show.Required.Description}}</description>
{{if .Show.Recommended.Owner}} 
<itunes:owner>
{{if .Show.Recommended.Owner.Name}} <itunes:name>{{.Show.Recommended.Owner.Name}}</itunes:name>{{ end }}
{{if .Show.Recommended.Owner.Email}} <itunes:email>{{.Show.Recommended.Owner.Email}}</itunes:email>{{ end }}
</itunes:owner>
{{ end }}
<itunes:explicit>{{.Show.Required.Explicit}}</itunes:explicit>
<itunes:image href="{{.Show.Required.Image}}" />
{{range .Show.Required.Categories}}
<itunes:category text="{{.Category}}">
{{if .Subcategory}} <itunes:category text="{{.Subcategory}}" />{{end}}
</itunes:category>
{{end}}
{{if .Show.Recommended.HostURL}} <atom:link href="{{.Show.Recommended.HostURL}}" rel="self" type="application/rss+xml" />{{end}}
</channel>
</rss>
`
