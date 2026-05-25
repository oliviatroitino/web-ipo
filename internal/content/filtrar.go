package content

import (
	"strings"
)

// Filtrar devuelve las técnicas que coinciden con todos los criterios dados.
// Los slices de filtros vacíos se ignoran (no filtran nada).
// La búsqueda por texto q es insensible a mayúsculas y busca en varios campos.
func Filtrar(tecnicas []Tecnica, q string, clasificacion, fases, recursos, duracion []string) []Tecnica {
	var result []Tecnica
	q = strings.ToLower(strings.TrimSpace(q))

	for _, t := range tecnicas {
		if len(clasificacion) > 0 && !contains(clasificacion, t.Clasificacion) {
			continue
		}
		if len(fases) > 0 && !anyMatch(fases, t.Fases) {
			continue
		}
		if len(recursos) > 0 && !contains(recursos, t.Recursos) {
			continue
		}
		if len(duracion) > 0 && !contains(duracion, t.Duracion) {
			continue
		}
		if q != "" && !matchesQuery(t, q) {
			continue
		}
		result = append(result, t)
	}
	return result
}

// matchesQuery comprueba si q aparece en alguno de los campos de texto de la técnica.
func matchesQuery(t Tecnica, q string) bool {
	if strings.Contains(strings.ToLower(t.Nombre), q) {
		return true
	}
	if strings.Contains(strings.ToLower(t.Clasificacion), q) {
		return true
	}
	for _, f := range t.Fases {
		if strings.Contains(strings.ToLower(f), q) {
			return true
		}
	}
	for _, s := range t.ParaQueSirve {
		if strings.Contains(strings.ToLower(s), q) {
			return true
		}
	}
	for _, p := range t.Pasos {
		if strings.Contains(strings.ToLower(p.Titulo), q) ||
			strings.Contains(strings.ToLower(p.Descripcion), q) {
			return true
		}
	}
	return false
}

// contains comprueba si val está en slice (comparación exacta).
func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// anyMatch comprueba si algún elemento de targets aparece en source.
func anyMatch(targets, source []string) bool {
	for _, t := range targets {
		if contains(source, t) {
			return true
		}
	}
	return false
}
