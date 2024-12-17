package main

func getResponses(name string) ([]string, error) {
	return []string{"Hello " + name, "Bye " + name}, nil
}
