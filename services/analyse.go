package services

import (
	"main/models"
	"strings"
)

type Analyse struct {
}

type AnalyseI interface {
	Analyse(entries []models.Entry) (catagories map[string]float64)
}

func (a Analyse) Analyse(entries []models.Entry) (catagories map[string]float64) {
	catagories = make(map[string]float64)

	lookup := getCategoryLookup()
	for _, entry := range entries {
		words := strings.Split(entry.Line, " ")
		lastWord := ""
		found := false
		for _, word := range words {
			word = strings.ToLower(word)
			if lastWord != "" {
				// Lookup two words together
				twoWordKey := lastWord + " " + word
				if category, exists := lookup[twoWordKey]; exists {
					catagories[category] += entry.Amount
					found = true
					break
				}
			}
			lastWord = word
		}
		if found {
			continue
		}
		for _, word := range words {
			word = strings.ToLower(word)
			// Lookup one word
			if category, exists := lookup[word]; exists {
				catagories[category] += entry.Amount
				found = true
				break
			}
		}
		if !found {
			// If not found then add the line as a category. This makes it easy to see whats not been found on the results page
			catagories[entry.Line] = entry.Amount
		}
	}

	return catagories
}

func getCategoryLookup() (lookup map[string]string) {
	// Added like this to make it easy to maintain
	categories := map[string][]string{
		"Groceries":     {"lidl", "tesco", "morrisons", "iceland", "sainsburys", "aldi", "waitrose", "co-op", "asda", "spar", "kwikimart", "grocery", "bakery", "portsmouth arena"},
		"Transport":     {"bus", "train", "uber", "taxi"},
		"Entertainment": {"cinema", "restaurant", "pub", "bar", "john baker"},
		"Restaurants":   {"pizza", "burger", "kebab", "chinese", "indian", "blue cobra", "subway", "dominos", "kfc", "mcdonalds", "pizza hut", "nandos", "taco bell", "wetherspoons", "wagamama", "takeaway"},
		"Utilities":     {"lebara", "water", "gas", "internet"},
		"Health":        {"pharmacy", "health", "boots", "superdrug"},
		"Subscriptions": {"netflix", "spotify", "amazon prime", "disney plus", "apple music", "google play"},
		"Amazon":        {"amazon"},
		"Homeware":      {"b&m", "other", "unknown", "home bargains", "poundland", "temu", "screwfix"},
		"Car":           {"protyre", "petrol", "shell", "esso", "bp", "tesco petrol", "morrisons petrol", "asda petrol", "sainsburys petrol", "car insurance", "car tax", "car maintenance"},
		"School":        {"scopay"},
		"Hotels":        {"hotel", "travelodge", "premier inn", "ibis", "marriott", "hilton", "holiday inn", "radisson", "best western"},
	}

	lookup = map[string]string{}

	for category, keywords := range categories {
		for _, keyword := range keywords {
			lookup[keyword] = category
		}
	}
	return
}
