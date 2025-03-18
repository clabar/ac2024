package nm

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func CreateGraph() {
	cityHash := func(c City) string {
		return c.Name
	}
	g := graph.New(cityHash, graph.Directed(), graph.Weighted())
	allDistances := graph.New(cityHash, graph.Directed(), graph.Weighted())

	fmt.Println("adding cities")
	populateCity(g)
	populateCity(allDistances)
	fmt.Println("adding edges")
	addEdges(g)

	fmt.Println("print")
	var wg sync.WaitGroup
	for i := 0; i < len(citta); i++ {
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			fmt.Println(fmt.Sprintf("Index = %d computing for %s\n", ind, citta[ind].Name))
			s := "Routes from " + citta[ind].Name
			for j := ind + 1; j < len(citta); j++ {
				path, _ := graph.ShortestPath(g, citta[ind].Name, citta[j].Name)
				l := len(path) - 1
				mustAddFullDuplexEdge(allDistances, path[0], path[l], l)
				s += fmt.Sprintf("from %s to %s len:%d \n", path[0], path[l], l)
				time.Sleep(1 * time.Millisecond)
			}
			fmt.Println(s + "\n")
		}(i)
	}
	wg.Wait()
	fmt.Println("waiting for all to complete")
	//path, _ := graph.ShortestPath(g, Oslo.Name, Bergen.Name)
	//fmt.Println(path)

	file, err := os.Create("./simple.gv")
	if err != nil {
		panic(err)
	}
	err = draw.DOT(g, file)
	if err != nil {
		panic(err)
	}
	allF, err := os.Create("./allDistances.gv")
	if err != nil {
		panic(err)
	}
	_ = draw.DOT(g, allF)
}

