package data

import (
	"context"

	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type MusicRepository interface {
	GetById(c context.Context, id int64) (domain.Music, error)
	Add(c context.Context, id int64) (domain.Music, error)
	GetAll(c context.Context) ([]domain.Music, error)
}

type AuthRepository interface {
	GetUser(c context.Context, email string) (Auth, error)
	CreateUser(c context.Context, email, password string) (*Auth, error)
}

type TrackRepository interface {
	SaveTrack(c context.Context, playListID int, title, album, artist string) (*Track, error)
}

type ProviderRepository interface {
	GetTokensData(c context.Context, authorizationCode string) (Tokens, error)
}

type JWTRepository interface {
	Sign(c context.Context, username string) (string, error)
	Verify(c context.Context, token string) (*Claims, error)
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type authUsecase struct {
	AuthRepository
	JWTRepository
}

type musicUsecase struct {
	MusicRepository
}

type trackUsecase struct {
	TrackRepository
}

type providerUsecase struct {
	ProviderRepository
}

func NewMusicUsecase(mr MusicRepository) domain.MusicUsecase {
	return &musicUsecase{
		MusicRepository: mr,
	}
}

func NewAuthUsecase(ar AuthRepository, jr JWTRepository) domain.AuthUsecase {
	return &authUsecase{
		AuthRepository: ar,
		JWTRepository:  jr,
	}
}

func NewTrackUsecase(tr TrackRepository) domain.TrackUsecase {
	return &trackUsecase{
		TrackRepository: tr,
	}
}

func NewProviderUsecase(pr ProviderRepository) domain.ProviderUsecase {
	return &providerUsecase{}
}

func (pu *providerUsecase) GetTokens(c context.Context, authorizationCode string) (domain.UserTokens, error) {
	userTokens := &domain.UserTokens{}
	tokens, err := pu.ProviderRepository.GetTokensData(c, authorizationCode)
	if err != nil {
		return *userTokens, err
	}
	userTokens.AccessToken = tokens.AccessToken
	userTokens.RefreshToken = tokens.RefreshToken
	return *userTokens, err
}

func (tu *trackUsecase) SaveTrack(c context.Context, playListID int, title, album, artist string) (domain.Track, error) {
	track, err := tu.TrackRepository.SaveTrack(c, playListID, title, album, artist)
	domainTrack := &domain.Track{}
	if err != nil {
		return *domainTrack, err
	}
	domainTrack.Album = track.Album
	domainTrack.Artist = track.Artist
	domainTrack.CreatedAt = track.CreatedAt
	domainTrack.DeletedAt = track.DeletedAt.Time
	domainTrack.PlayListID = track.PlayListID
	domainTrack.Title = track.Title
	domainTrack.UpdatedAt = track.UpdatedAt
	return *domainTrack, err
}

func (au *authUsecase) Signup(c context.Context, email, password string) (domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, err
	}
	auth, err := au.AuthRepository.CreateUser(c, email, string(hash))
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{Email: auth.Email}, nil
}

func (au *authUsecase) Login(c context.Context, email, password string) (domain.Auth, error) {
	user, err := au.AuthRepository.GetUser(c, email)
	if err != nil {
		return domain.Auth{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.Auth{}, err
	}
	token, err := au.JWTRepository.Sign(c, email)
	if err != nil {
		return domain.Auth{}, err
	}
	return domain.Auth{User: domain.User{Email: email}, Token: token}, err
}

func (mu *musicUsecase) Add(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.Add(c, id)
}

func (mu *musicUsecase) GetById(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.GetById(c, id)
}

func (mu *musicUsecase) GetAllMusics(c context.Context) ([]domain.Music, error) {
	return mu.MusicRepository.GetAll(c)
}
