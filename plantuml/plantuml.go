package plantuml

import (
	"github.com/hekonsek/plant4/model"
	"strings"
)

type Plotter struct {

}

func (*Plotter) Plot(landscape *model.Model) ([]string, error) {
	shared := map[string]bool{}
	for _, s := range landscape.Shared {
		shared[s.Name] = true
	}

	landscapeDiagram := "@startuml\n"
	if landscape.Name != "" {
		landscapeDiagram += "title System view: " + landscape.Name + "\n"
	}
	for _, system := range landscape.Systems {
		landscapeDiagram += "node " + system.Name + "\n"
	}
	for _, system := range landscape.Systems {
		for _, relation := range system.Relations {
			landscapeDiagram += system.Name + "-->" + relation.Target + " : " + relation.Name + "\n"
		}
	}
	landscapeDiagram +=  "@enduml"

	systemDiagrams := []string{}
	for _, system := range landscape.Systems {
		containerDiagram := "@startuml\n"
		containerDiagram += "title Container view for system: " + system.Name + "\n"
		containerDiagram += "node " + system.Name + " {\n"
		for _, container := range system.Containers {
			containerDiagram += "frame \"" + container.Name +  "\" as " + strings.ReplaceAll(container.Name, " ", "_") + "\n"
		}
		for _, container := range system.Containers {
			for _, relation := range container.Relations {
				if !shared[relation.Target] {
					containerDiagram += strings.ReplaceAll(container.Name, " ", "_") + "-->" + strings.ReplaceAll(relation.Target, " ", "_") + " : " + relation.Name + "\n"
				}
			}
		}
		containerDiagram += "}\n"
		for _, shared := range landscape.Shared {
			addedShared := map[string]bool{}
			for _, container := range system.Containers {
				for _, relation := range container.Relations {
					if relation.Target == shared.Name {
						if !addedShared[shared.Name] {
							containerDiagram += "frame \"" + shared.Name + "\" as " + strings.ReplaceAll(shared.Name, " ", "_") + "\n"
							addedShared[shared.Name] = true
						}
						containerDiagram += strings.ReplaceAll(container.Name, " ", "_") + "-->" + strings.ReplaceAll(relation.Target, " ", "_") + " : " + relation.Name + "\n"
					}
				}
			}
		}
		containerDiagram +=  "@enduml"
		systemDiagrams = append(systemDiagrams, containerDiagram)
	}

 	return append([]string{landscapeDiagram}, systemDiagrams...), nil
}