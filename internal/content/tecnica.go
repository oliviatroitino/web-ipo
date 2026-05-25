package content

// TiempoFase representa una fase de la técnica y su duración estimada.
type TiempoFase struct {
	Fase   string `yaml:"fase"`
	Tiempo string `yaml:"tiempo"`
}

// Paso representa un paso del procedimiento de la técnica.
type Paso struct {
	Numero      int    `yaml:"numero"`
	Titulo      string `yaml:"titulo"`
	Descripcion string `yaml:"descripcion"`
}

// Tecnica representa la ficha completa de una técnica de usabilidad.
type Tecnica struct {
	// Identificación
	ID     string `yaml:"id"`
	Nombre string `yaml:"nombre"`

	// Clasificación (usados como filtros)
	Clasificacion string   `yaml:"clasificacion"` // Inspección | Prueba | Indagación
	Fases         []string `yaml:"fases"`          // Descubrimiento | Prototipado | Evaluación | Implementación
	Duracion      string   `yaml:"duracion"`       // corta | media | larga
	Recursos      string   `yaml:"recursos"`       // bajo | medio | alto (coste global)

	// Recursos detallados
	CostoMonetario  string       `yaml:"costo_monetario"`  // bajo | medio | alto
	NumUsuariosMin  int          `yaml:"num_usuarios_min"`
	NumUsuariosMax  int          `yaml:"num_usuarios_max"` // 0 = sin límite
	Materiales      []string     `yaml:"materiales"`
	TiemposDetalle  []TiempoFase `yaml:"tiempos_detalle"`

	// Contenido de la ficha
	ParaQueSirve          []string `yaml:"para_que_sirve"`
	ParaQueNoSirve        []string `yaml:"para_que_no_sirve"`
	Pasos                 []Paso   `yaml:"pasos"`
	Resultado             string   `yaml:"resultado"`
	DimensionesUsabilidad []string `yaml:"dimensiones_usabilidad"`

	// Descargable — nombre del archivo PDF en static/plantillas/ (opcional)
	Plantilla string `yaml:"plantilla"`
}
