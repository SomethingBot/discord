package dizzy

import "testing"

func Test_syncedCounter(t *testing.T) {
	c := syncedCounter{}
	if c.count() != 0 {
		t.Fatal("syncedCounter not 0")
	}
	c.add()
	if c.count() != 1 {
		t.Fatal("syncedCounter not 1")
	}
	c.set(40)
	if c.count() != 40 {
		t.Fatal("syncedCounter not 40")
	}
}
