import threading

class DiningPhilosophers:
    def __init__(self):
        self.forks = [threading.Lock() for _ in range(5)]
        self.access_lock = threading.Lock()

    def wantsToEat(self,
                   philosopher: int,
                   pickLeftFork: 'Callable[[], None]',
                   pickRightFork: 'Callable[[], None]',
                   eat: 'Callable[[], None]',
                   putLeftFork: 'Callable[[], None]',
                   putRightFork: 'Callable[[], None]') -> None:
        
        left_fork = philosopher
        right_fork = (philosopher + 1) % 5

        with self.access_lock:
            # Lock the left fork, then the right fork in sequence
            with self.forks[left_fork]:
                pickLeftFork()
                with self.forks[right_fork]:
                    pickRightFork()
                    # Both forks are picked; philosopher can eat
                    eat()
                    # Put down forks after eating
                    putRightFork()
                putLeftFork()
