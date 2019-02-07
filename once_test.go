package xsync

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type one int

func (o *one) Increment() error {
	*o++

	return nil
}

func run(t *testing.T, once *Once, o *one, c chan bool) {

	once.Do(func() error {
		return o.Increment()
	})

	v := *o
	assert.EqualValues(t, 1, v, "once failed inside run: %d is not 1", v)

	c <- true
}

func TestOnce(t *testing.T) {

	o := new(one)
	once := new(Once)
	c := make(chan bool)

	const N = 10

	for i := 0; i < N; i++ {
		go run(t, once, o, c)
	}

	for i := 0; i < N; i++ {
		<-c
	}

	assert.EqualValues(t, 1, *o, "once failed outside run: %d is not 1", *o)
}

func TestOncePanic(t *testing.T) {

	var (
		once    Once
		counter = 0
	)

	func() {
		defer func() {
			if r := recover(); r == nil {
				assert.FailNow(t, "Once.Do did not panic")
			}
		}()

		once.Do(func() error {
			panic("failed")
		})
	}()

	once.Do(func() error {
		counter += 1
		return nil
	})
	assert.EqualValues(t, 1, counter)

	once.Do(func() error {
		counter += 1
		return nil
	})
	assert.EqualValues(t, 1, counter)
}

func TestOnceError(t *testing.T) {

	var (
		once    Once
		counter = 0
	)

	once.Do(func() error {
		counter += 1
		return errors.New("some error")
	})
	assert.EqualValues(t, 1, counter)

	once.Do(func() error {
		counter += 1
		return nil
	})
	assert.EqualValues(t, 2, counter)
}
