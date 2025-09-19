class Spreadsheet:
    def __init__(self, rows: int):
        self.rows = rows
        self.grid = [[0] * 26 for _ in range(rows)]

    def _parse_cell(self, cell: str):
        col = ord(cell[0]) - ord('A')
        row = int(cell[1:]) - 1
        return row, col

    def setCell(self, cell: str, value: int) -> None:
        r, c = self._parse_cell(cell)
        self.grid[r][c] = value

    def resetCell(self, cell: str) -> None:
        r, c = self._parse_cell(cell)
        self.grid[r][c] = 0

    def getValue(self, formula: str) -> int:
        parts = formula[1:].split('+')  # remove '=' and split
        total = 0
        for part in parts:
            if part[0].isalpha():  # cell reference
                r, c = self._parse_cell(part)
                total += self.grid[r][c]
            else:  # integer
                total += int(part)
        return total
