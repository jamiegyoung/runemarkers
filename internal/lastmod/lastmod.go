package lastmod

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/jamiegyoung/runemarkers-go/internal/entities"
	"github.com/jamiegyoung/runemarkers-go/internal/logger"
	bolt "go.etcd.io/bbolt"
)

type EntityMod struct {
	Entity        entities.Entity
	ModTimeString string
	ModTime       time.Time
}

const bucketName string = "lastmod"

var log = logger.New("lastmod")

func GenerateDb() error {
	db, err := db()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	return nil
}

// Updates or inserts a directory's last modified time into the database
func UpdateEntities(entity_arr []*entities.Entity) error {
	db, err := db()
	if err != nil {
		return err
	}
	defer db.Close()

	gob.Register(map[string]interface{}{})

	err = db.Batch(func(tx *bolt.Tx) error {
		for _, e := range entity_arr {
			log(fmt.Sprintf("Updating 1 %v", e.SafeUri))
			b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)

			modTime := time.Now().UTC()

			err = enc.Encode(
				EntityMod{
					Entity:        *e,
					ModTime:       modTime,
					ModTimeString: Format(modTime),
				},
			)
			if err != nil {
				return err
			}

			log("Updating " + e.SafeUri)
			err = b.Put([]byte(e.SafeUri), buf.Bytes())
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// Returns a subset of the entities provided that have been modified since the last time they were updated
func EntitiesDiff(entities_arr []*entities.Entity) ([]*entities.Entity, error) {
	db, err := db()
	modded := []*entities.Entity{}
	if err != nil {
		return nil, err
	}
	defer db.Close()

	gob.Register(map[string]interface{}{})

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("Bucket does not exist")
		}

		for _, e := range entities_arr {
			res := b.Get([]byte(e.SafeUri))
			dec := gob.NewDecoder(bytes.NewReader(res))
			if len(res) == 0 {
				log(fmt.Sprintf("Adding %v to modded list", e.SafeUri))
				modded = append(modded, e)
				continue
			}

			var decEntity EntityMod
			err = dec.Decode(&decEntity)
			if err != nil {
				return err
			}

			// Compare decoded entity to current entity
			isDiff := reflect.DeepEqual(e, decEntity.Entity)
			log("Checking " + e.SafeUri + " diff " + strconv.FormatBool(isDiff))
			if isDiff {
				log(fmt.Sprintf("Adding %v to modded list", e.SafeUri))
				modded = append(modded, e)
			}
		}
		return err
	})
	return modded, err
}

func GetEntites() ([]*EntityMod, error) {
	db, err := db()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	gob.Register(map[string]interface{}{})

	var entitiesArr []*EntityMod

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("Bucket does not exist")
		}

		b.ForEach(func(k, v []byte) error {
			log(fmt.Sprintf("%v found", string(k[:])))
			dec := gob.NewDecoder(bytes.NewReader(v))
			var decEntity EntityMod
			err = dec.Decode(&decEntity)
			if err != nil {
				return err
			}

			entitiesArr = append(entitiesArr, &decEntity)

			return nil
		})
		return nil
	})

	return entitiesArr, err
}

func FindLastMod(entMods []*EntityMod) time.Time {
	lastmod := entMods[0].ModTime
	for _, entMod := range entMods {
		if entMod.ModTime.After(lastmod) {
			lastmod = entMod.ModTime
		}
	}

	return lastmod
}

func Format(t time.Time) string {
	return t.Format(time.RFC3339)
}

func db() (*bolt.DB, error) {
	return bolt.Open("lastmod.db", 0600, nil)
}
