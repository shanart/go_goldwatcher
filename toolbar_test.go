package main

import "testing"

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("Wrong number of items in toolbar")
	}
}