func addEdges(g graph.Graph[string, City]) {
	mustAddCarEdge(g, Oslo.Name, Lillestrøm.Name)
	mustAddCarEdge(g, Oslo.Name, Ski.Name)
	mustAddCarEdge(g, Oslo.Name, Drammen.Name)

	mustAddCarEdge(g, Lillestrøm.Name, Eidsvoll.Name)
	mustAddCarEdge(g, Ski.Name, Askim.Name)
	mustAddCarEdge(g, Askim.Name, Sarpsborg.Name)
	mustAddCarEdge(g, Sarpsborg.Name, Halden.Name)
	mustAddCarEdge(g, Sarpsborg.Name, Fredrikstad.Name)
	mustAddCarEdge(g, Ski.Name, Drøbak.Name)
	mustAddCarEdge(g, Drøbak.Name, Moss.Name)
	mustAddCarEdge(g, Moss.Name, Fredrikstad.Name)

	mustAddCarEdge(g, Drammen.Name, Kongsberg.Name)
	mustAddCarEdge(g, Drammen.Name, Hønefoss.Name)
	mustAddCarEdge(g, Drammen.Name, Drøbak.Name)

	mustAddCarEdge(g, Kongsberg.Name, Tønsberg.Name)
	mustAddCarEdge(g, Kongsberg.Name, Notodden.Name)
	mustAddCarEdge(g, Tønsberg.Name, Horten.Name)
	mustAddCarEdge(g, Tønsberg.Name, Sandefjord.Name)
	mustAddCarEdge(g, Sandefjord.Name, Larvik.Name)
	mustAddCarEdge(g, Larvik.Name, Porsgrunn.Name)

	mustAddCarEdge(g, Notodden.Name, Skien.Name)
	mustAddCarEdge(g, Notodden.Name, Bø.Name)
	mustAddCarEdge(g, Notodden.Name, Rjukan.Name)
	mustAddCarEdge(g, Porsgrunn.Name, Skien.Name)
	mustAddCarEdge(g, Porsgrunn.Name, Kragerø.Name)

	mustAddCarEdge(g, Kragerø.Name, Risør.Name)
	mustAddCarEdge(g, Arendal.Name, Risør.Name)
	mustAddCarEdge(g, Arendal.Name, Grimstad.Name)
	mustAddCarEdge(g, Byglandsfjord.Name, Grimstad.Name)
	mustAddCarEdge(g, Byglandsfjord.Name, Bø.Name)

	mustAddCarEdge(g, Rjukan.Name, Odda.Name)
	mustAddCarEdge(g, Geilo.Name, Odda.Name)
	mustAddCarEdge(g, Geilo.Name, Hønefoss.Name)
	mustAddCarEdge(g, Geilo.Name, Gol.Name)
	mustAddCarEdge(g, Geilo.Name, Lærdal.Name)
	mustAddCarEdge(g, Sogndal.Name, Lærdal.Name)
	mustAddCarEdge(g, Sogndal.Name, Førde.Name)
	mustAddCarEdge(g, Florø.Name, Førde.Name)
	mustAddCarEdge(g, Voss.Name, Lærdal.Name)
	mustAddCarEdge(g, Voss.Name, Bergen.Name)
	mustAddCarEdge(g, Bergen.Name, Osøyro.Name)
	mustAddCarEdge(g, Bergen.Name, Askøy.Name)

	mustAddCarEdge(g, Kristiansand.Name, Grimstad.Name)
	mustAddCarEdge(g, Kristiansand.Name, Mandal.Name)
	mustAddCarEdge(g, Kristiansand.Name, Vennesla.Name)
	mustAddCarEdge(g, Flekkefjord.Name, Vennesla.Name)
	mustAddCarEdge(g, Flekkefjord.Name, Egersund.Name)
	mustAddCarEdge(g, Bryne.Name, Egersund.Name)
	mustAddCarEdge(g, Bryne.Name, Sandnes.Name)
	mustAddCarEdge(g, Stavanger.Name, Sandnes.Name)

	mustAddCarEdge(g, Raufoss.Name, Hønefoss.Name)
	mustAddCarEdge(g, Raufoss.Name, Gjøvik.Name)
	mustAddCarEdge(g, Eidsvoll.Name, Hurdal.Name)
	mustAddCarEdge(g, Eidsvoll.Name, Kongsvinger.Name)
	mustAddCarEdge(g, Eidsvoll.Name, Hamar.Name)
	mustAddCarEdge(g, Elverum.Name, Hamar.Name)
	mustAddCarEdge(g, Elverum.Name, Kongsvinger.Name)
	mustAddCarEdge(g, Elverum.Name, Trysil.Name)
	mustAddCarEdge(g, Røros.Name, Trysil.Name)
	mustAddCarEdge(g, Røros.Name, Melhus.Name)
	mustAddCarEdge(g, Hamar.Name, Brumunddal.Name)
	mustAddCarEdge(g, Lillehammer.Name, Brumunddal.Name)
	mustAddCarEdge(g, Lillehammer.Name, Hafjell.Name)
	mustAddCarEdge(g, Dovre.Name, Hafjell.Name)
	mustAddCarEdge(g, Dovre.Name, Volda.Name)
	mustAddCarEdge(g, Måløy.Name, Volda.Name)
	mustAddCarEdge(g, Ålesund.Name, Volda.Name)
	mustAddCarEdge(g, Dovre.Name, Melhus.Name)

	// Troendelag
	mustAddCarEdge(g, Melhus.Name, Trondheim.Name)
	mustAddCarEdge(g, Orkanger.Name, Trondheim.Name)
	mustAddCarEdge(g, Orkanger.Name, Kristiansund.Name)
	mustAddCarEdge(g, Orkanger.Name, Molde.Name)

	mustAddCarEdge(g, Malvik.Name, Trondheim.Name)
	mustAddCarEdge(g, Malvik.Name, Stjørdal.Name)
	mustAddCarEdge(g, Levanger.Name, Stjørdal.Name)
	mustAddCarEdge(g, Levanger.Name, Steinkjer.Name)
	mustAddCarEdge(g, Namsos.Name, Steinkjer.Name)
	mustAddCarEdge(g, Namsos.Name, Brønnøysund.Name)

	mustAddCarEdge(g, Mo.Name, Brønnøysund.Name)
	mustAddCarEdge(g, Mo.Name, Bodø.Name)
	mustAddCarEdge(g, Narvik.Name, Bodø.Name)
	mustAddCarEdge(g, Narvik.Name, Harstad.Name)
	mustAddCarEdge(g, Narvik.Name, Bardufoss.Name)

	mustAddCarEdge(g, Svolvær.Name, Henningsvær.Name)
	mustAddCarEdge(g, Henningsvær.Name, Å.Name)

	mustAddCarEdge(g, Tromsø.Name, Bardufoss.Name)
	mustAddCarEdge(g, Golddajávri.Name, Bardufoss.Name)
	mustAddCarEdge(g, Golddajávri.Name, Storslett.Name)
	mustAddCarEdge(g, Golddajávri.Name, Alta.Name)
	mustAddCarEdge(g, Kautokeino.Name, Alta.Name)
	mustAddCarEdge(g, Kautokeino.Name, Karasjok.Name)
	mustAddCarEdge(g, Lakselv.Name, Karasjok.Name)
	mustAddCarEdge(g, Lakselv.Name, Alta.Name)
	mustAddCarEdge(g, Lakselv.Name, Honningsvåg.Name)
	mustAddCarEdge(g, Lakselv.Name, Vadsø.Name)
	mustAddCarEdge(g, Kirkenes.Name, Vadsø.Name)
	mustAddCarEdge(g, Vardø.Name, Vadsø.Name)

	mustAddNaval(g, Haugesund.Name, Osøyro.Name, 2)
	mustAddNaval(g, Askøy.Name, Florø.Name, 3)
	mustAddNaval(g, Måløy.Name, Florø.Name, 1)
	mustAddNaval(g, Ålesund.Name, Molde.Name, 2)
	mustAddNaval(g, Hamar.Name, Gjøvik.Name, 1)
	mustAddNaval(g, Røst.Name, Å.Name, 1)
	mustAddCarEdge(g, Svolvær.Name, Å.Name)
	mustAddNaval(g, Svolvær.Name, Bodø.Name, 2)
	mustAddNaval(g, Svolvær.Name, Harstad.Name, 4)
	mustAddNaval(g, Finnsnes.Name, Harstad.Name, 2)
	mustAddNaval(g, Finnsnes.Name, Tromsø.Name, 1)
	mustAddNaval(g, Alta.Name, Hammerfest.Name, 2)
	mustAddNaval(g, Stavanger.Name, Haugesund.Name, 1)
	mustAddNaval(g, Lyngør.Name, Risør.Name, 1)
	mustAddNaval(g, Horten.Name, Moss.Name, 1)

	mustAddAirEdge(g, Oslo.Name, Trondheim.Name, 1)
	mustAddAirEdge(g, Oslo.Name, Tromsø.Name, 2)

	mustAddAirEdge(g, Stavanger.Name, Trondheim.Name, 2)
	mustAddAirEdge(g, Stavanger.Name, Honningsvåg.Name, 3)

	mustAddAirEdge(g, Bergen.Name, Trondheim.Name, 1)
	mustAddAirEdge(g, Bergen.Name, Tromsø.Name, 3)

	mustAddAirEdge(g, Trondheim.Name, Oslo.Name, 2)
	mustAddAirEdge(g, Trondheim.Name, Stavanger.Name, 1)
	mustAddAirEdge(g, Trondheim.Name, Bergen.Name, 1)
	mustAddAirEdge(g, Trondheim.Name, Honningsvåg.Name, 2)

	mustAddAirEdge(g, Tromsø.Name, Oslo.Name, 3)
	mustAddAirEdge(g, Tromsø.Name, Stavanger.Name, 2)
	mustAddAirEdge(g, Tromsø.Name, Bergen.Name, 2)

	mustAddAirEdge(g, Honningsvåg.Name, Trondheim.Name, 2)
	mustAddAirEdge(g, Honningsvåg.Name, Stavanger.Name, 3)
	mustAddAirEdge(g, Honningsvåg.Name, Bergen.Name, 3)

}

