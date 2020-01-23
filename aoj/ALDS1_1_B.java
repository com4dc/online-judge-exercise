import java.util.Scanner;

public class ALDS1_1_B {
    
    public static void main(String... args) {
        // Scanner scanner = new Scanner(System.in);
        
        // String[] numbers = scanner.nextLine().split(" ");
        
        // int x = Integer.parseInt(numbers[0]);
        // int y = Integer.parseInt(numbers[1]);
        int x = 93;
        int y = 93;
        
        if (x >= y) {
            System.out.println(gcd(y, x%y) + "");
        } else {
            System.out.println(gcd(x, y) + "");
        }
    }
    
    private static int gcd(int x, int y) {
        if (y == 0) {
            return x;
        }
        int d = 0;
        for (int i = 1; i <= y / 2; i++) {
            if (x % i == 0 && y % i == 0) {
                d = i;
            }
        }
        return d;
    }
}
