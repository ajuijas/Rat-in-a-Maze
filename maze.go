package main

import "fmt"

type Maze struct {
    Cells [][]Cell
    Start struct{ x, y int } // Start coordinates
    End   struct{ x, y int } // End coordinates
}

type Cell struct {
    south, west bool // we need to store the south and west walls only. other walls can be inferred from neighbor cells
}

func (m Maze) Print() {
    width := len(m.Cells[0])

    // Print top border
    for j := 0; j < width; j++ {
        fmt.Print("+--")
    }
    fmt.Println("+")

    for i, row := range m.Cells {
        // Print left wall of each row
        if row[0].west {
            fmt.Print("|")
        } else {
            fmt.Print(" ")
        }

        // Print cell content
        for j := range row {
            if i == m.Start.x && j == m.Start.y {
                fmt.Print(" S") // Mark start position
            } else if i == m.End.x && j == m.End.y {
                fmt.Print(" E") // Mark end position
            } else {
                fmt.Print("  ") // Empty space for the cell
            }
            if j == width-1 || m.Cells[i][j+1].west {
                fmt.Print("|") // Right wall
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()

        // Print bottom walls
        for _, cell := range row {
            fmt.Print("+")
            if cell.south {
                fmt.Print("--")
            } else {
                fmt.Print("  ")
            }
        }
        fmt.Println("+")
    }
}

var maze Maze = Maze{
    Cells: [][]Cell{
        {Cell{false, true}, Cell{false, false}, Cell{true, false}, Cell{true, false}},
        {Cell{false, true}, Cell{true, true}, Cell{true, false}, Cell{false, false}},
        {Cell{false, true}, Cell{true, true}, Cell{true, false}, Cell{true, false}},
        {Cell{true, true}, Cell{true, false}, Cell{true, false}, Cell{true, false}},
    },
    Start: struct{ x, y int }{0, 0},
    End:   struct{ x, y int }{2, 1},
}

func main() {
    maze.Print()
}