func mustAddCarEdge(g graph.Graph[string, City], from, to string) {
	mustAddFullDuplexEdge(g, from, to, carWeight)
}

func mustAddFullDuplexEdge(g graph.Graph[string, City], from, to string, weight int) {
	err := g.AddEdge(from, to, graph.EdgeWeight(weight))
	if err != nil {
		panic(fmt.Sprintf("impossibe to create edge %s->%s: %s", from, to, err))
	}
	err = g.AddEdge(to, from, graph.EdgeWeight(weight))
	if err != nil {
		panic(fmt.Sprintf("impossibe to create edge %s->%s: %s", from, to, err))
	}
}

func mustAddAirEdge(g graph.Graph[string, City], from, to string, multiplier int) {
	err := g.AddEdge(from, to, graph.EdgeWeight(multiplier*airWeight), graph.EdgeAttribute("label", "air-connection"))
	if err != nil {
		panic(fmt.Sprintf("impossibe to create edge %s->%s: %s", from, to, err))
	}
}

func mustAddNaval(g graph.Graph[string, City], from, to string, multiplier int) {
	err := g.AddEdge(from, to, graph.EdgeWeight(multiplier*navalWeight), graph.EdgeAttribute("label", "sea-connection"))
	if err != nil {
		panic(fmt.Sprintf("impossibe to create edge %s->%s: %s", from, to, err))
	}
	err = g.AddEdge(to, from, graph.EdgeWeight(multiplier*navalWeight), graph.EdgeAttribute("label", "sea-connection"))
	if err != nil {
		panic(fmt.Sprintf("impossibe to create edge %s->%s: %s", from, to, err))
	}
}

