package rest

type AppController struct {
	Music interface{ MusicController }
}
