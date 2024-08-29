package models

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var Albums = []Album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "Suegros", Artist: "Claudio", Year: 2023},
	{ID: "3", Title: "dsa", Artist: "Camila", Year: 2024},
	{ID: "4", Title: "ssssssss", Artist: "Cabello", Year: 2020},
}
