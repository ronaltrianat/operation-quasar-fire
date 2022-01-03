package domain

// Position Posicion X y Y en algun punto del plano
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Satellite Detalle de cada satelite. Nombre del tripulante del satelite,
// distancia del satelite a la nave imperial, el mensaje que esta trasmitiendo y
// su posicion actual.
type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
	Position Position `json:"position"`
}

// SatellitesData Cuerpo del servicio topsecret que devuelve
// la posicion de la nave imperial y el mensaje secreto
type SatellitesData struct {
	Satellites []Satellite `json:"satellites"`
}

// TopSecret docs
type TopSecret struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
