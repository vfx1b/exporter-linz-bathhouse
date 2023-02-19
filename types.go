package main

type App struct {
	BindAddr string
	BindPort string
}

type Location struct {
	ServerSideId int64
	Name         string
}

type Payload struct {
	Id                   int64 `json:"id"`
	RelativeCurrCapacity int64 `json:"relativeCurrCapacity"`
}

func findLocationNameForServerSideId(id int64) *Location {
	for _, location := range locations {
		if location.ServerSideId == id {
			return &location
		}
	}
	return nil
}
