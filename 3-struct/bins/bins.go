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

func NewBin(id, name string, private bool) *Bin {
	newB := Bin{
		CreatedAt: time.Now(),
		Id:        id,
		Private:   private,
		Name:      name,
	}
	if newB.Id == "" {
		rand.NewSource(time.Now().UnixNano())
		id := rand.Int63()
		newB.Id = strconv.FormatInt(id, 10)
	}
	if newB.Name == "" {
		newB.Name = "name_" + string(rune(rand.Intn(100)))
	}
	newB.Private = private

	return &newB
}
func (bins *Bins) ToBytes() []byte {
	data, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("не удалось привести к байтам bins")
		return nil
	}
	return data
}
func (bins *Bins) ShowBins() {
	for _, bin := range bins.Bins {
		fmt.Println(bin)
	}
}
