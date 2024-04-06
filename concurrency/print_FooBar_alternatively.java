import java.util.concurrent.Semaphore;

class FooBar {
    private int n;
    private Semaphore fooSemaphore = new Semaphore(1);
    private Semaphore barSemaphore = new Semaphore(0);

    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            fooSemaphore.acquire(); // Acquire the foo semaphore.
            printFoo.run(); // This outputs "foo".
            barSemaphore.release(); // Release the bar semaphore, allowing bar() to execute.
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            barSemaphore.acquire(); // Acquire the bar semaphore.
            printBar.run(); // This outputs "bar".
            fooSemaphore.release(); // Release the foo semaphore, allowing foo() to execute.
        }
    }

    public static void main(String[] args) {
        FooBar fooBar = new FooBar(3); // 
        Thread threadFoo = new Thread(() -> {
            try {
                fooBar.foo(() -> System.out.print("foo"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        Thread threadBar = new Thread(() -> {
            try {
                fooBar.bar(() -> System.out.print("bar"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        threadFoo.start();
        threadBar.start();
    }
}
