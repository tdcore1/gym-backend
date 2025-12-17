package app

import "gym/ent"

type App struct {
	Client *ent.Client
	Key    []byte
}
