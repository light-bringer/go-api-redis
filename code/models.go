package main

import {
	"time"
	"uuid"
}

type Song struct {
	Id        int       `json:"id`
	SongName  string    `json:"name"`
	Album     string    `json:"album"`
	Year      int       `json:"year"`
	TimeStamp time.Time `json:"timestamp"`
}

type MQ struct {
	

}

type Songs []Song