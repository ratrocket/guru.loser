package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Release struct{
	Id          int    `json:"database_id"`
	ArtistName  string `json:"artist_name"`
	ArtistSlug  string `json:"artist_slug"`
	ReleaseName string `json:"release_name"`
	ReleaseSlug string `json:"release_slug"`
	CatNum      string `json:"catalog_number"`
	Desc        string `json:"description"`
	BareNum     string `json:"bare_sp_num"`
	SPNum       string `json:"cat_with_sprefix"`
	DateNum     string `json:"release_date_numeric"`
	DateWords   string `json:"release_date_string"`
	ArtistLink  string `json:"artist_link"`
	SiteLink    string `json:"site_link"`
	ShopLink    string `json:"shop_link"`
	ImgSm       string `json:"img_sm"`
	ImgMd       string `json:"img_md"`
	ImgLg       string `json:"img_lg"`
}

type Releases []Release

type ViewData struct{
	Rs Releases
	ProcessingInterior bool
	CurrentPopupAnchor  string
	AssetPrefix string
	ImagePrefix string
}

// cloudfront path for guru.loser:
// http://d1nubnl3w13qxh.cloudfront.net
func main() {
	var jsonPath = flag.String("json", "json.large", "json file")
	var tmplPath = flag.String("tmpl", "main.tmpl", "template file")
	// Assume images are in an "img/" dir; optionally prepend something.
	var imgPrefix = flag.String("img", "", "path/url to \"img/\"")
	// TODO option for extension/no ext on "interior" files.
	//  I think Pages will want the extensions, the Go server won't work
	//  how I want it to *with* .html on the "interior" files.
	var asstPrefix = flag.String("asst", "", "asset prefix, prepended to paths like \"assets/javascripts/file.js\"")
	flag.Parse()

	vd := ViewData{AssetPrefix: *asstPrefix, ImagePrefix: *imgPrefix}

	jsonData, err := os.Open(*jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonData.Close()

	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &vd.Rs)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(vd.Rs); i++ {
		imageDir := "img/" // This structure is assumed.
		vd.Rs[i].ImgSm = fmt.Sprintf("%s%s%s", vd.ImagePrefix, imageDir, vd.Rs[i].ImgSm)
		vd.Rs[i].ImgMd = fmt.Sprintf("%s%s%s", vd.ImagePrefix, imageDir, vd.Rs[i].ImgMd)
		vd.Rs[i].ImgLg = fmt.Sprintf("%s%s%s", vd.ImagePrefix, imageDir, vd.Rs[i].ImgLg)
	}

	tmpl := template.Must(template.ParseFiles(*tmplPath))

	// change this so it writes a file
	if err := tmpl.Execute(os.Stdout, vd); err != nil {
		log.Fatalf("template execution: %s", err)
	}

	// Set for writing stuff on "interior" (non-index) pages
	vd.ProcessingInterior = true

	for i := 0; i < len(vd.Rs); i++ {
		vd.CurrentPopupAnchor = fmt.Sprintf("#a-popup-%s", vd.Rs[i].BareNum)

		//interiorFile := fmt.Sprintf("%s.html", vd.Rs[i].BareNum)
		// Try w/o extension
		interiorFile := fmt.Sprintf("%s", vd.Rs[i].BareNum)
		out, err := os.Create(interiorFile)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		if err := tmpl.Execute(out, vd); err != nil {
			log.Fatalf("template execution: %s", err)
		}

	}

}
