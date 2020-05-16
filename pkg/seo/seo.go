package seo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/0x13a/golang.cafe/pkg/database"
)

func GeneratePostAJobSEOLandingPages(conn *sql.DB) ([]string, error) {
	var seoLandingPages []string
	locs, err := database.GetSEOLocations(conn)
	if err != nil {
		return seoLandingPages, err
	}
	for _, loc := range locs {
		seoLandingPages = appendPostAJobSEOLandingPageForLocation(seoLandingPages, loc.Name)
	}

	return seoLandingPages, nil
}

func GenerateSalarySEOLandingPages(conn *sql.DB) ([]string, error) {
	var landingPages []string
	locs, err := database.GetSEOLocations(conn)
	if err != nil {
		return landingPages, err
	}
	for _, loc := range locs {
		landingPages = appendSalarySEOLandingPageForLocation(landingPages, loc.Name)
	}

	return landingPages, nil
}

func appendSalarySEOLandingPageForLocation(landingPages []string, loc string) []string {
	tmpl := `Golang-Developer-Salary-%s`
	if strings.ToLower(loc) == "remote" {
		return append(landingPages, `Remote-Golang-Developer-Salary`)
	}
	return append(landingPages, fmt.Sprintf(tmpl, strings.ReplaceAll(loc, " ", "-")))
}

func appendPostAJobSEOLandingPageForLocation(seoLandingPages []string, loc string) []string {
	tmpl := `Hire-Golang-Developers-In-%s`
	if strings.ToLower(loc) == "remote" {
		return append(seoLandingPages, `Hire-Remote-Golang-Developers`)
	}
	return append(seoLandingPages, fmt.Sprintf(tmpl, strings.ReplaceAll(loc, " ", "-")))
}

func GenerateSearchSEOLandingPages(conn *sql.DB) ([]database.SEOLandingPage, error) {
	var seoLandingPages []database.SEOLandingPage
	locs, err := database.GetSEOLocations(conn)
	if err != nil {
		return seoLandingPages, err
	}
	skills, err := database.GetSEOskills(conn)
	if err != nil {
		return seoLandingPages, err
	}

	for _, loc := range locs {
		seoLandingPages = appendSearchSEOLandingPageForLocationAndSkill(seoLandingPages, loc, database.SEOSkill{})
		// for _, skill := range skills {
		// 	seoLandingPages = appendSearchSEOLandingPageForLocationAndSkill(seoLandingPages, loc, skill)
		// }
	}
	for _, skill := range skills {
		seoLandingPages = appendSearchSEOLandingPageForLocationAndSkill(seoLandingPages, database.SEOLocation{}, skill)
	}

	return seoLandingPages, nil
}

func appendSearchSEOLandingPageForLocationAndSkill(seoLandingPages []database.SEOLandingPage, loc database.SEOLocation, skill database.SEOSkill) []database.SEOLandingPage {
	templateBoth := `Golang-%s-Jobs-In-%s`
	templateSkill := `Golang-%s-Jobs`
	templateLoc := `Golang-Jobs-In-%s`

	templateRemoteLoc := `Remote-Golang-Jobs`
	templateRemoteBoth := `Remote-Golang-%s-Jobs`
	loc.Name = strings.ReplaceAll(loc.Name, " ", "-")
	skill.Name = strings.ReplaceAll(skill.Name, " ", "-")

	// Skill only
	if loc.Name == "" {
		return append(seoLandingPages, database.SEOLandingPage{
			URI:   fmt.Sprintf(templateSkill, skill.Name),
			Skill: skill.Name,
		})
	}

	// Remote is special case
	if loc.Name == "Remote" {
		if skill.Name != "" {
			return append(seoLandingPages, database.SEOLandingPage{
				URI:      fmt.Sprintf(templateRemoteBoth, skill.Name),
				Location: loc.Name,
			})
		} else {
			return append(seoLandingPages, database.SEOLandingPage{
				URI:      templateRemoteLoc,
				Location: loc.Name,
				Skill:    skill.Name,
			})
		}
	}

	// Location only
	if skill.Name == "" {
		return append(seoLandingPages, database.SEOLandingPage{
			URI:      fmt.Sprintf(templateLoc, loc.Name),
			Location: loc.Name,
		})
	}

	// Both
	return append(seoLandingPages, database.SEOLandingPage{
		URI:      fmt.Sprintf(templateBoth, skill.Name, loc.Name),
		Skill:    skill.Name,
		Location: loc.Name,
	})
}
