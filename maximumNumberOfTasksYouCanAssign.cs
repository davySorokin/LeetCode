using System;
using System.Collections.Generic;

public class Solution {
    public int MaxTaskAssign(int[] tasks, int[] workers, int pills, int strength) {
        Array.Sort(tasks);
        Array.Sort(workers);

        int left = 0, right = Math.Min(tasks.Length, workers.Length), result = 0;

        while (left <= right) {
            int mid = (left + right) / 2;
            if (CanAssign(mid, tasks, workers, pills, strength)) {
                result = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return result;
    }

    private bool CanAssign(int k, int[] tasks, int[] workers, int pills, int strength) {
        var taskList = new List<int>();
        for (int i = 0; i < k; i++) taskList.Add(tasks[tasks.Length - k + i]);

        var deque = new LinkedList<int>();
        for (int i = 0; i < k; i++) deque.AddLast(workers[workers.Length - k + i]);

        int pillCount = pills;

        while (taskList.Count > 0) {
            int task = taskList[taskList.Count - 1];
            taskList.RemoveAt(taskList.Count - 1);

            if (deque.Count > 0 && deque.Last.Value >= task) {
                deque.RemoveLast();
            } else if (pillCount > 0) {
                if (deque.Count > 0 && deque.First.Value + strength >= task) {
                    deque.RemoveFirst();
                    pillCount--;
                } else {
                    return false;
                }
            } else {
                return false;
            }
        }

        return true;
    }
}
