import java.util.Scanner;

public class Main {
    
    public static void main(String... args) {
        Scanner scanner = new Scanner(System.in);
        int count = Integer.parseInt(scanner.nextLine());
        
        boolean[] club = new boolean[]{false,false,false,false,false,false,false,false,false,false,false,false,false};
        boolean[] spade = new boolean[]{false,false,false,false,false,false,false,false,false,false,false,false,false};
        boolean[] heart = new boolean[]{false,false,false,false,false,false,false,false,false,false,false,false,false};
        boolean[] dia = new boolean[]{false,false,false,false,false,false,false,false,false,false,false,false,false};
        
        for (int i=0; i<count; i++) {
            String[] cards = scanner.nextLine().split(" ");
            String symbol = cards[0];
            int number = Integer.parseInt(cards[1]) - 1;
            
            switch (symbol) {
                case "S":
                    spade[number] = true;
                    break;
                case "H":
                    heart[number] = true;
                    break;
                case "C":
                    club[number] = true;
                    break;
                case "D":
                    dia[number] = true;
                    break;                
                default:
                    break;
            }
        }
        output(spade, "S");
        output(heart, "H");
        output(club, "C");
        output(dia, "D");
    }
    
    private static void output(boolean[] bools, String symbol) {
        for (int i=0; i<bools.length; i++) {
            if (bools[i] == false) {
                System.out.println(symbol + " " + (i+1));
            }
        }
    }
}
