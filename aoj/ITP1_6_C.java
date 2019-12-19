import java.util.Scanner;

public class Main {
    
    private static final int FLOOR = 3;
    private static final int ROOMS = 10;
    
    public static void main (String... args) {
        int[][] w1 = new int[FLOOR][ROOMS]; // matrixで表記
        int[][] w2 = new int[FLOOR][ROOMS]; // matrixで表記
        int[][] w3 = new int[FLOOR][ROOMS]; // matrixで表記
        int[][] w4 = new int[FLOOR][ROOMS]; // matrixで表記
        
        Scanner scanner = new Scanner(System.in);
        int count = Integer.parseInt(scanner.nextLine());
        
        for (int i=0; i<count; i++) {
            String[] numbers = scanner.nextLine().split(" ");
            int w = Integer.parseInt(numbers[0]);
            int f = Integer.parseInt(numbers[1]) - 1; // indexに使う
            int r = Integer.parseInt(numbers[2]) - 1; // indexに使う
            int person = Integer.parseInt(numbers[3]);
            
            switch (w) {
                case 1: 
                    w1[f][r] = w1[f][r] + person;
                    break;
                case 2:
                    w2[f][r] = w2[f][r] + person;
                    break;
                case 3:
                    w3[f][r] = w3[f][r] + person;
                    break;
                case 4:
                    w4[f][r] = w4[f][r] + person;
                    break;
            }
        }
        outputMatrix(w1);
        System.out.println("####################");
        outputMatrix(w2);
        System.out.println("####################");
        outputMatrix(w3);
        System.out.println("####################");
        outputMatrix(w4);
    }
    
    private static void outputMatrix(int[][] matrix) {
        for (int i=0; i<FLOOR; i++) {
            for (int j=0; j<ROOMS; j++) {
                System.out.print(String.format(" %d", matrix[i][j]));
            }
            System.out.println();
        }
    }
}
