package main

import (
	"net/http"

	"url/cmd/sevice_provider"
	"url/internal/config"
)

func main() {
	//storageType := flag.String("storage", "postgresql", "Тип хранилища (in-memory или postgresql)")
	//flag.Parse()
	//var c *config.Config
	//
	//switch *storageType {
	//case "in-memory":
	//	c = config.NewConfig(*storageType)
	//case "postgresql":
	//	c = config.NewConfig(*storageType)
	//default:
	//	panic("данный тип не але")
	//}
	//fmt.Println(c.DatabaseType)

	sp := sevice_provider.NewServiceProvider(config.NewConfig())

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}

//манипулировать структурой url, а не полями как у меня
//возвращать json вместе с короткой длинную тоже
//сделать поиск по подстроке, так чтобы короткая сслыка была ссылкой
//переделать юзкейс:
//type Service struct {
//	storage storage.Storage
//}
//
//func NewService(store storage.Storage) *Service {
//	return &Service{storage: store}
//}
//
//const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
//
//func (s *Service) Shorten(url string) (string, error) {
//	shortURL := generateShortURL(10)
//	err := s.storage.Save(shortURL, url)
//	return shortURL, err
//}
//
//func (s *Service) GetOriginal(shortURL string) (string, error) {
//	return s.storage.Get(shortURL)
//}
//
//func generateShortURL(length int) string {
//	b := make([]byte, length)
//	for i := range b {
//		b[i] = charset[rand.Intn(len(charset))]
//	}
//	return string(b)
//}

// почитать про baseurl и  basepath
