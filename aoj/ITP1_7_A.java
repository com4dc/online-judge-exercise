import java.util.Scanner;

public class Main {
    
    public static void main(String... args) {
        Scanner scanner = new Scanner(System.in);
        
        while (true) {
            String[] scores = scanner.nextLine().split(" ");
            
            if (scores[0].equals("-1") && scores[1].equals("-1") && scores[2].equals("-1")) {
                break;
            }
            
            System.out.println(validateScore(scores));
        }
    }
    
    private static String validateScore(String[] target) {
        int m = Integer.parseInt(target[0]);
        int f = Integer.parseInt(target[1]);
        int r = Integer.parseInt(target[2]);
        
        if (m == -1 || f == -1) {
            return "F";
        } else {
            int total = m + f;
            if (total >= 80) return "A";
            if (total >= 65 && total < 80) return "B";
            if (total >= 50 && total < 65) return "C";
            if (total >= 30 && total < 50) {
                if (r >= 50) return "C";
                else return "D";
            }
            return "F";
        }
    }
}
