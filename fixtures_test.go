package fixtures_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/knqyf263/bolt-fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	bolt "go.etcd.io/bbolt"
)

func TestLoader_Load(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		f, err := ioutil.TempFile("", "TestLoad")
		require.NoError(t, err)
		require.NoError(t, f.Close())
		defer os.Remove(f.Name())

		l, err := fixtures.New(f.Name(), []string{"testdata/test.yaml"})
		require.NoError(t, err)

		// load
		require.NoError(t, l.Load())
		require.NoError(t, l.Close())

		// evaluate data
		db, err := bolt.Open(f.Name(), 0666, nil)
		require.NoError(t, err)

		err = db.View(func(tx *bolt.Tx) error {
			// first bucket
			root := tx.Bucket([]byte("abc"))
			nested := root.Bucket([]byte("def"))
			value := nested.Get([]byte("ghi"))
			assert.Equal(t, "jkl", string(value))

			// second bucket
			root = tx.Bucket([]byte("mno"))
			value = root.Get([]byte("pqr"))
			assert.Equal(t, "stu", string(value))
			value = root.Get([]byte("vwx"))
			assert.JSONEq(t, `{"foo":"abc","bar":123}`, string(value))
			return nil
		})
		require.NoError(t, err)
	})

	t.Run("invalid yaml", func(t *testing.T) {
		f, err := ioutil.TempFile("", "TestLoad")
		require.NoError(t, err)
		require.NoError(t, f.Close())
		defer os.Remove(f.Name())

		l, err := fixtures.New(f.Name(), []string{"testdata/invalid.yaml"})
		require.NoError(t, err)

		err = l.Load()
		assert.EqualError(t, err, "String node doesn't ArrayNode")

	})
}
