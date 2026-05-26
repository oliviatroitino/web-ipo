package content

import (
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// LoadTecnicas lee todos los archivos .yaml del directorio dado
// y devuelve un slice de Tecnica ordenado por primera fase y nombre.
func LoadTecnicas(dir string) ([]Tecnica, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var tecnicas []Tecnica
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}
		// Ignorar la plantilla
		if entry.Name() == "plantilla.yaml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}

		var t Tecnica
		if err := yaml.Unmarshal(data, &t); err != nil {
			return nil, err
		}
		tecnicas = append(tecnicas, t)
	}

	// Ordenar por primera fase y luego por nombre
	sort.Slice(tecnicas, func(i, j int) bool {
		faseI := ""
		if len(tecnicas[i].Fases) > 0 {
			faseI = tecnicas[i].Fases[0]
		}
		faseJ := ""
		if len(tecnicas[j].Fases) > 0 {
			faseJ = tecnicas[j].Fases[0]
		}
		if faseI != faseJ {
			return faseI < faseJ
		}
		return tecnicas[i].Nombre < tecnicas[j].Nombre
	})

	return tecnicas, nil
}

// FindByID busca una técnica por su ID. Devuelve nil si no existe.
func FindByID(tecnicas []Tecnica, id string) *Tecnica {
	for i := range tecnicas {
		if tecnicas[i].ID == id {
			return &tecnicas[i]
		}
	}
	return nil
}
