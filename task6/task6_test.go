package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCut(t *testing.T) {
	fields = []int{2, 4}
	var tab Tab
	row := Row{
		Separator: true,
		Column:    []string{"1", "2", "3", "4", "5"},
	}

	tab = append(tab, row)

	row = Row{
		Separator: false,
		Column:    []string{"1gjg5"},
	}

	tab = append(tab, row)

	row = Row{
		Separator: true,
		Column:    []string{"kpkpoki"},
	}

	tab = append(tab, row)

	var tabRes Tab
	rowRes := Row{
		Separator: false,
		Column:    []string{"2", "4"},
	}

	tabRes = append(tabRes, rowRes)

	var tabRes Tab
	rowRes := Row{
		Separator: false,
		Column:    []string{"2", "4"},
	}

	tabRes = append(tabRes, rowRes)

	var tabRes Tab
	rowRes := Row{
		Separator: false,
		Column:    []string{"2", "4"},
	}

	tabRes = append(tabRes, rowRes)

	res := cut(tab)

	require.Equal(t, tabRes, res)
}
