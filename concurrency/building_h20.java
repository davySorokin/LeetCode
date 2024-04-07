import java.util.concurrent.Semaphore;

class H2O {
    private final Semaphore hydro = new Semaphore(2);
    private final Semaphore oxy = new Semaphore(0);
    
    public H2O() {}

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        hydro.acquire(); // Hydrogen thread will acquire a permit
        releaseHydrogen.run();
        if (hydro.availablePermits() == 0) {
            oxy.release(); // If two hydrogens are released, it's time to release one oxygen
        }
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        oxy.acquire(); // Oxygen thread will try to acquire a permit
        releaseOxygen.run();
        hydro.release(2); // Once oxygen is released, it gives back permits for two more hydrogens
    }
}
