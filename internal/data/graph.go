package data

import (
	"sync"
)

type neighbor struct {
	name   string
	weight uint
}

type Route struct {
	Source string
	Length uint
	Target string
}

type Graph struct {
	mutex  *sync.Mutex
	Nodes  map[string]struct{}
	Routes map[string]*Route
}

func (g *Graph) AddNodes(newNodes []string) (bool, string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for _, nodeName := range newNodes {
		_, found := g.Nodes[nodeName]
		if found {
			return false, nodeName
		}
		g.Nodes[nodeName] = struct{}{}
	}
	return true, ""
}

func (g *Graph) AddRoutes(routes map[string]*Route) (bool, string, string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for routeName, route := range routes {
		_, found := g.Routes[routeName]
		if found {
			return false, routeName, ""
		}
		_, aFound := g.Nodes[route.Source]
		if !aFound {
			return false, "", route.Source
		}
		_, bFound := g.Nodes[route.Target]
		if !bFound {
			return false, "", route.Target
		}
		g.Routes[routeName] = route
	}
	return true, "", ""
}

func (g *Graph) DeleteNodes(newNodes []string) (bool, string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for _, nodeName := range newNodes {
		_, found := g.Nodes[nodeName]
		if !found {
			return false, nodeName
		}
		delete(g.Nodes, nodeName)
	}
	return true, ""
}

func (g *Graph) DeleteRoutes(routes []string) (bool, string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for _, routeName := range routes {
		if _, found := g.Routes[routeName]; !found {
			return false, routeName
		}
		delete(g.Routes, routeName)
	}
	return true, ""
}

// http://www.gitta.info/Accessibiliti/en/html/Dijkstra_learningObject1.html
func (g *Graph) Dijkstra(start, target string) ([]string, string) {
	// ToDo: Check if start and target are int the graph

	dist := map[string]uint{}
	Q := map[string]struct{}{}
	previous := map[string]string{}

	neighbors := map[string][]neighbor{}

	g.mutex.Lock()

	done := make(chan bool, 1)

	go func() {
		for nodeName := range g.Nodes {
			previous[nodeName] = ""
			if nodeName == start {
				continue
			}
			dist[nodeName] = ^uint(0)
			Q[nodeName] = struct{}{}
		}
		dist[start] = 0
		Q[start] = struct{}{}
		done <- true
	}()

	for _, route := range g.Routes {
		_, aFound := neighbors[route.Source]
		if !aFound {
			neighbors[route.Source] = []neighbor{}
		}

		neighbors[route.Source] = append(neighbors[route.Source],
			neighbor{
				name:   route.Target,
				weight: route.Length,
			},
		)
	}

	<-done

	g.mutex.Unlock()
	qLength := len(Q)
	for qLength > 0 {
		u := ""
		firstNode := true
		for node := range Q {
			if firstNode {
				firstNode = false
				u = node
				continue
			}
			if dist[node] <= dist[u] {
				u = node
			}
		}
		qLength--
		delete(Q, u)

		for _, v := range neighbors[u] {
			_, found := Q[v.name]
			if found {
				alt := dist[u] + v.weight
				if alt < dist[v.name] {
					dist[v.name] = alt
					previous[v.name] = u
				}
			}
		}
	}
	result := []string{target}
	current := target
	for {
		before := previous[current]
		if before == "" {
			break
		}
		result = append(result, before)
		current = before

	}
	return result, ""
}

func NewGraph() *Graph {
	return &Graph{
		mutex:  new(sync.Mutex),
		Nodes:  map[string]struct{}{},
		Routes: map[string]*Route{},
	}
}
