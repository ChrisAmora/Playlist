package presentation

type AppController struct {
	Music interface{ MusicController }
}
