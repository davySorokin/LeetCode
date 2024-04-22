from typing import List

class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        WORD_KEY = '$'
        
        # Create a trie data structure to hold all of our words
        trie = {}
        for word in words:
            node = trie
            for letter in word:
                # Retrieve the next node; If not found, create a new empty node.
                node = node.setdefault(letter, {})
            # Mark the existence of a word in trie node
            node[WORD_KEY] = word
        
        rowNum = len(board)
        colNum = len(board[0])
        matchedWords = []

        def backtracking(row, col, parent):
            letter = board[row][col]
            currNode = parent[letter]

            # Check if we find a match of word
            word_match = currNode.pop(WORD_KEY, False)
            if word_match:
                # Also we removed the matched word to avoid duplicates,
                # as the question stated that all words are unique.
                matchedWords.append(word_match)

            # Mark the current letter on the board as '#' to indicate this letter is used
            board[row][col] = '#'
            
            # Explore the neighbors in DFS
            for (rowOffset, colOffset) in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                newRow, newCol = row + rowOffset, col + colOffset     
                if newRow < 0 or newRow >= rowNum or newCol < 0 or newCol >= colNum:
                    continue
                if not board[newRow][newCol] in currNode:
                    continue
                backtracking(newRow, newCol, currNode)

            # End of exploration, we restore the letter in the board.
            board[row][col] = letter

            # Optimization: incrementally remove the leaf nodes in Trie.
            if not currNode:
                parent.pop(letter)

        for row in range(rowNum):
            for col in range(colNum):
                # Starting from each of the board's letters that matches the first letter of a word
                if board[row][col] in trie:
                    backtracking(row, col, trie)
        
        return matchedWords
