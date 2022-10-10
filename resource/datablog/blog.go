package datablog

type Blog struct {
	ID      uint8  `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Like    uint8  `json:"like"`
}

var Blogs = []Blog{
	{ID: 1, Title: "KKN : Kok Kaga Ngerti-ngerti?", Author: "Ale", Content: "Belum lama ini gw baru aja ikut kegiatan wajib dari kampus yang mengharuskan gw ketemu dan kerja bareng orang-orang baru."},
	{ID: 2, Title: "Quarter Life Crisis from Broadway", Author: "Ale", Content: "Bimeline waktu di film ini terbagi menjadi dua bagian; saat penayangan Rent & bagaimana keseharian Jonathan Larson selama membuat Superbia. Superbia merupakan projek pertama Jonathan yang dia buat selama 7 tahun, yang nantinya punya peran penting dalam pembuatan Rent."},
}