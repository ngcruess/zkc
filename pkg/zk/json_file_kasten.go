package zk

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonFileKasten struct {
	Kasten
}

func LoadJsonFileKasten(handle string) (*Kasten, error) {
	content, err := os.ReadFile(fmt.Sprintf("%s.json", handle))
	if err != nil {
		return nil, err
	}
	var k *Kasten
	err = json.Unmarshal(content, k)

	return k, err
}

func (k *JsonFileKasten) PersistMeta() error {
	content, err := json.MarshalIndent(k, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", k.Label), content, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

func (k *JsonFileKasten) PersistZettel(z *Zettel) error {
	content, err := json.MarshalIndent(z, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", z.Address), content, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

func (k *JsonFileKasten) RetrieveZettel(address Address) (*Zettel, error) {
	content, err := os.ReadFile(fmt.Sprintf("%s.json", address))
	if err != nil {
		return nil, err
	}
	var z *Zettel
	err = json.Unmarshal(content, z)

	return z, err
}
