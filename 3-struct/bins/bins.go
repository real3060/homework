package bins

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type Bins struct {
	Bins [](*Bin)
}

func NewBin(id, name string, private bool) (*Bin, error) {

	if id == "" {
		rand.NewSource(time.Now().UnixNano())
		id = strconv.FormatInt(rand.Int63(), 10)
	}
	if name == "" {
		name = "name_" + string(rune(rand.Intn(100)))
	}

	return &Bin{
		CreatedAt: time.Now(),
		Id:        id,
		Private:   private,
		Name:      name,
	}, nil
}
func (bins *Bins) ToBytes() ([]byte, error) {
	data, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("не удалось привести к байтам bins")
		return nil, err
	}
	return data, nil
}
func (bins *Bins) ShowBins() {
	if len(bins.Bins) == 0 {
		fmt.Println("тут еще пока ничего нет")
		return
	}
	for _, bin := range bins.Bins {
		fmt.Println(bin)
	}
}