func populateCity(g graph.Graph[string, City]) {
	for _, city := range citta {
		if city.Type == Aereoporto {
			g.AddVertex(city,
				graph.VertexAttribute("colorscheme", "greens3"),
				graph.VertexAttribute("style", "filled"),
				graph.VertexAttribute("color", "2"),
				graph.VertexAttribute("fillcolor", "1"),
			)
		} else if city.Type == Porto {
			g.AddVertex(city,
				graph.VertexAttribute("colorscheme", "blues3"),
				graph.VertexAttribute("style", "filled"),
				graph.VertexAttribute("color", "2"),
				graph.VertexAttribute("fillcolor", "1"),
			)
		} else {
			g.AddVertex(city,
				graph.VertexAttribute("colorscheme", "purples3"),
				graph.VertexAttribute("style", "filled"),
				graph.VertexAttribute("color", "2"),
				graph.VertexAttribute("fillcolor", "1"),
			)
		}
	}
}

type City struct {
	Name string
	Type CityType `default:"d"`
}

var citta = []City{
	Oslo,
	Bergen,
	Trondheim,
	Stavanger,
	Tromsø,
	Honningsvåg,
	Røros,
	Orkanger,
	Malvik,
	Namsos,
	Steinkjer,
	Stjørdal,
	Levanger,
	Melhus,
	Volda,
	Kristiansund,
	Molde,
	Haugesund,
	Egersund,
	Bryne,
	Sandnes,
	Rjukan,
	Notodden,
	Kragerø,
	Grimstad,
	Lyngør,
	Risør,
	Arendal,
	Mandal,
	Flekkefjord,
	Odda,
	Måløy,
	Lærdal,
	Sogndal,
	Førde,
	Florø,
	Voss,
	Askøy,
	Osøyro,
	Byglandsfjord,
	Bø,
	Harstad,
	Finnsnes,
	Bardufoss,
	Lakselv,
	Golddajávri,
	Storslett,
	Alta,
	Hammerfest,
	Kirkenes,
	Karasjok,
	Vardø,
	Å,
	Henningsvær,
	Brønnøysund,
	Røst,
	Svolvær,
	Kongsvinger,
	Brumunddal,
	Trysil,
	Elverum,
	Hamar,
	Hafjell,
	Dovre,
	Raufoss,
	Gjøvik,
	Vadsø,
	Narvik,
	Bodø,
	Kautokeino,
	Lillehammer,
	Mo,
	Halden,
	Fredrikstad,
	Vennesla,
	Kristiansand,
	Larvik,
	Sandefjord,
	Horten,
	Tønsberg,
	Skien,
	Porsgrunn,
	Lillestrøm,
	Drøbak,
	Ski,
	Askim,
	Moss,
	Sarpsborg,
	Ålesund,
	Kongsberg,
	Geilo,
	Gol,
	Hønefoss,
	Drammen,
	Hurdal,
	Eidsvoll,
}

