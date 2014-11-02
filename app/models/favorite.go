package models

type Favorite struct {
	Id, Created, PinId int64
}

func (f Favorite) All() FavoriteView {

	var favorites []Favorite
	DB.Find(&favorites)

	return FavoriteView{Favorites: favorites}
}
