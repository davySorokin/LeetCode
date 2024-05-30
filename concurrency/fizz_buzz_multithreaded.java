import java.util.concurrent.Semaphore;
import java.util.function.IntConsumer;

class FizzBuzz {
    private int n;
    private Semaphore semFizz = new Semaphore(0);
    private Semaphore semBuzz = new Semaphore(0);
    private Semaphore semFizzBuzz = new Semaphore(0);
    private Semaphore semNumber = new Semaphore(1);

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 != 0) {
                semFizz.acquire();
                printFizz.run();
                semNumber.release();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 5 == 0 && i % 3 != 0) {
                semBuzz.acquire();
                printBuzz.run();
                semNumber.release();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            if (i % 15 == 0) {
                semFizzBuzz.acquire();
                printFizzBuzz.run();
                semNumber.release();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i++) {
            semNumber.acquire();
            if (i % 3 != 0 && i % 5 != 0) {
                printNumber.accept(i);
                semNumber.release();
            } else if (i % 3 == 0 && i % 5 != 0) {
                semFizz.release();
            } else if (i % 5 == 0 && i % 3 != 0) {
                semBuzz.release();
            } else if (i % 15 == 0) {
                semFizzBuzz.release();
            }
        }
    }
}
