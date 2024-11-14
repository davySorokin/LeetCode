class Solution {
    public int[] findProductsOfElements(long[][] queries) {
        int[] results = new int[queries.length];
        
        for(int i = 0; i < results.length; i++) {
            results[i] = calculateProduct(queries[i][0], queries[i][1] + 1, queries[i][2]);
        }
        
        return results;
    }
    
    public int calculateProduct(long start, long target, long modulus) {
        long initialLimit = start > 0 ? determineLimit(1L, 100000_00000_00000L, start) : 0L;
        long[] initialCounts = start > 0 ? computePowerCounts(initialLimit) : new long[51];
        adjustCounts(initialCounts, initialLimit, start);

        long finalLimit = determineLimit(1L, 100000_00000_00000L, target);
        long[] targetCounts = computePowerCounts(finalLimit);
        adjustCounts(targetCounts, finalLimit, target);

        for(int i = 0; i < 50; i++) {
            targetCounts[i] -= initialCounts[i];
        }
        
        long[] powers = new long[50];
        powers[0] = 1L;
        for(int i = 1; i < 50; i++) {
            powers[i] = (powers[i - 1] << 1) % modulus;
        }
        
        long product = 1L;
        long nextPower = 2L;
        targetCounts[50] = 0;

        for(int i = 1; i < 50; i++) {
            int remainder = (int)((targetCounts[i] * i) % (i + 1));
            long quotient = (targetCounts[i] * i) / (i + 1);
            product = (product * powers[remainder]) % modulus;
            
            targetCounts[i + 1] += quotient;
            nextPower <<= 1;
        }
        
        long base = nextPower % modulus;
        while(targetCounts[50] > 0) {
            if (targetCounts[50] % 2 == 1) product = (product * base) % modulus;
            base = (base * base) % modulus;
            targetCounts[50] >>= 1;
        }
        
        return (int)(product % modulus);
    }
    
    public void adjustCounts(long[] counts, long number, long goal) {
        int index = 0;
        while(counts[50] != goal) {
            if(number == 0) {
                System.out.println("Error");
                break;
            }
            if(number % 2 != 0) {
                counts[index]++;
                counts[50]++;
            }
            
            number >>= 1;
            index++;
        }
    }
    
    public long determineLimit(long lower, long upper, long target) {
        if(lower == upper) {
            return lower;
        }
        long midpoint = lower + (upper + 1 - lower) / 2;
        long[] counts = computePowerCounts(midpoint);
        if(counts[50] > target) {
            return determineLimit(lower, midpoint - 1, target);
        } else {
            return determineLimit(midpoint, upper, target);
        }
    }
    
    public long[] computePowerCounts(long value) {
        long[] counts = new long[51];
        long total = 0;
        int index = 0;
        long current = 1L;
        
        while(current < value) {
            counts[index] = (value / (2 * current)) * current + Math.max((value % (2 * current)) - current, 0);
            total += counts[index];
            index++;
            current <<= 1;
        }
        
        counts[50] = total;
        return counts;
    }
}
