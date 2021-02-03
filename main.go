package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ProgramConfig struct for all the program configs.
type ProgramConfig struct {
	SearchFolder    string `yaml:"SearchFolder"`
	Podcastfilename string `yaml:"Podcastfilename"`
}

// Podcast single podcast detail
type Podcast struct {
	Show     ShowDetail `yaml:"Show"`
	Web      WebDetail  `yaml:"Website"`
	Episodes []Episode  `yaml:"Episodes"`
}

// RequiredEpisodeDetails Required details about a show
type RequiredEpisodeDetails struct {
	Title     string                   `yaml:"Title"`
	Enclosure RequiredEpisodeEnclosure `yaml:"Enclosure"`
}

// RequiredEpisodeEnclosure Required details about a show length URL and format
type RequiredEpisodeEnclosure struct {
	URL    string `yaml:"URL"`
	Length string `yaml:"Length"`
	Type   string `yaml:"Type"`
}

// RecommendedEpisodeDetails Recommended details about the episode
type RecommendedEpisodeDetails struct {
	GUID        string `yaml:"GUID"`
	PublishDate string `yaml:"PublishDate"`
	Duration    string `yaml:"Duration"`
	Link        string `yaml:"Link"`
	Image       string `yaml:"Image"`
	Explicit    string `yaml:"Explicit"`
}

// SituationalEpisodeDetails Situational details about the episode
type SituationalEpisodeDetails struct {
	EpisodeNumber string `yaml:"EpisodeNumber"`
	Season        string `yaml:"Season"`
	EpisodeType   string `yaml:"EpisodeType"`
	Block         string `yaml:"Block"`
}

// Episode Holds the Detail about an Episode
type Episode struct {
	Required    RequiredEpisodeDetails    `yaml:"Required"`
	Recommended RecommendedEpisodeDetails `yaml:"Recommended"`
}

// WebDetail the details about how the url RSS will be hosted
type WebDetail struct {
	URL string `yaml:"URL"`
}

// RequiredShowDetails information required by itunes RSS feed
type RequiredShowDetails struct {
	Title       string     `yaml:"Title"`
	Description string     `yaml:"Description"`
	Language    string     `yaml:"Language"`
	Image       string     `yaml:"Image"`
	Categories  []Category `yaml:"Categories"`
	Explicit    string     `yaml:"Explicit"`
}

// Category Categories defined by Apple
type Category struct {
	Category    string `yaml:"Category"`
	Subcategory string `yaml:"Subcategory"`
}

// SituationalShowDetails the situational details about your
type SituationalShowDetails struct {
	Type       string `yaml:"Type"`
	Copyright  string `yaml:"Copyright"`
	NewURLFeed string `yaml:"NewURLFeed"`
	Block      string `yaml:"Block"`
	Complete   string `yaml:"Complete"`
}

// ShowDetail is the details about you podcast
type ShowDetail struct {
	Required    RequiredShowDetails    `yaml:"Required"`
	Recommended RecommendedShowDetails `yaml:"Recommended"`
	Situational SituationalShowDetails `yaml:"Situational"`
}

// RecommendedShowDetails information recommended by itunes RSS feed
type RecommendedShowDetails struct {
	Website string       `yaml:"Website"`
	Author  string       `yaml:"Author"`
	Owner   OwnerDetails `yaml:"Owner"`
	HostURL string       `yaml:"HostURL"`
}

// OwnerDetails information about the owner of the podcast
type OwnerDetails struct {
	Name  string `yaml:"Name"`
	Email string `yaml:"Email"`
}

func main() {
	var Config string
	flag.StringVar(&Config, "f", "./config/config.yaml", "YAML file to parse.")
	flag.Parse()
	ProgramConfiguration, err := initializeConfig(Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Looking for Folders under the provided Search Folder")
	PodCasts := ProgramConfiguration.createPodcastsStruct()
	log.Printf("Found %d Podcasts", len(PodCasts))
	for _, podcast := range PodCasts {
		CreateFeeds(podcast)
	}

}

func (pc *ProgramConfig) createPodcastsStruct() []Podcast {
	var Podcasts []Podcast
	files, err := ioutil.ReadDir(pc.SearchFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		filename := filepath.Join(pc.SearchFolder, f.Name())
		//log.Println("Found: " + filename)
		show, _ := getShowDetails(filepath.Join(filename, pc.Podcastfilename))
		log.Printf("%+v\n", show)
		Podcasts = append(Podcasts, show)
	}
	return Podcasts

}

func getShowDetails(filename string) (Podcast, error) {
	var PC Podcast
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
		return PC, err
	}
	err = yaml.Unmarshal(yamlFile, &PC)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
		return PC, err
	}

	return PC, nil
}

func initializeConfig(file string) (ProgramConfig, error) {
	var programconfiguration ProgramConfig
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
		return programconfiguration, err
	}
	err = yaml.Unmarshal(yamlFile, &programconfiguration)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
		return programconfiguration, err
	}
	return programconfiguration, err
}
