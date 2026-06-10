package main

type Dictionary map[string]string

func (d Dictionary) Serach(word string) string {
	return d[word]
}