var (
	Oslo          = City{Name: "Oslo", Type: Aereoporto}
	Bergen        = City{Name: "Bergen", Type: Aereoporto}
	Trondheim     = City{Name: "Trondheim", Type: Aereoporto}
	Stavanger     = City{Name: "Stavanger", Type: Aereoporto}
	Tromsø        = City{Name: "Tromsø", Type: Aereoporto}
	Honningsvåg   = City{Name: "Honningsvåg", Type: Aereoporto}
	Røros         = City{Name: "Røros"}
	Orkanger      = City{Name: "Orkanger"}
	Malvik        = City{Name: "Malvik"}
	Namsos        = City{Name: "Namsos"}
	Steinkjer     = City{Name: "Steinkjer"}
	Stjørdal      = City{Name: "Stjørdal"}
	Levanger      = City{Name: "Levanger"}
	Melhus        = City{Name: "Melhus"}
	Volda         = City{Name: "Volda"}
	Kristiansund  = City{Name: "Kristiansund"}
	Molde         = City{Name: "Molde"}
	Haugesund     = City{Name: "Haugesund"}
	Egersund      = City{Name: "Egersund"}
	Bryne         = City{Name: "Bryne"}
	Sandnes       = City{Name: "Sandnes"}
	Rjukan        = City{Name: "Rjukan"}
	Notodden      = City{Name: "Notodden"}
	Kragerø       = City{Name: "Kragerø"}
	Grimstad      = City{Name: "Grimstad"}
	Lyngør        = City{Name: "Lyngør"}
	Risør         = City{Name: "Risør"}
	Arendal       = City{Name: "Arendal"}
	Mandal        = City{Name: "Mandal"}
	Flekkefjord   = City{Name: "Flekkefjord"}
	Odda          = City{Name: "Odda"}
	Måløy         = City{Name: "Måløy"}
	Lærdal        = City{Name: "Lærdal"}
	Sogndal       = City{Name: "Sogndal"}
	Førde         = City{Name: "Førde"}
	Florø         = City{Name: "Florø"}
	Voss          = City{Name: "Voss"}
	Askøy         = City{Name: "Askøy"}
	Osøyro        = City{Name: "Osøyro"}
	Byglandsfjord = City{Name: "Byglandsfjord"}
	Bø            = City{Name: "Bø"}
	Harstad       = City{Name: "Harstad"}
	Finnsnes      = City{Name: "Finnsnes"}
	Bardufoss     = City{Name: "Bardufoss"}
	Lakselv       = City{Name: "Lakselv"}
	Golddajávri   = City{Name: "Golddajávri"}
	Storslett     = City{Name: "Storslett"}
	Alta          = City{Name: "Alta"}
	Hammerfest    = City{Name: "Hammerfest"}
	Kirkenes      = City{Name: "Kirkenes"}
	Karasjok      = City{Name: "Karasjok"}
	Vardø         = City{Name: "Vardø"}
	Å             = City{Name: "Å"}
	Henningsvær   = City{Name: "Henningsvær"}
	Brønnøysund   = City{Name: "Brønnøysund"}
	Røst          = City{Name: "Røst"}
	Svolvær       = City{Name: "Svolvær"}
	Kongsvinger   = City{Name: "Kongsvinger"}
	Brumunddal    = City{Name: "Brumunddal"}
	Trysil        = City{Name: "Trysil"}
	Elverum       = City{Name: "Elverum"}
	Hamar         = City{Name: "Hamar"}
	Hafjell       = City{Name: "Hafjell"}
	Dovre         = City{Name: "Dovre"}
	Raufoss       = City{Name: "Raufoss"}
	Gjøvik        = City{Name: "Gjøvik"}
	Vadsø         = City{Name: "Vadsø"}
	Narvik        = City{Name: "Narvik"}
	Bodø          = City{Name: "Bodø"}
	Kautokeino    = City{Name: "Kautokeino"}
	Lillehammer   = City{Name: "Lillehammer"}
	Mo            = City{Name: "Mo i Rana"}
	Halden        = City{Name: "Halden"}
	Fredrikstad   = City{Name: "Fredrikstad"}
	Vennesla      = City{Name: "Vennesla"}
	Kristiansand  = City{Name: "Kristiansand"}
	Larvik        = City{Name: "Larvik"}
	Sandefjord    = City{Name: "Sandefjord"}
	Horten        = City{Name: "Horten"}
	Tønsberg      = City{Name: "Tønsberg"}
	Skien         = City{Name: "Skien"}
	Porsgrunn     = City{Name: "Porsgrunn"}
	Lillestrøm    = City{Name: "Lillestrøm"}
	Drøbak        = City{Name: "Drøbak"}
	Ski           = City{Name: "Ski"}
	Askim         = City{Name: "Askim"}
	Moss          = City{Name: "Moss"}
	Sarpsborg     = City{Name: "Sarpsborg"}
	Ålesund       = City{Name: "Ålesund"}
	Kongsberg     = City{Name: "Kongsberg"}
	Geilo         = City{Name: "Geilo"}
	Gol           = City{Name: "Gol"}
	Hønefoss      = City{Name: "Hønefoss"}
	Drammen       = City{Name: "Drammen"}
	Hurdal        = City{Name: "Hurdal"}
	Eidsvoll      = City{Name: "Eidsvoll"}
)

type CityType string

const (
	carWeight   = 287
	navalWeight = 575
	airWeight   = 820

	Aereoporto CityType = "a"
	Default    CityType = "d"
	Porto      CityType = "p"
)
