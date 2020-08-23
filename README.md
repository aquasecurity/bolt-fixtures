# bolt-fixtures
Test fixtures for [BoltDB](https://github.com/etcd-io/bbolt). Write tests against a real database.

## Prepare test data

Example YAML fixture:

```
- bucket: abc
  pairs:
    - bucket: def
      pairs:
        - key: ghi
          value: jkl
- bucket: mno
  pairs:
    - key: pqr
      value: stu
    - key: vwx
      value:
        foo: abc
        bar: 123
```

## Load into BoltDB

Example integration for your project:

```
func TestSomething(t *testing.T) {
	f, _ := ioutil.TempFile("", "TestSomething")
	defer os.Remove(f.Name())

	fixtures := []string{"testdata/test.yaml"})
	l, _ := New(f.Name(), fixtures)
	defer l.Close()

	_ = l.Load()

	// do something
```