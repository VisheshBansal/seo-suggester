package main

type CheckMetaTags struct {
	Charset			   string `json:"charset"`
	ViewPort           string `json:"viewport"`
	Description        string `json:"description"`
	KeyWords           string `json:"keywords"`
	Author             string `json:"author"`
	Copyright          string `json:"copyright"`
	Language           string `json:"language"`
	MetaURL            string `json:"meta_url"`
	Category           string `json:"category"`
	Coverage           string `json:"coverage"`
	Rating             string `json:"rating"`
	OgEmail            string `json:"og_email"`
	OgCountryName      string `json:"og_country_name"`
	OgRegion           string `json:"og_region"`
	OgSiteName         string `json:"og_site_name"`
	OgType             string `json:"og_type"`
	OgTitle            string `json:"og_title"`
	OgDescription      string `json:"og_description"`
	OgURL              string `json:"og_url"`
	TwitterCard        string `json:"twitter_card"`
	TwitterWebsite     string `json:"twitter_website"`
	TwitterTitle       string `json:"twitter_title"`
	TwitterDescription string `json:"twitter_description"`
	TwitterCreator     string `json:"twitter_creator"`
	TwitterImage       string `json:"twitter_image"`
}

// Contains Main, Open Graph and Twitter Meta Tags