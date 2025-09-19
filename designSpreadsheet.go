package main

import (
    "strconv"
    "unicode"
)

type Spreadsheet struct {
    rows int
    grid [][]int
}

func Constructor(rows int) Spreadsheet {
    grid := make([][]int, rows)
    for i := range grid {
        grid[i] = make([]int, 26)
    }
    return Spreadsheet{rows: rows, grid: grid}
}

func (s *Spreadsheet) parseCell(cell string) (int, int) {
    col := int(cell[0] - 'A')
    row, _ := strconv.Atoi(cell[1:])
    return row - 1, col
}

func (s *Spreadsheet) SetCell(cell string, value int) {
    r, c := s.parseCell(cell)
    s.grid[r][c] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
    r, c := s.parseCell(cell)
    s.grid[r][c] = 0
}

func (s *Spreadsheet) GetValue(formula string) int {
    parts := []rune(formula[1:]) // skip '='
    splitIdx := -1
    for i, ch := range parts {
        if ch == '+' {
            splitIdx = i
            break
        }
    }
    left := string(parts[:splitIdx])
    right := string(parts[splitIdx+1:])

    return s.evalToken(left) + s.evalToken(right)
}

func (s *Spreadsheet) evalToken(token string) int {
    if len(token) > 0 && unicode.IsLetter(rune(token[0])) {
        r, c := s.parseCell(token)
        return s.grid[r][c]
    }
    val, _ := strconv.Atoi(token)
    return val
}
