package main

import "fmt"

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func main() {

	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)
	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

}
