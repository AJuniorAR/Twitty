package models

import "time"

//GraboTweet es el formato o estructura que tendra nuestro Tweet
type GraboTweet struct {
	UserID  string    `bson:"userId" json:"userId,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
