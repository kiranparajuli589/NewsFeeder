package newsfeed


type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

type ItemList struct {
	Items []Item
}

func New() *ItemList {
	return &ItemList{
		Items: []Item{},
	}
}

func (r *ItemList) Add(item Item)  {
	r.Items = append(r.Items, item)
}

func (r *ItemList) GetAll() []Item  {
	return r.Items
}
