from typing import List

class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        square, circular = students.count(0), students.count(1)
        
        for sandwich in sandwiches:
            if sandwich == 0 and square > 0:
                square -= 1
            elif sandwich == 1 and circular > 0:
                circular -= 1
            else:
                # if no student wants the current type of sandwich, break the loop
                break
                
        # the remaining students are unable to eat
        return square + circular
