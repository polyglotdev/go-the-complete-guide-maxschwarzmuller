package main

import "fmt"

type UdemyMap map[string]float64

func main() {
	cloudProviders := map[string]string{
		"aws":     "Amazon Web Services",
		"gcp":     "Google Cloud Platform",
		"az":      "Microsoft Azure",
		"ibm":     "IBM Cloud",
		"alibaba": "Alibaba Cloud",
	}
	for key, value := range cloudProviders {
		fmt.Printf("%s:, %s\n", key, value)
	}

	fmt.Println("------------------------------")

	userNames := map[string]string{
		"elijah": "Elijah Hallan",
		"ezra":   "Ezra Hallan",
	}
	for key, value := range userNames {
		fmt.Printf("%s:, %s\n", key, value)
	}
	// add to usernames
	userNames["dom"] = "Dom Hallan"
	for key, value := range userNames {
		fmt.Printf("%s:, %s\n", key, value)
	}

	nbaTeams := []string{"Atlanta Hawks", "Boston Celtics", "Brooklyn Nets", "Charlotte Hornets", "Chicago Bulls", "Cleveland Cavaliers", "Dallas Mavericks", "Denver Nuggets", "Detroit Pistons", "Golden State Warriors", "Houston Rockets", "Indiana Pacers", "Los Angeles Clippers", "Los Angeles Lakers", "Memphis Grizzlies", "Miami Heat", "Milwaukee Bucks", "Minnesota Timberwolves", "New Orleans Pelicans", "New York Knicks", "Oklahoma City Thunder", "Orlando Magic", "Philadelphia 76ers", "Phoenix Suns", "Portland Trail Blazers", "Sacramento Kings", "San Antonio Spurs", "Toronto Raptors", "Utah Jazz", "Washington Wizards"}
	for i, value := range nbaTeams {
		fmt.Printf("%d:, %s\n", i, value)
	}
	fmt.Println("---------------------------------------------")
	courseRatings := map[string]float64{
		"Go Fundamentals":         4.8,
		"Python Fundamentals":     4.7,
		"Java Fundamentals":       4.6,
		"JavaScript Fundamentals": 4.5,
		"Ruby Fundamentals":       4.4,
	}
	for i, value := range courseRatings {
		fmt.Printf("Course: %s\nScore: %.1f\n\n", i, value)
	}

	fmt.Println("---------------------------------------------")
	values := make([]string, 0, len(cloudProviders))
	values = append(values, "Amazon Web Services")
	values = append(values, "Google Cloud Platform")
	values = append(values, "Microsoft Azure")
	values = append(values, "IBM Cloud")
	values = append(values, "Alibaba Cloud")
	for i, value := range values {
		fmt.Printf("%d:, %s\n", i+1, value)
	}

	fmt.Println("---------------------------------------------")
	udemyCourses := make(UdemyMap)
	udemyCourses["Go the Complete Guide"] = 4.8
	udemyCourses["Python Data Structures"] = 4.7
	udemyCourses["All the Java your ass can handle"] = 4.6
	udemyCourses["JavaScript Bootcamp is for the ♥️"] = 4.5
	udemyCourses["Ruby Basics for the the Gangstas"] = 4.4
	for i, courses := range udemyCourses {
		fmt.Printf("Course: %s\nRating: %.1f\n\n", i, courses)
	}
}
