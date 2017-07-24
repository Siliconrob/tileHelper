package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"math/rand"
    "time"
)

type tomlConfig struct {
	Title   string
	Management mgrInfo
	TileServers tileGenerators
}

type mgrInfo struct {
	Url string
}

type tileGenerators struct {
	ImageExt string
	RequestFormat string
	Hosts []int
}

type TileRequest struct {
	Format string
	Host int
	X int
	Y int
	Z int
	ImageType string	
}

func mutate(inputs []int) []int {
	newSet := inputs
    for i := range newSet {
        j := rand.Intn(i + 1)
        newSet[i], newSet[j] = newSet[j], newSet[i]
    }
    return newSet
}

func (req TileRequest) String() string {
	return fmt.Sprintf(req.Format, req.Host, req.Z, req.X, req.Y, req.ImageType)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var config tomlConfig
	if _, err := toml.DecodeFile("tileHelper.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Management url: %s\n", config.Management.Url)
	fmt.Printf("Image format %s\n", config.TileServers.ImageExt)

	a := TileRequest { config.TileServers.RequestFormat,
		mutate(config.TileServers.Hosts)[0],
		1, 2, 3,
		config.TileServers.ImageExt };
	fmt.Println(a);
}