package podcastrss

var rsstemplate = `
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom">
<channel>
<title>{{.Show.Required.Title}}</title>
{{if .Show.Recommended.Website}}<link>{{.Show.Recommended.Website}}</link>{{ end }}
<language>{{.Show.Required.Language}}</language>
{{if .Show.Situational.Copyright}}<copyright>{{.Show.Situational.Copyright}}</copyright>{{end}}
{{if .Show.Recommended.Author}}<itunes:author>{{.Show.Recommended.Author}}</itunes:author> {{end}}
<itunes:summary>{{.Show.Required.Description}}</itunes:summary>
<description>{{.Show.Required.Description}}</description>
{{if .Show.Situational.Type}}<itunes:type>{{.Show.Situational.Type}}</itunes:type>{{end}}
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
{{if .Show.Situational.Block}}<itunes:block>{{.Show.Situational.Block}}</itunes:block>{{end}}
{{if .Show.Situational.NewURLFeed}}
<itunes:new-feed-url>
{{.Show.Situational.NewURLFeed}}
</itunes:new-feed-url>
{{end}}
{{range .Episodes}}
<item>
{{if .Situational.EpisodeType}}<itunes:episodeType>{{.Situational.EpisodeType}}</itunes:episodeType>{{end}}
{{if .Situational.Block}}<itunes:block>{{.Situational.Block}}</itunes:block>{{end}}
{{if .Situational.EpisodeNumber}}<itunes:episode>{{.Situational.EpisodeNumber}}</itunes:episode>{{end}}
{{if .Situational.Season}}<itunes:season>{{.Situational.Season}}</itunes:season>{{end}}
<title>{{.Required.Title}}</title>
{{if .Recommended.Description}}
<description>
{{.Recommended.Description}}
</description>
{{end}}
{{if .Recommended.Link}}<link>{{.Recommended.Link}}</link>{{end}}
{{if .Recommended.Image}}
<itunes:image href="{{.Recommended.Image}}" />
{{end}}
<enclosure 
  length="{{.Required.Enclosure.Length}}" 
  type="{{.Required.Enclosure.Type}}" 
  url="{{.Required.Enclosure.URL}}"
/>
{{if .Recommended.GUID}}<guid isPermaLink="false">{{.Recommended.GUID}}</guid>{{end}}
{{if .Recommended.PublishDate}}<pubDate>{{.Recommended.PublishDate}}</pubDate>{{end}}
{{if .Recommended.Duration}}<itunes:duration>{{.Recommended.Duration}}</itunes:duration>{{end}}
{{if .Recommended.Explicit}}<itunes:explicit>{{.Recommended.Explicit}}</itunes:explicit>{{end}}
</item>
{{end}}
</channel>
</rss>
`
